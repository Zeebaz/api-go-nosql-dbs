package rediscontroller

import (
	"context"
	"os"

	"github.com/go-redis/redis/v9"
)

func RedisConnection() *redis.Client {
	/* err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
		fmt.Println(err)
	} */

	redisHost := os.Getenv("REDIS_HOST")
	redisPass := os.Getenv("REDIS_PASSWORD")
	// redisDB := os.Getenv("REDIS_DB")

	rdb := redis.NewClient(&redis.Options{
		Addr:     redisHost,
		Password: redisPass,
		DB:       0,
	})
	return rdb
}

func SetSADD(key string, member string) (int64, error) {
	rdb := RedisConnection()
	ctx := context.Background()

	return rdb.SAdd(ctx, key, member).Result()
}

func GetSMEMBERS(key string) ([]string, error) {
	rdb := RedisConnection()
	ctx := context.Background()

	return rdb.SMembers(ctx, key).Result()
}

func SetHINCRBY(key string, field string, increment int64) (int64, error) {
	rdb := RedisConnection()
	ctx := context.Background()

	return rdb.HIncrBy(ctx, key, field, increment).Result()
}
func GetHGETALL(key string) (map[string]string, error) {
	rdb := RedisConnection()
	ctx := context.Background()

	return rdb.HGetAll(ctx, key).Result()
}
