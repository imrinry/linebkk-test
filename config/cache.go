package config

import (
	"context"
	"log"
	"os"

	"github.com/go-redis/redis/v8"
)

var Redis *redis.Client

func ConnectRedis() {
	Redis = redis.NewClient(&redis.Options{
		Addr: os.Getenv("REDIS_ADDR"),
	})

	if err := Redis.Ping(context.Background()).Err(); err != nil {
		log.Fatal("Failed to connect to Redis:", err)
	}

	log.Println("Connected to Redis")
}

func GetRedisInstance() *redis.Client {
	return Redis
}
