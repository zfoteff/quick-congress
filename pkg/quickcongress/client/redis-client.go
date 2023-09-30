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

type QuickCongressRedisClient struct {
	redisHost     string
	redisPassword string
	redisClient   *redis.Client
}

var cacheLogger = bin.NewLogger("Cache", "cache.log")

/**
 * Create new instance of a Redis connection and return the pointer to the new connecti
 */
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

/**
 * Create a new connection to the Redis cache
 */
func (q *QuickCongressRedisClient) Reconnect() {
	cacheLogger.Info("[*] Reconnecting to Redis Cache")
	q.redisClient.Close()
	q = NewRedisClient()
	cacheLogger.Info("[+] Reconnected to Redis Cache")
}

/**
 * Set a value in the cache using the URL as a key and the response as a value
 */
func (q *QuickCongressRedisClient) SetCacheValue(url string, response interface{}) bool {
	err := q.redisClient.Set(context.TODO(), url, response, time.Hour).Err()

	if err != nil {
		cacheLogger.Error("Error setting value in the cache", err)
	}

	return true
}

/**
 * Get a cached response using the URL as the key
 */
func (q *QuickCongressRedisClient) GetCacheValue(url string) (bool, string) {
	value, err := q.redisClient.Get(context.TODO(), url).Result()

	if err != nil {
		// Cache miss
		cacheLogger.Warning(fmt.Sprintf("Cache miss: %s", url))
		return false, ""
	}

	// Cache hit
	cacheLogger.Info("Cache hit")
	return true, value
}
