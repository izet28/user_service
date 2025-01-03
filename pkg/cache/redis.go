package cache

import (
	"context"
	"encoding/json"
	"fmt"
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
	// Convert struct to JSON
	jsonData, err := json.Marshal(value)
	if err != nil {
		fmt.Printf("Failed to marshal data for key: %s, error: %v\n", key, err)
		return err
	}

	// Store JSON in Redis
	err = rc.Client.Set(ctx, key, jsonData, expiration).Err()
	if err != nil {
		fmt.Printf("Failed to set data in Redis for key: %s, error: %v\n", key, err)
	} else {
		fmt.Printf("Successfully set data in Redis for key: %s\n", key)
	}
	return err
}

func (rc *RedisCache) Get(key string) (string, error) {
	data, err := rc.Client.Get(ctx, key).Result()
	if err == redis.Nil {
		// Log jika data tidak ditemukan di Redis
		fmt.Printf("Cache miss for key: %s\n", key)
		return "", err
	} else if err != nil {
		// Log jika terjadi error saat mengakses Redis
		fmt.Printf("Error accessing cache for key: %s, error: %v\n", key, err)
		return "", err
	}

	// Log jika data ditemukan di Redis
	fmt.Printf("Cache hit for key: %s\n", key)
	return data, nil
}

func (rc *RedisCache) Delete(key string) error {
	return rc.Client.Del(ctx, key).Err()
}
