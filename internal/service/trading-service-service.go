// Package service contains business-logic methods
package service

import (
	"context"
	"fmt"

	"github.com/eugenshima/trading-service/internal/model"
	"github.com/sirupsen/logrus"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// TradingService struct represents a trading-service-service
type TradingService struct {
	rps             TradingRepository
	priceServiceRps PriceServiceRepository
	balanceRps      BalanceRepository
	positionManager *model.PositionManager
}

// NewTradingService creates a new TradingService
func NewTradingService(rps TradingRepository, priceServiceRps PriceServiceRepository, balanceRps BalanceRepository, positionManager *model.PositionManager) *TradingService {
	return &TradingService{
		rps:             rps,
		priceServiceRps: priceServiceRps,
		balanceRps:      balanceRps,
		positionManager: positionManager,
	}
}

// TradingRepository interface represents a trading-service-repository methods
type TradingRepository interface {
	CreatePosition(context.Context, *model.Position) error
	DeletePosition(context.Context, uuid.UUID) error
	GetPositionByID(context.Context, uuid.UUID) (*model.Position, error)
	GetAllIDsPositions(context.Context, uuid.UUID) ([]*model.Position, error)
}

// PriceServiceRepository interface represents a price-service-repository methods
type PriceServiceRepository interface {
	AddSubscriber(context.Context, []string) (*model.Share, error)
}

// BalanceRepository interface represents balance-repository methods
type BalanceRepository interface {
	GetBalance(context.Context, uuid.UUID) (*model.Balance, error)
	UpdateBalance(context.Context, *model.Balance) error
}

// addPositionToMap method adds a position to position manager
func (s *TradingService) addPositionToMap(ProfileID uuid.UUID, position *model.Position) error {
	s.positionManager.Mu.Lock()
	defer s.positionManager.Mu.Unlock()

	s.positionManager.OpenedPositions[ProfileID] = make(map[uuid.UUID]*model.OpenedPosition)
	openedPosition := &model.OpenedPosition{
		PositionID:      position.ID,
		ShareOpenPrice:  position.SharePrice,
		ShareClosePrice: position.StopLoss,
		ShareAmount:     position.ShareAmount,
		IsOpened:        true,
	}
	if _, ok := s.positionManager.OpenedPositions[ProfileID][position.ID]; !ok {
		s.positionManager.OpenedPositions[ProfileID][position.ID] = openedPosition
		return nil
	}
	return fmt.Errorf("error opening position on ID: %v", ProfileID)
}

// deletePositionFromMap method deletes a position from position manager
func (s *TradingService) deletePositionFromMap(ProfileID, positionID uuid.UUID) error {
	s.positionManager.Mu.Lock()
	defer s.positionManager.Mu.Unlock()
	if _, ok := s.positionManager.OpenedPositions[ProfileID]; !ok {
		delete(s.positionManager.Closed, positionID)
		delete(s.positionManager.OpenedPositions[ProfileID], positionID)
		return nil
	}
	return fmt.Errorf("error closing position on ID: %v", ProfileID)
}

// OpenPosition creates a position for a given ID with checking all the necessary conditions
func (s *TradingService) OpenPosition(ctx context.Context, position *model.Position) error {
	balance, err := s.balanceRps.GetBalance(ctx, position.ProfileID)
	if err != nil {
		return fmt.Errorf("GetBalance: %w", err)
	}

	updatedBalance, err := checkIfEnoughMoneyOnBalance(balance, position.Total)
	if err != nil {
		return fmt.Errorf("not enough money on you balance: %w", err)
	}

	selectedShares := []string{position.ShareName}
	share, err := s.priceServiceRps.AddSubscriber(ctx, selectedShares)
	if err != nil {
		return fmt.Errorf("AddSubscriber:%w", err)
	}

	shareAmount, err := calculateAmountOfShares(ctx, position.Total, share.SharePrice)
	if err != nil {
		return fmt.Errorf("calculateAmountOfShares:%w", err)
	}

	position.ShareAmount = shareAmount
	position.SharePrice = share.SharePrice

	err = s.addPositionToMap(position.ID, position)
	if err != nil {
		return fmt.Errorf("addPositionToMap: %w", err)
	}

	err = s.rps.CreatePosition(ctx, position)
	if err != nil {
		return fmt.Errorf("CreatePosition:%w", err)
	}

	err = s.balanceRps.UpdateBalance(ctx, updatedBalance)
	if err != nil {
		return fmt.Errorf("UpdateBalance:%w", err)
	}

	return nil
}

// ClosePosition method closes the position of given ID
func (s *TradingService) ClosePosition(ctx context.Context, PositionID uuid.UUID) (float64, error) {
	position, err := s.rps.GetPositionByID(ctx, PositionID)
	if err != nil {
		return 0, fmt.Errorf("GetPositionByID: %w", err)
	}
	err = s.deletePositionFromMap(position.ProfileID, PositionID)
	if err != nil {
		return 0, fmt.Errorf("deletePositionFromMap:%w", err)
	}
	balance, err := s.balanceRps.GetBalance(ctx, position.ProfileID)
	if err != nil {
		return 0, fmt.Errorf("GetBalance: %w", err)
	}
	share, err := s.priceServiceRps.AddSubscriber(ctx, []string{position.ShareName})
	if err != nil {
		return 0, fmt.Errorf("AddSubscriber:%w", err)
	}

	currentTotal, PnL, err := calculateProfitAndLoss(ctx, position, share.SharePrice, balance.Balance)
	if err != nil {
		return 0, fmt.Errorf("calculateProfitAndLoss:%w", err)
	}
	updatedBalance := &model.Balance{
		BalanceID: balance.BalanceID,
		ProfileID: position.ProfileID,
		Balance:   currentTotal,
	}
	err = s.balanceRps.UpdateBalance(ctx, updatedBalance)
	if err != nil {
		return 0, fmt.Errorf("UpdateBalance:%w", err)
	}
	err = s.rps.DeletePosition(ctx, PositionID)
	if err != nil {
		return 0, fmt.Errorf("DeletePosition:%w", err)
	}

	return PnL, nil
}

// CheckForDestinationAmount function checks for given price of share in goven position
func (s *TradingService) CheckForShareClosePrice(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			logrus.Info("stream ended (ctx done)")
			return
		default:
			s.positionManager.Mu.RLock()
			for PositionID, OpenedPosition := range s.positionManager.OpenedPositions {
				if !s.positionManager.Closed[PositionID] {
					continue
				}
				position, err := s.rps.GetPositionByID(ctx, PositionID)
				if err != nil {
					logrus.Errorf("Error getting position: %v", err)
				}
				share, err := s.priceServiceRps.AddSubscriber(ctx, []string{position.ShareName})
				if err != nil {
					logrus.Errorf("Error getting share: %v", err)
				}
				fmt.Println(
					"Position id -> ", PositionID,
					"ShareOpenPrice -> ", OpenedPosition[PositionID].ShareOpenPrice,
					"CurrentSharePrice -> ", share.SharePrice,
					"ShareClosePrice -> ", OpenedPosition[PositionID].ShareClosePrice)
			}
			s.positionManager.Mu.RUnlock()
		}
	}
}

