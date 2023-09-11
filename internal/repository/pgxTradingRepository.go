package repository

import (
	"context"
	"fmt"

	"github.com/eugenshima/trading-service/internal/model"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sirupsen/logrus"
)

type TradingRepository struct {
	pool *pgxpool.Pool
}

func NewTradingRepository(pool *pgxpool.Pool) *TradingRepository {
	return &TradingRepository{pool: pool}
}

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
	_, err = tx.Exec(ctx, "INSERT INTO trading.trading VALUES($1,$2,$3,$4,$5)", position.ID, position.IsLong, position.Share, position.PurchasePrice, position.SellingPrice)
	if err != nil {
		return fmt.Errorf("exec: %w", err)
	}
	return nil
}
