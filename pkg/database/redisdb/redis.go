package redisdb

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/onemgvv/linkshot.git/internal/config"
)

func Init(cfg *config.Config) *redis.Client {
	addr := fmt.Sprintf("%s:%s", cfg.Redis.Host, cfg.Redis.Port)
	return redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: cfg.Redis.Password,
		DB:       0,
	})
}

func Close(rds *redis.Client) error {
	if err := rds.Close(); err != nil {
		return err
	}

	return nil
}
