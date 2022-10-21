package rediscontroller

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/go-redis/redis/v9"
	"github.com/joho/godotenv"
)

func RedisConnection() *redis.Client {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	redisHost := os.Getenv("REDIS_HOST")

	rdb := redis.NewClient(&redis.Options{
		Addr:     redisHost,
		Password: "",
		DB:       0,
	})
	return rdb
}

func SetNewMatch(key string, field string) map[string]string {
	rdb := RedisConnection()
	ctx := context.Background()

	val, err := rdb.HGetAll(ctx, "matches").Result()
	switch {
	case err == redis.Nil:
		fmt.Println("key does not exist")
	case err != nil:
		fmt.Println("Get failed", err)
	}

	return val
}

func GetHGETALL(key string) (map[string]string, error) {
	rdb := RedisConnection()
	ctx := context.Background()

	val, err := rdb.HGetAll(ctx, key).Result()

	return val, err
}
