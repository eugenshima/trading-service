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

// Share struct represents one share
type Share struct {
	ShareName  string  `json:"share_name"`
	SharePrice float64 `json:"share_price"`
}

// PubSub struct represents a model for subscriptions(subscriber part)
type PositionManager struct {
	Mu              sync.RWMutex
	OpenedPositions map[uuid.UUID]chan *Position
	Closed          map[uuid.UUID]bool
}

// NewPositionManager creates a new position manager
func NewPositionManager() *PositionManager {
	return &PositionManager{
		OpenedPositions: make(map[uuid.UUID]chan *Position),
		Closed:          make(map[uuid.UUID]bool),
	}
}
