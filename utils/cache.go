package utils

import (
	"context"
	"github.com/go-redis/redis/v8"
	"log"

	"os"
)

// ConnectToRedis func
func ConnectToRedis() *redis.Client {
	ctx := context.Background()
	host := os.Getenv("REDIS_HOST")
	password := os.Getenv("REDIS_PASSWORD")

	client := redis.NewClient(&redis.Options{
		Addr: host,
		Password: password,
		DB: 0,
	})

	if err := client.Ping(ctx).Err(); err != nil {
		log.Printf("[r] unable to connect to redis because %v", err)
		os.Exit(3)
	}

	log.Printf("[r] connected to redis")
	return client
}
