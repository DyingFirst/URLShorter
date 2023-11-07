package config

import (
	"github.com/caarlos0/env/v9"
)

type Config struct {
	AppPort             string `env:"APP_PORT" envDefault:"8080"`
	GrpcPort            string `env:"GRPC_PORT" envDefault:"9090"`
	RedisConnURL        string `env:"REDIS_CONN_URL" envDefault:"localhost:6379"`
	PGConnURL           string `env:"PG_CONN_URL" envDefault:"postgres://designer:gh5vMBSuLThnaZWrKf@localhost:5432/shoredurls?sslmode=disable"`
	UseInMemoryStorage  bool   `env:"USE_MEMORY_STORAGE" envDefault:"true"`
	UseOutMemoryStorage bool   `env:"USE_OUT_MEMORY_STORAGE" envDefault:"true"`
}

func NewConfig() (*Config, error) {
	cfg := Config{}
	err := env.Parse(&cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}
