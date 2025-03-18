package core

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/redis/go-redis/v9"
)

type Redis struct {
	client *redis.Client
}

type IRedis interface {
	Get(key string) (string, error)
	Set(key string, value string, ttl time.Duration) error
	Delete(key string) error
}

var _ IRedis = (*Redis)(nil)

func (r *Redis) Get(key string) (string, error) {
	ctx := context.Background()
	return r.client.Get(ctx, key).Result()
}

func (r *Redis) Set(key string, value string, ttl time.Duration) error {
	ctx := context.Background()
	return r.client.Set(ctx, key, value, ttl).Err()
}

func (r *Redis) Delete(key string) error {
	ctx := context.Background()
	return r.client.Del(ctx, key).Err()
}

func NewCacheServer() *Redis {
	config := redis.Options{
		Addr:     os.Getenv("REDIS_ADDR"),
		Password: os.Getenv("REDIS_PASSWORD"),
	}

	client := redis.NewClient(&config)

	err := client.Ping(context.Background()).Err()

	// Check if Redis connection is successful
	if err != nil {
		panic(fmt.Sprintf("Failed to connect to Redis at %s: %v", config.Addr, err))
	}

	return &Redis{
		client: client,
	}
}
