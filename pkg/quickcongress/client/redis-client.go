package client

import (
	"log"
	"time"

	"os"

	"github.com/go-redis/redis"
	"github.com/joho/godotenv"
	"github.com/zfoteff/quick-congress/pkg/quickcongress/model"
)

type QuickCongressRedisClient struct {
	redisHost     string
	redisPassword string
	redisClient   *redis.Client
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
		}),
	}
}

func (q *QuickCongressRedisClient) SetCacheValue(url string, response model.CongressSuccessRes) {
	err := q.redisClient.Set(url, response, time.Hour)

	if err != nil {
		log.Fatal(err)
	}

}

func (q *QuickCongressRedisClient) GetCacheValue(url string) (bool, string) {
	value, err := q.redisClient.Get("").Result()

	if err != nil {
		log.Fatal(err)
	}

	println(value)
	return false, value
}

func (q *QuickCongressRedisClient) Reconnect() {
	q.redisClient.Close()
	q = NewRedisClient()
}
