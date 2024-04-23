package cache

import (
	"context"
	"encoding/json"
	"github.com/redis/go-redis/v9"
	"sync"
	"time"
)

type RedisCache struct {
	client *redis.Client
	ctx    context.Context
}

type RedisInstance struct {
	RedisCache
}

type RedisConf struct {
	Addr     string
	Password string
	Db       int
}

var (
	redisInstance *RedisInstance
	once          sync.Once
)

func GetCacheIns(redisConf RedisConf) *RedisInstance {
	once.Do(func() {
		redisInstance = &RedisInstance{
			RedisCache: *NewRedisCache(redisConf),
		}
	})
	return redisInstance
}

func NewRedisCache(redisConf RedisConf) *RedisCache {
	rdb := redis.NewClient(&redis.Options{
		Addr:     redisConf.Addr,
		Password: redisConf.Password,
		DB:       redisConf.Db,
	})
	return &RedisCache{
		client: rdb,
		ctx:    context.Background(), // Initializing the context
	}
}

// CacheGet retrieves a value from Redis.
func (rc *RedisCache) CacheGet(key string) (any, error) {
	var value any
	str, err := rc.client.Get(rc.ctx, key).Result()
	if err != nil {
		return value, err
	}
	err = json.Unmarshal([]byte(str), &value)
	if err != nil {
		return value, err
	}
	return value, nil
}

func (rc *RedisCache) CacheSet(key string, value any, duration time.Duration) error {
	val, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return rc.client.Set(rc.ctx, key, val, duration).Err()
}

func (rc *RedisCache) CacheDel(key string) error {
	return rc.client.Del(rc.ctx, key).Err()
}
