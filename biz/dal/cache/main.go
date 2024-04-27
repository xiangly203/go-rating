package cache

func RedisInit() *RedisCache {
	return GetRedisIns(RedisConf{
		Addr:     "localhost:6379",
		Password: "",
		Db:       0,
		PoolSize: 10,
	})
}
