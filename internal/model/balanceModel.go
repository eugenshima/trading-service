// Package model provides data Structures
package model

import "github.com/google/uuid"

// Balance struct represents the current balance
type Balance struct {
	ProfileID uuid.UUID `json:"profile_id"`
	Balance   float64   `json:"balance"`
}
