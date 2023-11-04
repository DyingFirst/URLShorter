package redis

import (
	"context"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
)

func InitRedis(address string, log *logrus.Logger) (rdb *redis.Client) {
	db := redis.NewClient(&redis.Options{
		Addr:     address,
		Password: "",
		DB:       0,
	})

	if err := db.Ping(context.Background()); err.Err() != nil {
		log.WithField("error", err).Fatal("Can't connect to Redis")
		return nil
	}
	return db
}
