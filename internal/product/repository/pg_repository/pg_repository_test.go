package pg_repository

import (
	"context"
	"errors"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewValue(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	pgRepo := NewPGRepo(sqlxDB)

	t.Run("success", func(t *testing.T) {
		mock.ExpectExec(regexp.QuoteMeta(querySetOriginalUrlByID)).
			WithArgs("OriginalURL", "ShortURL", sqlmock.AnyArg()).
			WillReturnResult(sqlmock.NewResult(1, 1))

		err := pgRepo.NewValue(context.TODO(), "ShortURL", "OriginalURL")
		assert.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		mock.ExpectExec(regexp.QuoteMeta(querySetOriginalUrlByID)).
			WithArgs("OriginalURL", "ShortURL", sqlmock.AnyArg()).
			WillReturnError(errors.New("exec error"))

		err := pgRepo.NewValue(context.TODO(), "ShortURL", "OriginalURL")
		assert.Error(t, err)
	})
}

func TestGetValue(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	pgRepo := NewPGRepo(sqlxDB)

	t.Run("success", func(t *testing.T) {
		mock.ExpectQuery(regexp.QuoteMeta(queryGetOriginalUrlByID)).
			WithArgs("ShortURL").
			WillReturnRows(sqlmock.NewRows([]string{"value"}).AddRow("OriginalURL"))

		value, err := pgRepo.GetValue(context.TODO(), "ShortURL")
		assert.NoError(t, err)
		assert.Equal(t, "OriginalURL", value)
	})

	t.Run("error", func(t *testing.T) {
		mock.ExpectQuery(regexp.QuoteMeta(queryGetOriginalUrlByID)).
			WithArgs("ShortURL").
			WillReturnError(errors.New("exec error"))

		value, err := pgRepo.GetValue(context.TODO(), "ShortURL")
		assert.Error(t, err)
		assert.Empty(t, value)
	})
}
