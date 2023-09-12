package repository

import (
	"context"
	"fmt"

	proto "github.com/eugenshima/balance/proto"
	"github.com/eugenshima/trading-service/internal/model"
	"github.com/google/uuid"
)

type BalanceRepository struct {
	client proto.BalanceServiceClient
}

func NewBalanceRepository(client proto.BalanceServiceClient) *BalanceRepository {
	return &BalanceRepository{client: client}
}

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

func (r *BalanceRepository) UpdateBalance(ctx context.Context, ID uuid.UUID) error {

	//response, err := r.client.UpdateUserBalance(ctx, &proto.UserUpdateRequest{})
	return nil
}
