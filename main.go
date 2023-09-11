package main

import (
	"context"
	"fmt"
	"net"

	"github.com/eugenshima/trading-service/internal/config"
	"github.com/eugenshima/trading-service/internal/handlers"
	"github.com/eugenshima/trading-service/internal/repository"
	"github.com/eugenshima/trading-service/internal/service"
	proto "github.com/eugenshima/trading-service/proto"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

// NewDBPsql function provides Connection with PostgreSQL database
func NewDBPsql(env string) (*pgxpool.Pool, error) {
	// Initialization a connect configuration for a PostgreSQL using pgx driver
	config, err := pgxpool.ParseConfig(env)
	if err != nil {
		return nil, fmt.Errorf("error connection to PostgreSQL: %v", err)
	}

	// Establishing a new connection to a PostgreSQL database using the pgx driver
	pool, err := pgxpool.ConnectConfig(context.Background(), config)
	if err != nil {
		return nil, fmt.Errorf("error connection to PostgreSQL: %v", err)
	}
	// Output to console
	fmt.Println("Connected to PostgreSQL!")

	return pool, nil
}

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
	rps := repository.NewTradingRepository(pool)
	srv := service.NewTradingService(rps)
	handler := handlers.NewTradingHandler(srv)

	lis, err := net.Listen("tcp", "127.0.0.1:8083")
	if err != nil {
		logrus.Fatalf("cannot create listener: %s", err)
	}

	serverRegistrar := grpc.NewServer()
	proto.RegisterPriceServiceServer(serverRegistrar, handler)
	err = serverRegistrar.Serve(lis)
	if err != nil {
		logrus.Fatalf("cannot start server: %s", err)
	}
}
