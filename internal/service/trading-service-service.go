// Package service contains business-logic methods
package service

import (
	"context"
	"fmt"

	"github.com/eugenshima/trading-service/internal/model"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// TradingService struct represents a trading-service-service
type TradingService struct {
	rps             TradingRepository
	priceServiceRps PriceServiceRepository
	balanceRps      BalanceRepository
}

// NewTradingService creates a new TradingService
func NewTradingService(rps TradingRepository, priceServiceRps PriceServiceRepository, balanceRps BalanceRepository) *TradingService {
	return &TradingService{
		rps:             rps,
		priceServiceRps: priceServiceRps,
		balanceRps:      balanceRps,
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
}

// OpenPosition creates a position for a given ID
func (s *TradingService) OpenPosition(ctx context.Context, position *model.Position) error {
	balance, err := s.balanceRps.GetBalance(ctx, position.ProfileID)
	if err != nil {
		return fmt.Errorf("GetBalance: %w", err)
	}

	err = checkIfEnoughMoneyOnBalance(balance.Balance, position.SharePrice)
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

	return nil
}

// ClosePosition method closes the position of given ID
func (s *TradingService) ClosePosition(ctx context.Context, ID uuid.UUID) error {
	return s.rps.DeletePosition(ctx, ID)
}

// CheckForDestinationAmount function checks for given price of share in goven position
func (s *TradingService) CheckForDestinationAmount(ctx context.Context, ID uuid.UUID) {

}

// CalculateProfitAndLoss function calculates profit and loss for given position
func (s *TradingService) CalculateProfitAndLoss(ctx context.Context, ID uuid.UUID) (float64, error) {
	return 0, nil
}

// CheckForTakeProfitAndStopLoss function checks if price reaches given stop loss or take profit
func (s *TradingService) CheckForTakeProfitAndStopLoss(ctx context.Context, ID uuid.UUID) error {
	return nil
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
func checkIfEnoughMoneyOnBalance(balance, moneyAmount float64) error {

	balanceDecimal := decimal.NewFromFloatWithExponent(balance, -2)
	moneyAmountDecimal := decimal.NewFromFloatWithExponent(moneyAmount, -2)
	ok := balanceDecimal.GreaterThanOrEqual(moneyAmountDecimal)
	if !ok {
		return fmt.Errorf("GreaterThanOrEqual:%v", ok)
	}
	return nil
}
