// Package config contains configuration information
package config

import "github.com/caarlos0/env"

// Config struct
type Config struct {
	PgxDBAddr             string `env:"PGXCONN" envDefault:"postgres://eugen:ur2qly1ini@localhost:5432/trading_db"`
	TradingServiceAddress string `env:"TRADING_SERVICE_ADDRESS" envDefault:"127.0.0.1:8083"`
}

// NewConfig creates a new Config instance
func NewConfig() (*Config, error) {
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}
