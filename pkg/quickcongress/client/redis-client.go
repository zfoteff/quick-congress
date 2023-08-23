package client

import (
	"log"
	"time"

	"os"

	"github.com/go-redis/cache/v8"
	"github.com/go-redis/redis/v8"
	"github.com/joho/godotenv"
)

type QuickCongressRedisClient struct {
	redisHost     string
	redisPassword string
	redisClient   *redis.Client
	redisCache    *cache.Cache
}

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
		}), redisCache: nil,
	}
}

func (q *QuickCongressRedisClient) ConnectToCache(client *QuickCongressRedisClient) *QuickCongressRedisCache {
	return cache.New(&cache.Options{
		Redis: q.redisClient,
	})
}

func (q *QuickCongressRedisClient) Reconnect() {
	log.Print("[*] Disconnecting from Redis Cache ...")
	q.redisClient.Close()
	log.Print("[-] Disconnected from Redis Cache ...")
	log.Print("[*] Reconnecting to Redis Cache ...")
	q = NewRedisClient()
	log.Print("[+] Reconnected to Redis Cache ...")
}

func (q *QuickCongressRedisClient) SetCacheValue(url string, response interface{}) bool {
	err := q.redisClient.Set(url, response, time.Hour)

	if err != nil {
		log.Fatal(err)
	}

	return true
}

func (q *QuickCongressRedisClient) GetCacheValue(url string) (bool, string) {
	value, err := q.redisClient.Get("").Result()

	if err != nil {
		log.Fatal(err)
	}

	println(value)
	return true, value
}
