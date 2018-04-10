package database

import "github.com/go-redis/redis"

func OpenRedis() (*redis.Client) {
	db := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "alexbrasser", // no password set
		DB:       0,             // use default DB
	})
	return db
}
