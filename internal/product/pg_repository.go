package product

import "context"

type PGRepo interface {
	NewValue(ctx context.Context, OriginalURL string, ShortURL string) error
	GetValue(ctx context.Context, ShortURL string) (string, error)
}
