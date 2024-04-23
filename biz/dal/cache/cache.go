package cache

import (
	"time"
)

type Cache interface {
	CacheGet(key string) (any, error)
	CacheSet(key string, value any, duration time.Duration) error
	CacheDel(key string) error
}
