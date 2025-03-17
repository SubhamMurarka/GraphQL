package config

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

func NewRedisDatabase() (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", Config.RedisHost, Config.RedisPort),
		Password: "",
		DB:       0,
		PoolSize: 5,
	})

	ctx := context.Background()
	if err := client.Ping(ctx).Err(); err != nil {
		return nil, err
	}

	return client, nil
}
