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
	balanceID, err := uuid.Parse(response.Balance.BalanceID)
	if err != nil {
		return nil, fmt.Errorf("GetUserByID: %w", err)
	}
	balance := &model.Balance{
		BalanceID: balanceID,
		ProfileID: profileID,
		Balance:   response.Balance.Balance,
	}
	return balance, nil
}

// UpdateBalance method updates a balance of given ID
func (r *BalanceRepository) UpdateBalance(ctx context.Context, balance *model.Balance) error {
	protoBalance := &proto.Balance{
		BalanceID: balance.BalanceID.String(),
		ProfileID: balance.ProfileID.String(),
		Balance:   balance.Balance,
	}
	_, err := r.client.UpdateUserBalance(ctx, &proto.UserUpdateRequest{Balance: protoBalance})
	if err != nil {
		return fmt.Errorf("UpdateUserBalance: %w", err)
	}
	return nil
}
