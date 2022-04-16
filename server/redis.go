package server

import (
	"fmt"
	"github.com/go-redis/redis/v8"
)

var Client *redis.Client

func InitializeRedisClient() {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%v:%v", "localhost", "6379"),
		Password: "", // no password set
		//DB:       conf.RedisDefaultDb(), // use default DB,
		PoolSize: 20,
	})
	Client = redisClient
}
