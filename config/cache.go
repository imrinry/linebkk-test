package config

import (
	"context"
	"line-bk-api/pkg/logs"
	"log"
	"os"

	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
)

var Redis *redis.Client

func ConnectRedis() {
	Redis = redis.NewClient(&redis.Options{
		Addr: os.Getenv("REDIS_ADDR"),
	})

	if err := Redis.Ping(context.Background()).Err(); err != nil {
		log.Fatal("Failed to connect to Redis:", err)
	}

	ping := Redis.Ping(context.Background())
	if ping.Err() != nil {
		log.Fatal("Failed to ping Redis:", ping.Err())
	}

	logs.Info("Connected to Redis Ping", zap.Any("ping", ping))
}

func GetRedisInstance() *redis.Client {
	return Redis
}
