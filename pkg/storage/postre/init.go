package postre

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/sirupsen/logrus"
)

func InitPsqlDB(connectionURL string, log *logrus.Logger) (pdb *pgx.Conn) {
	db, err := pgx.Connect(context.Background(), connectionURL)
	if err != nil {
		log.Fatal("Can't connect to PostgresSQL", err)
		return nil
	}
	if err = db.Ping(context.Background()); err != nil {
		return nil
	}
	return db
}