// PublishToAllOpenedPositions
func (s *TradingService) PublishToAllOpenedPositions(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			logrus.Info("stream ended (ctx done)")
		default:
			for PositionID, OpenedPosition := range s.positionManager.OpenedPositions {
				s.positionManager.Mu.Lock()
				position, err := s.rps.GetPositionByID(ctx, PositionID)
				if err != nil {
					logrus.Errorf("Error getting position: %v", err)
					return
				}
				share, err := s.priceServiceRps.AddSubscriber(ctx, []string{position.ShareName})
				if err != nil {
					logrus.Errorf("Error getting share: %v", err)
					return
				}
				fmt.Println("Position id -> ", PositionID, "ShareOpenPrice -> ", OpenedPosition[PositionID].ShareOpenPrice, "CurrentSharePrice -> ", share.SharePrice, "ShareClosePrice -> ", OpenedPosition[PositionID].ShareClosePrice)
				s.positionManager.Mu.Unlock()
			}
		}
	}
}

// CheckForTakeProfitAndStopLoss function checks if price reaches given stop loss or take profit
func (s *TradingService) CheckForTakeProfitAndStopLoss(ctx context.Context) {
	for {

		s.positionManager.Mu.RLock()
		for ID, openedPosition := range s.positionManager.OpenedPositions {
			select {
			case <-ctx.Done():
				logrus.Info("stream ended (ctx done)")
				return
			default:
				position, err := s.rps.GetPositionByID(ctx, ID)
				if err != nil {
					return
				}
				fmt.Println(position, openedPosition)

			}
		}
		s.positionManager.Mu.RUnlock()
	}
}

