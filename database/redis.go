package database

import (
	"my_diary/config"

	"strconv"
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

var (
	client *redis.Client
)

func initRedisClient() {
	redisConfig := config.GetRedisConfig()
	client = redis.NewClient(&redis.Options{
		Addr:     redisConfig.Host + ":" + strconv.Itoa(redisConfig.Port),
		Password: redisConfig.Password,
		DB:       0,
	})
}

func SetValue(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	return client.Set(ctx, key, value, expiration).Err()
}

func GetValue(ctx context.Context, key string) (string, error) {
	return client.Get(ctx, key).Result()
}

func DeleteValue(ctx context.Context, key string) error {
	return client.Del(ctx, key).Err()
}

