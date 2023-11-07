package main

import (
	"URLShorter/config"
	"URLShorter/internal/server"
	"context"
	"errors"
	"sync"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
)

func main() {
	var rdb redis.UniversalClient
	var pdb *sqlx.DB
	rdb = nil
	pdb = nil
	logger := logrus.New()
	ctx := context.Background()

	cfg, err := config.NewConfig()
	if err != nil {
		logger.Fatal("Can't get config ", err)
	}

	if !cfg.UseInMemoryStorage && !cfg.UseOutMemoryStorage {
		logger.Fatal(errors.New("invalid storage"))
	}

	if cfg.UseOutMemoryStorage {
		pdb, err = sqlx.Open("postgres", cfg.PGConnURL)
		if err != nil {
			logger.Fatal("Can't connect to PostgresSQL", err)
		}
		if err = pdb.Ping(); err != nil {
			logger.Fatal("Can't ping db ", err)
		}
	}

	if cfg.UseInMemoryStorage {
		rdb = redis.NewUniversalClient(&redis.UniversalOptions{
			Addrs:    []string{cfg.RedisConnURL},
			Password: "",
			DB:       0,
		})
		if err := rdb.Ping(context.Background()); err.Err() != nil {
			logger.Fatal("Can't ping redis ", err)
		}
	}

	app := server.NewServer(ctx, cfg, logger, pdb, rdb)

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := app.RunHttp(); err != nil {
			logger.Fatal("HTTP server error: ", err)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := app.RunGRPC(); err != nil {
			logger.Fatal("gRPC server error: ", err)
		}
	}()

	wg.Wait()
}
