// Package repository contains methods to communicate with postgres and gRPC servers
package repository

import (
	"context"
	"fmt"

	"github.com/eugenshima/trading-service/internal/model"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sirupsen/logrus"
)

// TradingRepository structure ....
type TradingRepository struct {
	pool *pgxpool.Pool
}

// NewTradingRepository creates a new TradingRepository
func NewTradingRepository(pool *pgxpool.Pool) *TradingRepository {
	return &TradingRepository{pool: pool}
}

// CreatePosition method creates a new Position
func (repo *TradingRepository) CreatePosition(ctx context.Context, position *model.Position) error {
	tx, err := repo.pool.BeginTx(ctx, pgx.TxOptions{IsoLevel: "repeatable read"})
	if err != nil {
		return fmt.Errorf("BeginTx: %w", err)
	}
	defer func() {
		if err != nil {
			err = tx.Rollback(ctx)
			if err != nil {
				logrus.Errorf("Rollback: %v", err)
				return
			}
		} else {
			err = tx.Commit(ctx)
			if err != nil {
				logrus.Errorf("Commit: %v", err)
				return
			}
		}
	}()
	_, err = tx.Exec(
		ctx,
		"INSERT INTO trading.trading VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9)",
		position.ID, position.ProfileID, position.IsLong, position.ShareName, position.SharePrice, position.Total, position.ShareAmount, position.StopLoss, position.TakeProfit)
	if err != nil {
		return fmt.Errorf("exec: %w", err)
	}
	return nil
}

// DeletePosition method deletes a position from database
func (repo *TradingRepository) DeletePosition(ctx context.Context, ID uuid.UUID) error {
	tx, err := repo.pool.BeginTx(ctx, pgx.TxOptions{IsoLevel: "repeatable read"})
	if err != nil {
		return fmt.Errorf("BeginTx: %w", err)
	}
	defer func() {
		if err != nil {
			err = tx.Rollback(ctx)
			if err != nil {
				logrus.Errorf("Rollback: %v", err)
				return
			}
		} else {
			err = tx.Commit(ctx)
			if err != nil {
				logrus.Errorf("Commit: %v", err)
				return
			}
		}
	}()
	_, err = tx.Exec(ctx, "DELETE FROM trading.trading WHERE id=$1", ID)
	if err != nil {
		return fmt.Errorf("exec: %w", err)
	}
	return nil
}

// GetPositionByID functions returns the position of the given ID from database
func (repo *TradingRepository) GetPositionByID(ctx context.Context, PositionID uuid.UUID) (*model.Position, error) {
	tx, err := repo.pool.BeginTx(ctx, pgx.TxOptions{IsoLevel: "repeatable read"})
	if err != nil {
		return nil, fmt.Errorf("BeginTx: %w", err)
	}
	defer func() {
		if err != nil {
			err = tx.Rollback(ctx)
			if err != nil {
				logrus.Errorf("Rollback: %v", err)
				return
			}
		} else {
			err = tx.Commit(ctx)
			if err != nil {
				logrus.Errorf("Commit: %v", err)
				return
			}
		}
	}()
	position := &model.Position{}
	err = tx.QueryRow(ctx, "SELECT id, profile_id, is_long, share_name, share_price, total, shares_amount, stop_loss, take_profit FROM trading.trading WHERE id=$1", PositionID).Scan(&position.ID, &position.ProfileID, &position.IsLong, &position.ShareName, &position.SharePrice, &position.Total, &position.ShareAmount, &position.StopLoss, &position.TakeProfit)
	if err != nil {
		return nil, fmt.Errorf("QueryRow: %w", err)
	}
	return position, nil
}

// GetAllIDsPositions functions returns all positions of the given user
func (repo *TradingRepository) GetAllIDsPositions(ctx context.Context, profileID uuid.UUID) ([]*model.Position, error) {
	tx, err := repo.pool.BeginTx(ctx, pgx.TxOptions{IsoLevel: "repeatable read"})
	if err != nil {
		return nil, fmt.Errorf("BeginTx: %w", err)
	}
	defer func() {
		if err != nil {
			err = tx.Rollback(ctx)
			if err != nil {
				logrus.Errorf("Rollback: %v", err)
				return
			}
		} else {
			err = tx.Commit(ctx)
			if err != nil {
				logrus.Errorf("Commit: %v", err)
				return
			}
		}
	}()
	rows, err := tx.Query(ctx, "SELECT id, profile_id, is_long, share_name, share_price, total, shares_amount, stop_loss, take_profit FROM trading.trading")
	if err != nil {
		return nil, fmt.Errorf("Query(): %w", err)
	}
	defer rows.Close()

	var positions []*model.Position

	for rows.Next() {
		position := &model.Position{}
		err := rows.Scan(&position.ID, &position.ProfileID, &position.IsLong, &position.ShareName, &position.SharePrice, &position.Total, &position.ShareAmount, &position.StopLoss, &position.TakeProfit)
		if err != nil {
			return nil, fmt.Errorf("Scan(): %w", err)
		}
		positions = append(positions, position)
	}
	return positions, rows.Err()
}
