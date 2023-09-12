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

type TradingHandler struct {
	srv TradingService
	vl  *validator.Validate
	proto.UnimplementedTradingServiceServer
}

func NewTradingHandler(srv TradingService, vl *validator.Validate) *TradingHandler {
	return &TradingHandler{srv: srv, vl: vl}
}

type TradingService interface {
	OpenLongPosition(context.Context, *model.Position) error
	OpenShortPosition(context.Context, *model.Position) error
	ClosePosition(context.Context, uuid.UUID) error
}

func (h *TradingHandler) customValidator(ctx context.Context, i interface{}) error {
	if val, ok := i.(*model.Position); ok {
		err := h.vl.VarCtx(ctx, val.ID, "required")
		if err != nil {
			return fmt.Errorf("VarCtx: %w", err)
		}
		err = h.vl.VarCtx(ctx, val.PurchasePrice, "required")
		if err != nil {
			return fmt.Errorf("VarCtx: %w", err)
		}
		err = h.vl.VarCtx(ctx, val.SellingPrice, "required")
		if err != nil {
			return fmt.Errorf("VarCtx: %w", err)
		}
		err = h.vl.VarCtx(ctx, val.Share, "required")
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

func (h *TradingHandler) OpenPosition(ctx context.Context, req *proto.OpenPositionRequest) (*proto.OpenPositionResponse, error) {
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
	err = h.customValidator(ctx, position)
	if err != nil {
		logrus.WithFields(logrus.Fields{"position": position}).Errorf("customValidator: %v", err)
		return nil, fmt.Errorf("customValidator: %w", err)
	}
	if position.IsLong {
		err := h.srv.OpenLongPosition(ctx, position)
		if err != nil {
			logrus.WithFields(logrus.Fields{"position": position}).Errorf("OpenLongPosition: %v", err)
			return nil, fmt.Errorf("OpenLongPosition: %w", err)
		}
	} else {
		err := h.srv.OpenShortPosition(ctx, position)
		if err != nil {
			logrus.WithFields(logrus.Fields{"position": position}).Errorf("OpenShortPosition: %v", err)
			return nil, fmt.Errorf("OpenShortPosition: %w", err)
		}
	}

	return &proto.OpenPositionResponse{ID: ID.String()}, nil
}

func (h *TradingHandler) ClosePosition(ctx context.Context, req *proto.ClosePositionRequest) (*proto.ClosePositionResponse, error) {
	ID, err := uuid.Parse(req.ID)
	if err != nil {
		logrus.WithFields(logrus.Fields{"ID": req.ID}).Errorf("Parse: %v", err)
		return nil, fmt.Errorf("parse: %w", err)
	}
	err = h.srv.ClosePosition(ctx, ID)
	if err != nil {
		logrus.WithFields(logrus.Fields{"ID": ID}).Errorf("ClosePosition: %v", err)
		return nil, fmt.Errorf("ClosePosition: %w", err)
	}
	return &proto.ClosePositionResponse{}, nil
}
