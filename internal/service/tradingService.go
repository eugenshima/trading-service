package service

import (
	"context"

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
	return s.rps.CreatePosition(ctx, position)
}
