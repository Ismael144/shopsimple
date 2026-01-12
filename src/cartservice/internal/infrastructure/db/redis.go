package db

import (
	"github.com/redis/go-redis/v9"
	"time"
)

func NewRedis(addr string) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "",
		DB:       0,

		// Retry config
		MaxRetries:      3,
		MinRetryBackoff: 10 * time.Millisecond,
		MaxRetryBackoff: 512 * time.Millisecond,

		// Timeout config
		DialTimeout:  5 * time.Second,
		ReadTimeout:  3 * time.Second,
		WriteTimeout: 3 * time.Second,
		PoolTimeout:  30 * time.Second,
	})

	return rdb
}
