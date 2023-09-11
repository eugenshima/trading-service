package handlers

import (
	"context"
	"fmt"

	"github.com/eugenshima/trading-service/internal/model"
	proto "github.com/eugenshima/trading-service/proto"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type TradingHandler struct {
	srv TradingService
	proto.UnimplementedPriceServiceServer
}

func NewTradingHandler(srv TradingService) *TradingHandler {
	return &TradingHandler{srv: srv}
}

type TradingService interface {
	OpenLongPosition(context.Context, *model.Position) error
}

func (h *TradingHandler) customValidator(ctx context.Context, i interface{}) error {
	return nil
}

func (h *TradingHandler) OpenPosition(ctx context.Context, req *proto.OpenPositionRequest) (*proto.OpenPositionResponse, error) {
	h.customValidator(ctx, req)
	ID, err := uuid.Parse(req.Position.Id)
	if err != nil {
		logrus.WithFields(logrus.Fields{"ID": req.Position.Id}).Errorf("Parse: %v", err)
		return nil, fmt.Errorf("parse: %w", err)
	}
	position := &model.Position{
		ID:            ID,
		IsLong:        req.Position.IsLong,
		Share:         req.Position.Share,
		PurchasePrice: req.Position.PurchasePrice,
		SellingPrice:  req.Position.SellingPrice,
	}
	h.srv.OpenLongPosition(ctx, position)
	return &proto.OpenPositionResponse{ID: ID.String()}, nil
}
