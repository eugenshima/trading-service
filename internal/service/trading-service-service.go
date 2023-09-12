// Package service contains business-logic methods
package service

import (
	"context"
	"fmt"

	"github.com/eugenshima/trading-service/internal/model"

	"github.com/google/uuid"
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
}

// PriceServiceRepository interface represents a price-service-repository methods
type PriceServiceRepository interface {
	AddSubscriber(context.Context, []string) error
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
	if position.IsLong {
		fmt.Println("Opening long position. current balance -> ", balance.Balance)
	} else {
		fmt.Println("Opening short position")
	}
	selectedShares := []string{}
	selectedShares = append(selectedShares, position.Share)
	err = s.priceServiceRps.AddSubscriber(ctx, selectedShares)
	if err != nil {
		return fmt.Errorf("AddSubscriber:%w", err)
	}

	return s.rps.CreatePosition(ctx, position)
}

// ClosePosition method closes the position of given ID
func (s *TradingService) ClosePosition(ctx context.Context, ID uuid.UUID) error {
	return s.rps.DeletePosition(ctx, ID)
}
