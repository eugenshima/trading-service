package model

import "github.com/google/uuid"

type Balance struct {
	ProfileID uuid.UUID `json:"profile_id"`
	Balance   float64   `json:"balance"`
}
