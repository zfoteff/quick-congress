package client

import (
	"context"
	"fmt"
	"log"
	"time"

	"os"

	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
	"github.com/zfoteff/quick-congress/bin"
)

var cacheLogger = bin.NewLogger("Cache", "cache.log")

type QuickCongressRedisClient struct {
	redisHost     string
	redisPassword string
	redisClient   *redis.Client
}

// Create new instance of a Redis connection and return the pointer to the new connection
func NewRedisClient() *QuickCongressRedisClient {
	if goEnvErr := godotenv.Load(".env"); goEnvErr != nil {
		log.Fatal(goEnvErr)
	}

	host := os.Getenv("REDIS_HOST")
	password := os.Getenv("REDIS_PASSWORD")

	return &QuickCongressRedisClient{
		redisHost:     host,
		redisPassword: password,
		redisClient: redis.NewClient(&redis.Options{
			Addr:     host,
			Password: password,
			DB:       0,
		}),
	}
}

// Set a value in the cache using the URL as a key and the response as a value
func (q *QuickCongressRedisClient) SetCacheValue(url string, response interface{}) error {
	err := q.redisClient.Set(context.TODO(), url, response, time.Hour).Err()

	if err != nil {
		cacheLogger.Error("Error setting value in the cache", err)
		return err
	}

	return nil
}

// Get a cached response using the URL as the key
func (q *QuickCongressRedisClient) GetCacheValue(url string) (bool, interface{}) {
	value, err := q.redisClient.Get(context.TODO(), url).Result()

	if err != nil {
		cacheLogger.Warning(fmt.Sprintf("Cache miss for key: '%s'", url))
		return false, ""
	}

	cacheLogger.Info(fmt.Sprintf("Cache hit for key: '%s'", url))
	return true, value
}
