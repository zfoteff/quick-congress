package client

import (
	"log"
	"time"

	"os"

	"github.com/go-redis/redis"
	"github.com/joho/godotenv"
)

type QuickCongressRedisClient struct {
	redisHost     string
	redisPassword string
	redisClient   *redis.Client
}

type QuickCongressRedisCache struct {
	QuickCongressRedisClient
}

func newRedisClient() *QuickCongressRedisClient {
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

func NewRedisCache() *

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

func (q *QuickCongressRedisClient) Reconnect() {
	log.Print("[*] Disconnecting from Redis Cache ...")
	q.redisClient.Close()
	log.Print("[-] Disconnected from Redis Cache ...")
	log.Print("[*] Reconnecting to Redis Cache ...")
	q = NewRedisClient()
	log.Print("[+] Reconnected to Redis Cache ...")
}
