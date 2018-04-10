package cache

import (
	"github.com/go-redis/redis"
)

func OpenRedis() (*redis.Client) {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "alexbrasser",
		DB:       0,
	})
	return client
}
