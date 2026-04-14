package config

import (
	"log"

	"github.com/redis/go-redis/v9"
)

func ConnectRedis(redisString string) *redis.Client {

	opts, err := redis.ParseURL(redisString)
	if err != nil {
		panic(err)
	}
	log.Println("connected to redis")
	return redis.NewClient(opts)
}
