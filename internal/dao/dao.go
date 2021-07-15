package dao

import (
	"github.com/go-redis/cache/v8"
	"github.com/go-redis/redis/v8"
	"github.com/jmoiron/sqlx"
)

type Dao struct {
	db    *sqlx.DB
	rdb   *redis.Client
	cache *cache.Cache
}

func New(db *sqlx.DB, rdb *redis.Client, cache *cache.Cache) *Dao {
	return &Dao{
		db:    db,
		rdb:   rdb,
		cache: cache,
	}
}
