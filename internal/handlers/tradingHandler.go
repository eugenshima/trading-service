// Package handlers for the various types of events
package handlers

import (
	"context"
	"fmt"

	"github.com/eugenshima/trading-service/internal/model"
	proto "github.com/eugenshima/trading-service/proto"
	"github.com/go-playground/validator"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

// TradingHandler struct ....
type TradingHandler struct {
	srv TradingService
	vl  *validator.Validate
	proto.UnimplementedTradingServiceServer
}

// NewTradingHandler creates a new TradingHandler
func NewTradingHandler(srv TradingService, vl *validator.Validate) *TradingHandler {
	return &TradingHandler{srv: srv, vl: vl}
}

// TradingService interface represents the underlying TradingService
type TradingService interface {
	OpenPosition(context.Context, *model.Position) error
	ClosePosition(context.Context, uuid.UUID) (float64, error)
}

// customValidator function for validation of requests
func (h *TradingHandler) customValidator(ctx context.Context, i interface{}) error {
	if val, ok := i.(*model.Position); ok {
		err := h.vl.VarCtx(ctx, val.ID, "required")
		if err != nil {
			return fmt.Errorf("VarCtx: %w", err)
		}
		err = h.vl.VarCtx(ctx, val.ShareName, "required")
		if err != nil {
			return fmt.Errorf("VarCtx: %w", err)
		}
	} else {
		err := h.vl.VarCtx(ctx, i, "required")
		if err != nil {
			return fmt.Errorf("VarCtx: %w", err)
		}
	}
	return nil
}

// OpenPosition function opens position for user
func (h *TradingHandler) OpenPosition(ctx context.Context, req *proto.OpenPositionRequest) (*proto.OpenPositionResponse, error) {
	ID, err := uuid.Parse(req.Position.Id)
	if err != nil {
		logrus.WithFields(logrus.Fields{"ID": req.Position.Id}).Errorf("Parse: %v", err)
		return nil, fmt.Errorf("parse: %w", err)
	}
	position := &model.Position{
		ID:          uuid.New(),
		ProfileID:   ID,
		IsLong:      req.Position.IsLong,
		ShareName:   req.Position.ShareName,
		Total:       req.Position.Total,
		ShareAmount: req.Position.ShareAmount,
		StopLoss:    req.Position.StopLoss,
		TakeProfit:  req.Position.TakeProfit,
	}
	err = h.customValidator(ctx, position)
	if err != nil {
		logrus.WithFields(logrus.Fields{"position": position}).Errorf("customValidator: %v", err)
		return nil, fmt.Errorf("customValidator: %w", err)
	}
	err = h.srv.OpenPosition(ctx, position)
	if err != nil {
		logrus.WithFields(logrus.Fields{"position": position}).Errorf("OpenPosition: %v", err)
		return nil, fmt.Errorf("OpenPosition: %w", err)
	}

	return &proto.OpenPositionResponse{ID: position.ID.String()}, nil
}

// ClosePosition function closes position for user
func (h *TradingHandler) ClosePosition(ctx context.Context, req *proto.ClosePositionRequest) (*proto.ClosePositionResponse, error) {
	ID, err := uuid.Parse(req.ID)
	if err != nil {
		logrus.WithFields(logrus.Fields{"ID": req.ID}).Errorf("Parse: %v", err)
		return nil, fmt.Errorf("parse: %w", err)
	}
	profitAndLoss, err := h.srv.ClosePosition(ctx, ID)
	if err != nil {
		logrus.WithFields(logrus.Fields{"ID": ID}).Errorf("ClosePosition: %v", err)
		return nil, fmt.Errorf("ClosePosition: %w", err)
	}
	return &proto.ClosePositionResponse{PnL: profitAndLoss}, nil
}
