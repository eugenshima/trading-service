// Package service contains business-logic methods
package service

import (
	"context"
	"fmt"
	"time"

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
func (s *TradingService) addPositionToMap(ID uuid.UUID) error {
	s.positionManager.Mu.Lock()
	defer s.positionManager.Mu.Unlock()
	if _, ok := s.positionManager.OpenedPositions[ID]; !ok {
		s.positionManager.OpenedPositions[ID] = make(chan *model.Position, 1)
		return nil
	}
	return fmt.Errorf("error opening position on ID: %v", ID)
}

// deletePositionFromMap method deletes a position from position manager
func (s *TradingService) deletePositionFromMap(ID uuid.UUID) error {
	s.positionManager.Mu.Lock()
	defer s.positionManager.Mu.Unlock()
	if _, ok := s.positionManager.OpenedPositions[ID]; !ok {
		delete(s.positionManager.Closed, ID)
		close(s.positionManager.OpenedPositions[ID])
		return nil
	}
	return fmt.Errorf("error closing position on ID: %v", ID)
}

// OpenPosition creates a position for a given ID with checking all the necessary conditions
func (s *TradingService) OpenPosition(ctx context.Context, position *model.Position) error {
	err := s.addPositionToMap(position.ID)
	if err != nil {
		return fmt.Errorf("addPositionToMap: %w", err)
	}

	balance, err := s.balanceRps.GetBalance(ctx, position.ProfileID)
	if err != nil {
		return fmt.Errorf("GetBalance: %w", err)
	}

	updatedBalance, err := checkIfEnoughMoneyOnBalance(balance, position.Total)
	if err != nil {
		return fmt.Errorf("not enough money on you balance: %w", err)
	}

	selectedShares := []string{}
	selectedShares = append(selectedShares, position.ShareName)
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
func (s *TradingService) ClosePosition(ctx context.Context, ID uuid.UUID) (float64, error) {
	balance, err := s.balanceRps.GetBalance(ctx, ID)
	if err != nil {
		return 0, fmt.Errorf("GetBalance: %w", err)
	}
	position, err := s.rps.GetPositionByID(ctx, ID)
	if err != nil {
		return 0, fmt.Errorf("GetPositionByID: %w", err)
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
		ProfileID: ID,
		Balance:   currentTotal,
	}
	err = s.balanceRps.UpdateBalance(ctx, updatedBalance)
	if err != nil {
		return 0, fmt.Errorf("UpdateBalance:%w", err)
	}
	err = s.rps.DeletePosition(ctx, ID)
	if err != nil {
		return 0, fmt.Errorf("DeletePosition:%w", err)
	}
	return PnL, nil
}

// CheckForDestinationAmount function checks for given price of share in goven position
func (s *TradingService) CheckForDestinationAmount(ctx context.Context, ID uuid.UUID) {

}

// CheckForTakeProfitAndStopLoss function checks if price reaches given stop loss or take profit
func (s *TradingService) CheckForTakeProfitAndStopLoss(ctx context.Context) {
	for {
		time.Sleep(1 * time.Second)
		s.positionManager.Mu.Lock()
		for ID, _ := range s.positionManager.OpenedPositions {
			select {
			case <-ctx.Done():
				logrus.Info("stream ended (ctx done)")
				return
			default:
				logrus.Info()
				position, err := s.rps.GetPositionByID(ctx, ID)
				if err != nil {
					return
				}
				s.positionManager.OpenedPositions[ID] <- position
			}
		}
		s.positionManager.Mu.Unlock()
	}
}

// Calculations

// CalculateProfitAndLoss function calculates profit and loss for given position
func calculateProfitAndLoss(ctx context.Context, position *model.Position, currentSharePrice, balance float64) (float64, float64, error) {
	// initializing given float64 variables
	totalDecimal := decimal.NewFromFloatWithExponent(position.Total, -2)
	currentSharePriceDecimal := decimal.NewFromFloatWithExponent(currentSharePrice, -2)
	shareAmountDecimal := decimal.NewFromFloatWithExponent(position.ShareAmount, -4)
	balanceDecimal := decimal.NewFromFloatWithExponent(balance, -2)

	// calculating of PnL
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
	ok := balanceDecimal.GreaterThanOrEqual(moneyAmountDecimal)
	if !ok {
		return nil, fmt.Errorf("GreaterThanOrEqual:%v", ok)
	}
	updatedBalanceDecimal := balanceDecimal.Sub(moneyAmountDecimal)

	balance.Balance, _ = updatedBalanceDecimal.Float64()

	return balance, nil
}
