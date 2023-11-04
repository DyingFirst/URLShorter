package product

import "context"

type RedisRepo interface {
	NewValue(ctx context.Context, ShortURL string, OriginalURL string) error
	GetValue(ctx context.Context, ShortURL string) (string, error)
}
