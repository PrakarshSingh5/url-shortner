package database

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
)

var (
	Rdb *redis.Client
	Ctx = context.Background()
)

// ConnectRedis initializes Redis client
func ConnectRedis(addr string, password string, db int) {
	Rdb = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})

	_, err := Rdb.Ping(Ctx).Result()
	if err != nil {
		log.Fatalf("failed to connect to Redis: %v", err)
	}
	log.Println("Connected to Redis")
}
