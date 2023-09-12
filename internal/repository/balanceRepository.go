// Package repository contains methods to communicate with postgres and gRPC servers
package repository

import (
	"context"
	"fmt"

	proto "github.com/eugenshima/balance/proto"
	"github.com/eugenshima/trading-service/internal/model"
	"github.com/google/uuid"
)

// BalanceRepository represents a repository that contains balance microservice methods
type BalanceRepository struct {
	client proto.BalanceServiceClient
}

// NewBalanceRepository creates a new BalanceRepository
func NewBalanceRepository(client proto.BalanceServiceClient) *BalanceRepository {
	return &BalanceRepository{client: client}
}

// GetBalance method returns a balance by given ID
func (r *BalanceRepository) GetBalance(ctx context.Context, ID uuid.UUID) (*model.Balance, error) {
	response, err := r.client.GetUserByID(ctx, &proto.UserGetByIDRequest{ProfileID: ID.String()})
	if err != nil {
		return nil, fmt.Errorf("GetUserByID: %w", err)
	}
	profileID, err := uuid.Parse(response.Balance.ProfileID)
	if err != nil {
		return nil, fmt.Errorf("GetUserByID: %w", err)
	}
	balance := &model.Balance{
		ProfileID: profileID,
		Balance:   response.Balance.Balance,
	}
	return balance, nil
}

// UpdateBalance ....
func (r *BalanceRepository) UpdateBalance(_ context.Context, _ uuid.UUID) error {
	return nil
}
