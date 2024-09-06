package models

import (
	"context"

	"github.com/go-redis/redis/v8"
)

var (
	RedisDb *redis.Client
)

// 创建 redis 链接
func init() {
	var c = context.Background()
	RedisDb = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	_, err := RedisDb.Ping(c).Result()
	if err != nil {
		println(err)
	}
}
