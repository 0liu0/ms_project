package dao

import (
	"time"

	"github.com/go-redis/redis"
)

var Rc *RedisCache

type RedisCache struct {
	rdb *redis.Client
}

func init() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "liuche17.cn:3061",
		Password: "liuche", // no password set
		DB:       0,        // use default DB
	})
	Rc = &RedisCache{
		rdb: rdb,
	}
}

func (rc *RedisCache) Put(key string, value string, expire time.Duration) error {
	error := rc.rdb.Set(key, value, expire).Err()
	return error
}

func (rc *RedisCache) Get(key string) (string, error) {
	res, err := rc.rdb.Get(key).Result()
	return res, err
}
