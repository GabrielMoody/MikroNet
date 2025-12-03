package model

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
)

func RedisConnect() *redis.Client {
	ctx := context.Background()

	rdb := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})

	err := rdb.Ping(ctx).Err()

	if err != nil {
		log.Fatalf("error connecting to redis: %s", err.Error())
	}

	return rdb
}
