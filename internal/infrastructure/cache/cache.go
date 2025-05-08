package cache

import (
	"time"
)

type Cache interface {
	CacheGet(key string) (any, error)
	CacheSet(key string, value any, duration time.Duration) error
	CacheDel(key string) error
}

var Ins Cache

func CacheInit() {
	Ins = RedisInit()
}

func CacheGet(key string) (any, error) {
	return Ins.CacheGet(key)
}

func CacheSet(key string, value any, duration time.Duration) error {
	return Ins.CacheSet(key, value, duration)
}

func CacheDel(key string) error {
	return Ins.CacheDel(key)
}
