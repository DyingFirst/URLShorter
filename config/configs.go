package config

import (
	"github.com/caarlos0/env/v9"
)

type Config struct {
	AppPort                string `env:"APP_PORT" envDefault:"8080"`
	RedisConnURL           string `env:"REDIS_CONN_URL"`
	PGConnURL              string `env:"PG_CONN_URL"`
	UseOnlyInMemoryStorage bool   `env:"USE_ONLY_IN_MEMORY_STORAGE" envDefault:"false"`
}

func NewConfig() (*Config, error) {
	cfg := Config{}
	err := env.Parse(&cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}
