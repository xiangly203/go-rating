package cache

import (
	"context"
	"github.com/redis/go-redis/v9"
	"sync"
	"time"
)

type RedisCache struct {
	client *redis.Client
	ctx    context.Context
}

type RedisConf struct {
	Addr     string
	Password string
	Db       int
	PoolSize int
}

var (
	once sync.Once
)

func GetRedisIns(redisConf RedisConf) *RedisCache {
	var redisIns *RedisCache
	once.Do(func() {
		rdb := redis.NewClient(&redis.Options{
			Addr:     redisConf.Addr,
			Password: redisConf.Password,
			DB:       redisConf.Db,
		})
		redisIns = &RedisCache{
			client: rdb,
			ctx:    context.Background(),
		}
	})
	return redisIns
}

func (rc *RedisCache) CacheGet(key string) (any, error) {
	var value any
	value, err := rc.client.Get(rc.ctx, key).Result()
	if err != nil {
		return value, err
	}
	return value, err
}

func (rc *RedisCache) CacheSet(key string, value any, duration time.Duration) error {
	return rc.client.Set(rc.ctx, key, value, duration).Err()
}

func (rc *RedisCache) CacheDel(key string) error {
	return rc.client.Del(rc.ctx, key).Err()
}
