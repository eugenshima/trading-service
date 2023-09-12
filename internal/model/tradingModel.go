// Package model provides data Structures
package model

import (
	"sync"

	"github.com/google/uuid"
)

// Position struct represents an position by user
type Position struct {
	ID            uuid.UUID `json:"id"`
	ProfileID     uuid.UUID `json:"profile_id"`
	IsLong        bool      `json:"is_long"`
	Share         string    `json:"share"`
	PurchasePrice float64   `json:"purch_price"`
	SellingPrice  float64   `json:"sell_price"`
	StopLoss      float64   `json:"stop_loss"`
	TakeProfit    float64   `json:"take_profit"`
}

// PubSub struct represents a model for subscriptions(subscriber part)
type PubSub struct {
	Mu   sync.RWMutex
	Subs map[uuid.UUID][]string
	// SubsShares map[uuid.UUID]chan []*Share
	Closed map[uuid.UUID]bool
}
