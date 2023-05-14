package pkg_redis

import (
	"github.com/redis/go-redis/v9"
)

func NewRedisClient() *redis.Client {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "Secret_Password_1234",
		DB:       0, // use default DB
	})

	return redisClient
}
