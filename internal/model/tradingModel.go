// Package model provides data Structures
package model

import (
	"sync"

	"github.com/google/uuid"
)

// Position struct represents an user's position
type Position struct {
	ID          uuid.UUID `json:"id"`
	ProfileID   uuid.UUID `json:"profile_id"`
	IsLong      bool      `json:"is_long"`
	ShareName   string    `json:"share_name"`
	SharePrice  float64   `json:"share_price"`
	Total       float64   `json:"total"`
	ShareAmount float64   `json:"share_amount"`
	StopLoss    float64   `json:"stop_loss"`
	TakeProfit  float64   `json:"take_profit"`
}

type OpenedPosition struct {
	PositionID      uuid.UUID `json:"position_id"`
	ShareOpenPrice  float64   `json:"share_open_price"`
	ShareClosePrice float64   `json:"share_close_price"`
	ShareAmount     float64   `json:"share_amount"`
	IsOpened        bool      `json:"is_closed"`
}

// Share struct represents one share
type Share struct {
	ShareName  string  `json:"share_name"`
	SharePrice float64 `json:"share_price"`
}

// PubSub struct represents a model for subscriptions(subscriber part)
type PositionManager struct {
	Mu              sync.RWMutex
	OpenedPositions map[uuid.UUID]map[uuid.UUID]*OpenedPosition
	Closed          map[uuid.UUID]bool
}

// NewPositionManager creates a new position manager
func NewPositionManager() *PositionManager {
	return &PositionManager{
		OpenedPositions: make(map[uuid.UUID]map[uuid.UUID]*OpenedPosition),
		Closed:          make(map[uuid.UUID]bool),
	}
}
