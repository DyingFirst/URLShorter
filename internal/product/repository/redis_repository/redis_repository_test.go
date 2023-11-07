package redis_repository

import (
	"context"
	"testing"
	"time"

	"URLShorter/internal/product/repository/redis_repository/mocks"

	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
)

func TestRedisRepo_NewValue(t *testing.T) {
	mockClient := &mocks.MockRedisClient{}
	repo := NewRedisRepo(mockClient)

	t.Run("success", func(t *testing.T) {
		mockClient.SetFunc = func(ctx context.Context, key string, value interface{}, expiration time.Duration) *redis.StatusCmd {
			assert.Equal(t, "short", key)
			assert.Equal(t, "original", value)
			assert.Equal(t, time.Hour, expiration)
			return redis.NewStatusCmd(ctx, "OK", nil)
		}

		err := repo.NewValue(context.Background(), "short", "original")
		assert.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		mockClient.SetFunc = func(ctx context.Context, key string, value interface{}, expiration time.Duration) *redis.StatusCmd {
			cmd := redis.NewStatusCmd(ctx, "")
			cmd.SetErr(assert.AnError)
			return cmd
		}

		err := repo.NewValue(context.Background(), "short", "original")
		assert.Error(t, err)
	})
}

func TestRedisRepo_GetValue(t *testing.T) {
	mockClient := &mocks.MockRedisClient{}
	repo := NewRedisRepo(mockClient)

	t.Run("success", func(t *testing.T) {
		mockClient.GetFunc = func(ctx context.Context, key string) *redis.StringCmd {
			assert.Equal(t, "short", key)
			cmd := redis.NewStringResult("original", nil)
			return cmd
		}

		val, err := repo.GetValue(context.Background(), "short")
		assert.NoError(t, err)
		assert.Equal(t, "original", val)
	})

	t.Run("error", func(t *testing.T) {
		mockClient.GetFunc = func(ctx context.Context, key string) *redis.StringCmd {
			return redis.NewStringResult("", assert.AnError) // returns an error
		}

		_, err := repo.GetValue(context.Background(), "short")
		assert.Error(t, err)
	})
}
