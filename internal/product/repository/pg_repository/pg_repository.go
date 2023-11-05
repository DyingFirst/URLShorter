package pg_repository

import (
	"URLShorter/internal/product"
	"context"
	"github.com/jackc/pgx/v5"
	"time"
)

type PGRepo struct {
	db *pgx.Conn
}

func NewPGRepo(db *pgx.Conn) product.PGRepo {
	return &PGRepo{db: db}
}

func (pg *PGRepo) NewValue(ctx context.Context, ShortURL string, OriginalURL string) error {
	Now := time.Now()
	UpdateDate := Now.Add(time.Hour)
	_, err := pg.db.Exec(ctx, querySetOriginalUrlByID, OriginalURL, ShortURL, UpdateDate)
	return err
}
func (pg *PGRepo) GetValue(ctx context.Context, ShortURL string) (string, error) {
	value, err := pg.db.Exec(ctx, queryGetOriginalUrlByID, ShortURL)
	return value.String(), err
}
