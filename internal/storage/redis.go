package storage

import (
	"AutoPayment/config"
	"github.com/redis/go-redis/v9"
)

func NewRedisDB(cfg config.Redis) (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     cfg.Address(),
		Password: "",
		DB:       0,
	})

	return rdb, nil
}
