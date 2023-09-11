package service

import (
	"context"
	"fmt"

	"github.com/eugenshima/trading-service/internal/model"
)

type TradingService struct {
	rps TradingRepository
}

func NewTradingService(rps TradingRepository) *TradingService {
	return &TradingService{rps: rps}
}

type TradingRepository interface {
	CreatePosition(context.Context, *model.Position) error
}

func (s *TradingService) OpenLongPosition(ctx context.Context, position *model.Position) error {
	fmt.Println("Opening long position")
	return s.rps.CreatePosition(ctx, position)
}

func (s *TradingService) OpenShortPosition(ctx context.Context, position *model.Position) error {
	fmt.Println("Opening short position")
	return s.rps.CreatePosition(ctx, position)
}
