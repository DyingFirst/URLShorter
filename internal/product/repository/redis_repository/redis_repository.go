package redis_repository

import (
	"URLShorter/internal/product"
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

type RedisRepo struct {
	rdb *redis.Client
}

func NewRedisRepo(rdb *redis.Client) product.RedisRepo {
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
