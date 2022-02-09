package redis

import "github.com/go-redis/redis/v8"

func InitRedis(endpoint string) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     endpoint,
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	return rdb
}
