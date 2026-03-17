package redis

import (
	"context"
	"github.com/redis/go-redis/v9"
	"log"
)

const (
	redisHost = "localhost:6379"
	password  = ""
	db        = 0
	protocol  = 2
)

func NewRedisConnection() (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     redisHost,
		Password: password,
		DB:       db,
		Protocol: protocol,
	})
	ctx := context.Background()
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Connecting to redis error: %v\n", err)
		return nil, err
	}
	log.Printf("Connected to redis succesfully\n")
	return rdb, err
}
