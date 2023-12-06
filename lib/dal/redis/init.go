package redis

import (
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

func GetRedisVal(key string) (string, error) {
	//opt, err := redis.ParseURL("redis://"":<pass>@localhost:6379/<db>")
	//if err != nil {
	//	panic(err)
	//}
	//rdb := redis.NewClient(opt)
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // 没有密码，默认值
		DB:       0,  // 默认DB 0
	})
	ctx := context.Background()
	val, err := rdb.Get(ctx, key).Result()
	return val, err
}

func SetRedisVal(key string, value interface{}, duration time.Duration) string {
	//opt, err := redis.ParseURL("redis://<user>:<pass>@localhost:6379/<db>")
	//if err != nil {
	//	panic(err)
	//}
	//rdb := redis.NewClient(opt)
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // 没有密码，默认值
		DB:       0,  // 默认DB 0
	})
	ctx := context.Background()
	get := rdb.Set(ctx, key, value, duration)
	return get.Val()
}
