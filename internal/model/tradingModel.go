// Package model provides data Structures
package model

import (
	"sync"

	"github.com/google/uuid"
)

// Position struct represents an position by user
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

type Share struct {
	ShareName  string  `json:"share_name"`
	SharePrice float64 `json:"share_price"`
}

// PubSub struct represents a model for subscriptions(subscriber part)
type PubSub struct {
	Mu   sync.RWMutex
	Subs map[uuid.UUID][]string
	// SubsShares map[uuid.UUID]chan []*Share
	Closed map[uuid.UUID]bool
}
