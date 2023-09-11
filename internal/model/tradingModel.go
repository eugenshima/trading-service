package model

import "github.com/google/uuid"

type Position struct {
	ID            uuid.UUID `json:"id"`
	IsLong        bool      `json:"is_long"`
	Share         string    `json:"share"`
	PurchasePrice float64   `json:"purch_price"`
	SellingPrice  float64   `json:"sell_price"`
}
