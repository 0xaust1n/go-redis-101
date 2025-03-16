package core

import (
	"fmt"
	"os"
	"time"

	"github.com/go-redis/redis"
)

type Redis struct {
	client *redis.Client
}

type IRedis interface {
	Get(key string) (string, error)
	Set(key string, value string, ttl time.Duration) error
}

var _ IRedis = (*Redis)(nil)

func (r *Redis) Get(key string) (string, error) {
	return r.client.Get(key).Result()
}

func (r *Redis) Set(key string, value string, ttl time.Duration) error {
	return r.client.Set(key, value, ttl).Err()
}

func NewCacheServer() *Redis {
	config := redis.Options{
		Addr:     os.Getenv("REDIS_ADDR"),
		Password: os.Getenv("REDIS_PASSWORD"),
	}

	client := redis.NewClient(&config)

	// Check if Redis connection is successful
	_, err := client.Ping().Result()
	if err != nil {
		panic(fmt.Sprintf("Failed to connect to Redis at %s: %v", config.Addr, err))
	}

	return &Redis{
		client: client,
	}
}
