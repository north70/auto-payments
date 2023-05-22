package redis

import (
	"fmt"
	"github.com/redis/go-redis/v9"
)

func NewRedisDB(host, port string) (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", host, port),
		Password: "",
		DB:       0,
	})

	return rdb, nil
}
