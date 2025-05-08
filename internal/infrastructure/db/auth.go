package db

import (
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

func GetRedisVal(key string) (string, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // 没有密码，默认值
		DB:       0,  // 默认DB 0
	})
	ctx := context.Background()
	val, err := rdb.Get(ctx, key).Result()
	return val, err
}

func SetRedisVal(key string, value interface{}, duration time.Duration) (string, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // 没有密码，默认值
		DB:       0,  // 默认DB 0
	})
	ctx := context.Background()
	val, err := rdb.Set(ctx, key, value, duration).Result()
	return val, err
}
