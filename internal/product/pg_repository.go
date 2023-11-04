package product

import "context"

type PGRepo interface {
	NewValue(ctx context.Context, ShortURL string, OriginalURL string) error
	GetValue(ctx context.Context, ShortURL string) (string, error)
}
