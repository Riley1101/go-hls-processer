package config

import (
	"github.com/redis/go-redis/v9"
	"log/slog"
)

type RedisConfig struct {
	Host string
	Port string
	DB   int
}

func (c *RedisConfig) getRedisOptions() *redis.Options {
	opt, err := redis.ParseURL("redis://" + c.Host + ":" + c.Port)
	if err != nil {
		slog.Error("redis parse url error: ", err)
	}
	return opt
}

func ConnectRedis() *redis.Client {
	redisConfig := RedisConfig{
		Host: "127.0.0.1",
		Port: "6379",
		DB:   0,
	}
	client := redis.NewClient(redisConfig.getRedisOptions())
	return client
}
