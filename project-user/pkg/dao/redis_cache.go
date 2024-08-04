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
		Addr:        "liuche17.cn:3061",
		Password:    "liuche",        // 设置密码，没有密码为空即可
		DB:          0,               // 设置数据库
		DialTimeout: 5 * time.Second, // 设置五秒超时时间
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
