package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"os"

	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
	"github.com/zfoteff/quick-congress/bin"
)

var cacheLogger = bin.NewLogger("Cache", "cache.log")

type QuickCongressRedisCache struct {
	redisHost     string
	redisPassword string
	redisCache    *redis.Client
}

// Create new instance of a Redis Cache
func NewQuickCongressRedisCache() *QuickCongressRedisCache {
	if goEnvErr := godotenv.Load(".env"); goEnvErr != nil {
		log.Fatal(goEnvErr)
	}

	host := os.Getenv("REDIS_HOST")
	password := os.Getenv("REDIS_PASSWORD")

	return &QuickCongressRedisCache{
		redisHost:     host,
		redisPassword: password,
		redisCache: redis.NewClient(&redis.Options{
			Addr:         host,
			Password:     password,
			DB:           0,
			MaxRetries:   5,
			MaxIdleConns: 5,
		}),
	}
}

// Set a value in the cache using the URL as a key and the response as a value
func (q *QuickCongressRedisCache) SetCacheValue(url string, response interface{}) error {
	var resBytes bytes.Buffer

	if err := json.NewEncoder(&resBytes).Encode(response); err != nil {
		cacheLogger.Error(fmt.Sprintf("Error unmarshalling object when setting value for key: %s", url), err)
		return err
	}

	if cacheErr := q.redisCache.Set(context.Background(), url, resBytes, time.Hour).Err(); cacheErr != nil {
		cacheLogger.Error("Error setting value in the cache", cacheErr)
		return cacheErr
	}

	return nil
}

// Get a cached response using the URL as the key
func (q *QuickCongressRedisCache) GetCacheValue(url string, response interface{}) error {
	value, cacheErr := q.redisCache.Get(context.Background(), url).Bytes()
	if cacheErr != nil {
		// Cache miss, return error to exit function
		cacheLogger.Warning(fmt.Sprintf("Cache miss for key: '%s'", url))
		return cacheErr
	}

	if err := json.NewDecoder(bytes.NewReader(value)).Decode(response); err != nil {
		cacheLogger.Error("Error hydrating response from cache", err)
		return err
	}

	cacheLogger.Info(fmt.Sprintf("Cache hit for key: '%s'", url))
	return nil
}

func (q *QuickCongressRedisCache) healthCheck(client *redis.Client) error {
	if _, err := q.redisCache.Ping(context.TODO()).Result(); err != nil {
		return err
	}

	return nil
}

// func (q *QuickCongressRedisCache) healthCheck(client *redis.Client) error {
// 	if _, err := q.redisCache.Ping(context.TODO()).Result(); err != nil {
// 		return err
// 	}

// 	return nil
// }
