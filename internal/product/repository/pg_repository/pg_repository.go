package pg_repository

import (
	"context"
	"time"

	"URLShorter/internal/product"

	"github.com/jmoiron/sqlx"
)

type PGRepo struct {
	db *sqlx.DB
}

func NewPGRepo(dbConn *sqlx.DB) product.PGRepo {
	return &PGRepo{db: dbConn}
}

func (pg *PGRepo) NewValue(ctx context.Context, ShortURL string, OriginalURL string) error {
	Now := time.Now()
	UpdateDate := Now.Add(time.Hour)
	_, err := pg.db.ExecContext(ctx, querySetOriginalUrlByID, OriginalURL, ShortURL, UpdateDate)
	return err
}

func (pg *PGRepo) GetValue(ctx context.Context, ShortURL string) (string, error) {
	var value string
	err := pg.db.GetContext(ctx, &value, queryGetOriginalUrlByID, ShortURL)
	return value, err
}