// CheckForServerUpdates function checks if server updates
func (s *TradingService) CheckForServerUpdates() {

}

// Calculations

// CalculateProfitAndLoss function calculates profit and loss for given position
func calculateProfitAndLoss(ctx context.Context, position *model.Position, currentSharePrice, balance float64) (float64, float64, error) {
	// initializing given float64 variables
	totalDecimal := decimal.NewFromFloatWithExponent(position.Total, -2)
	currentSharePriceDecimal := decimal.NewFromFloatWithExponent(currentSharePrice, -2)
	shareAmountDecimal := decimal.NewFromFloatWithExponent(position.ShareAmount, -4)
	balanceDecimal := decimal.NewFromFloatWithExponent(balance, -2)

	// calculating PnL
	currentTotalDecimal := currentSharePriceDecimal.Mul(shareAmountDecimal)
	pnl := currentTotalDecimal.Div(totalDecimal).Mul(decimal.New(100, -2))
	actualPnL := pnl.Sub(decimal.New(100, -2))
	actualPnL = actualPnL.Mul(decimal.New(100, 0))

	// converting back to float
	actualFloatPnL, _ := actualPnL.Float64()
	currentTotal, _ := currentTotalDecimal.Float64()
	fmt.Printf("old total: %f \nnew total: %f \nPnL: %f%% \n", position.Total, currentTotal, actualFloatPnL)

	// calculating of balance
	balanceDecimal = balanceDecimal.Add(currentTotalDecimal)
	balance, _ = balanceDecimal.Float64()
	return balance, actualFloatPnL, nil
}

// calculateAmountOfShares calculates the amount of shares for given amount of money
func calculateAmountOfShares(ctx context.Context, moneyAmount float64, sharePrice float64) (shareAmount float64, err error) {
	moneyAmountDecimal := decimal.NewFromFloatWithExponent(moneyAmount, -2)
	sharePriceDecimal := decimal.NewFromFloatWithExponent(sharePrice, -2)

	shareAmountDecimal := moneyAmountDecimal.Div(sharePriceDecimal)

	shareAmount, _ = shareAmountDecimal.Float64()
	return shareAmount, nil
}

// checkIfEnoughMoneyOnBalance function returns error if not enough money on balance
func checkIfEnoughMoneyOnBalance(balance *model.Balance, moneyAmount float64) (*model.Balance, error) {
	balanceDecimal := decimal.NewFromFloatWithExponent(balance.Balance, -2)
	moneyAmountDecimal := decimal.NewFromFloatWithExponent(moneyAmount, -2)
	greater := balanceDecimal.GreaterThanOrEqual(moneyAmountDecimal)
	if !greater {
		return nil, fmt.Errorf("GreaterThanOrEqual:%v", greater)
	}
	updatedBalanceDecimal := balanceDecimal.Sub(moneyAmountDecimal)

	balance.Balance, _ = updatedBalanceDecimal.Float64()

	return balance, nil
}
