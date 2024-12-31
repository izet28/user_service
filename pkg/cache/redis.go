package cache

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

type RedisCache struct {
	Client *redis.Client
}

func NewRedisCache(addr string, db int) *RedisCache {
	client := redis.NewClient(&redis.Options{
		Addr: addr,
		DB:   db,
	})
	return &RedisCache{Client: client}
}

func (rc *RedisCache) Set(key string, value interface{}, expiration time.Duration) error {
	return rc.Client.Set(ctx, key, value, expiration).Err()
}

func (rc *RedisCache) Get(key string) (string, error) {
	return rc.Client.Get(ctx, key).Result()
}

func (rc *RedisCache) Delete(key string) error {
	return rc.Client.Del(ctx, key).Err()
}
