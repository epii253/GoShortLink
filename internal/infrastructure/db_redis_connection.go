package infrastructure

import (
	"context"
	settings "project/internal/settings"

	"github.com/redis/go-redis/v9"
)

func NewRedisConnection(cnf *settings.Config) (*redis.Client, error) {
	connection := redis.NewClient(&redis.Options{
		Addr: cnf.RedisAddr,
		Password: cnf.RedisPass,
	})

	if err := connection.Ping(context.Background()).Err(); err != nil {
		return nil, err
	}

	return connection, nil
}