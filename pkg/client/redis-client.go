package client

import (
	"fmt"
	"log"
	"os"

	"github.com/go-redis/redis"
	"github.com/joho/godotenv"
)

err := godotenv.Load(".env")

if err != nil {
	log.Fatalf("Error loading .env file")
}

const (
	RedisHost=os.Getenv("REDIS_HOST")
	RedisPassword=os.Getenv("REDIS_PASSWORD")
)

type QuickCongressRedisClient struct {
	redisHost string
	redisPassword string
	redisClient *redis.Client	
}

func (q *QuickCongressRedisClient) NewClient() *QuickCongressRedisClient {
	return &QuickCongressRedisClient{
		redisHost: RedisHost,
		redisPassword: RedisPassword,
		redisClient: &redis.NewClient(&redis.Options{
			Addr: redisHost,
			Password: redisPassword,
			DB: 0,
		},),
	},
}

func ReconnectRedis(client *QuickCongressRedisClient) {
	client.redisClient.Close()
	client.redisClient = NewClient()
	return
}

func SetValue(client *QuickCongressRedisClient, key string, value *congresses.CongressRes) {
	client.redisClient.SetValue()
}

func GetValue(client *QuickCongQuickCongressRedisClient, key string) interface{}