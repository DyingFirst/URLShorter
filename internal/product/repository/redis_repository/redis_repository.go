package redis_repository

import (
	"URLShorter/internal/product"
	"context"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
	"time"
)

type RedisRepo struct {
	logger *logrus.Logger
	rdb    *redis.Client
}

func NewRedisRepo(logger *logrus.Logger, rdb *redis.Client) product.RedisRepo {
	return &RedisRepo{logger: logger, rdb: rdb}
}

func (r *RedisRepo) NewValue(ctx context.Context, ShortURL string, OriginalURL string) error {
	err := r.rdb.Set(ctx, ShortURL, OriginalURL, time.Hour).Err()
	return err
}

func (r *RedisRepo) GetValue(ctx context.Context, ShortURL string) (string, error) {
	value, err := r.rdb.Get(ctx, ShortURL).Result()
	return value, err
}
