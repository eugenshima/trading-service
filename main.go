// Package main is the entry-point for the package
package main

import (
	"context"
	"fmt"
	"net"

	balanceServiceProto "github.com/eugenshima/balance/proto"
	priceServiceProto "github.com/eugenshima/price-service/proto"
	"github.com/eugenshima/trading-service/internal/config"
	"github.com/eugenshima/trading-service/internal/handlers"
	"github.com/eugenshima/trading-service/internal/model"
	"github.com/eugenshima/trading-service/internal/repository"
	"github.com/eugenshima/trading-service/internal/service"
	readingServiceProto "github.com/eugenshima/trading-service/proto"

	"github.com/go-playground/validator"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

// NewDBPsql function provides Connection with PostgreSQL database
func NewDBPsql(env string) (*pgxpool.Pool, error) {
	// Initialization a connect configuration for a PostgreSQL using pgx driver
	cfg, err := pgxpool.ParseConfig(env)
	if err != nil {
		return nil, fmt.Errorf("error connection to PostgreSQL: %v", err)
	}

	// Establishing a new connection to a PostgreSQL database using the pgx driver
	pool, err := pgxpool.ConnectConfig(context.Background(), cfg)
	if err != nil {
		return nil, fmt.Errorf("error connection to PostgreSQL: %v", err)
	}
	// Output to console
	fmt.Println("Connected to PostgreSQL!")

	return pool, nil
}

// nolint:staticcheck // noinspection
func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		fmt.Printf("Error extracting env variables: %v", err)
		return
	}
	pool, err := NewDBPsql(cfg.PgxDBAddr)
	if err != nil {
		logrus.WithFields(logrus.Fields{"PgxDBAddr: ": cfg.PgxDBAddr}).Errorf("NewDBPsql: %v", err)
		return
	}

	priceServiceConn, err := grpc.Dial(":8080", grpc.WithInsecure())
	if err != nil {
		return
	}
	defer func() {
		err = priceServiceConn.Close()
		if err != nil {
			fmt.Println("error closing price service connection")
		}
	}()

	balanceConn, err := grpc.Dial(":8081", grpc.WithInsecure())
	if err != nil {
		return
	}
	defer func() {
		err = balanceConn.Close()
		if err != nil {
			fmt.Println("error closing balance connection")
		}
	}()

	priceServiceClient := priceServiceProto.NewPriceServiceClient(priceServiceConn)
	balanceServiceClient := balanceServiceProto.NewBalanceServiceClient(balanceConn)

	rps := repository.NewTradingRepository(pool)
	priceServiceRps := repository.NewPriceServiceClient(priceServiceClient)
	balanceServiceRps := repository.NewBalanceRepository(balanceServiceClient)

	positionManager := model.NewPositionManager()

	srv := service.NewTradingService(rps, priceServiceRps, balanceServiceRps, positionManager)
	handler := handlers.NewTradingHandler(srv, validator.New())

	// go srv.CheckForTakeProfitAndStopLoss(context.TODO())

	lis, err := net.Listen("tcp", cfg.TradingServiceAddress)
	if err != nil {
		logrus.Fatalf("cannot create listener: %s", err)
	}

	serverRegistrar := grpc.NewServer()
	readingServiceProto.RegisterTradingServiceServer(serverRegistrar, handler)
	err = serverRegistrar.Serve(lis)
	if err != nil {
		logrus.Fatalf("cannot start server: %s", err)
	}
}
