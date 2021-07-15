package global

import (
	"context"
	"encoding/json"
	"github.com/go-redis/cache/v8"
	"github.com/go-redis/redis/v8"
)

var (
	RedisDB    *redis.Client
	RedisCache *cache.Cache
)

func SetupRedisDB() error {
	RedisDB = redis.NewClient(&redis.Options{
		Addr: RedisSetting.Addr,
		DB:   RedisSetting.DB,
	})

	err := RedisDB.Ping(context.Background()).Err()
	if err != nil {
		return err
	}

	return nil
}

func SetupRedisCache() {
	RedisCache = cache.New(&cache.Options{
		Redis: RedisDB,
		//LocalCache:   cache.NewTinyLFU(1000, time.Minute),
		Marshal:   json.Marshal,
		Unmarshal: json.Unmarshal,
	})
}

func SetupRedis() error {
	err := SetupRedisDB()
	if err != nil {
		return err
	}
	SetupRedisCache()

	return nil
}
