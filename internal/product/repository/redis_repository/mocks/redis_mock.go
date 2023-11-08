package mocks

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type MockRedisClient struct {
	redis.UniversalClient
	SetFunc  func(ctx context.Context, key string, value interface{}, expiration time.Duration) *redis.StatusCmd
	GetFunc  func(ctx context.Context, key string) *redis.StringCmd
	PingFunc func(ctx context.Context) *redis.StatusCmd
}

func (m *MockRedisClient) Set(
	ctx context.Context,
	key string,
	value interface{},
	expiration time.Duration,
) *redis.StatusCmd {
	return m.SetFunc(ctx, key, value, expiration)
}

func (m *MockRedisClient) Get(ctx context.Context, key string) *redis.StringCmd {
	return m.GetFunc(ctx, key)
}

func (m *MockRedisClient) Ping(ctx context.Context) *redis.StatusCmd {
	return m.PingFunc(ctx)
}
