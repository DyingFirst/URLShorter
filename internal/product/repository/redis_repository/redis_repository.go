package redis_repository

import (
	"context"
	"time"

	"URLShorter/internal/product"

	"github.com/redis/go-redis/v9"
)

type RedisRepo struct {
	rdb redis.UniversalClient
}

func NewRedisRepo(rdb redis.UniversalClient) product.RedisRepo {
	return &RedisRepo{rdb: rdb}
}

func (r *RedisRepo) NewValue(ctx context.Context, ShortURL string, OriginalURL string) error {
	err := r.rdb.Set(ctx, ShortURL, OriginalURL, time.Hour).Err()
	return err
}

func (r *RedisRepo) GetValue(ctx context.Context, ShortURL string) (string, error) {
	value, err := r.rdb.Get(ctx, ShortURL).Result()
	return value, err
}
