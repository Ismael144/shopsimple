package repository

import "github.com/redis/go-redis/v9"

type CartRepositoryRedis struct {
	rdb *redis.Client
}

func NewCartRepositoryRedis(rdb *redis.Client) *CartRepositoryRedis {
	return &CartRepositoryRedis{rdb: rdb}
}

