package cache

import (
	"context"

	"github.com/NaturalSelectionLabs/RSS3-PreGod/shared/pkg/config"
	"github.com/NaturalSelectionLabs/RSS3-PreGod/shared/pkg/logger"
	"github.com/go-redis/redis/v8"
)

var rdb *redis.Client
var log = logger.Logger()

func Setup() error {
	rdb = redis.NewClient(&redis.Options{
		Addr:     config.Redis.Addr,
		Password: config.Redis.Password,
		DB:       config.Redis.DB,
	})

	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		return err
	}

	return nil
}
