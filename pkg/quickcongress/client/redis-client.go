package client

import (
	"log"
	"time"

	"os"

	"github.com/go-redis/cache/v8"
	"github.com/go-redis/cache/v9"
	"github.com/go-redis/redis"
	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
	"github.com/zfoteff/quick-congress/bin"
)

var logger = bin.NewLogger("Cache", "cache.log")

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

	ring := redis.NewRing(&redis.RingOptions{
		Addrs: map[string]string{
			host: "6379",
		},
	})

	cache := cache.New(&cache.Options{
		Redis:      ring,
		LocalCache: cache.NewTinyLFU(1000, time.Minute),
	})

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

func (q *QuickCongressRedisClient) Reconnect() {
	logger.Info("[*] Reconnecting to Redis Cache")
	q.redisClient.Close()
	q = NewRedisClient()
	logger.Info("[+] Reconnected to Redis Cache")
}

func (q *QuickCongressRedisClient) SetCacheValue(url string, response interface{}) bool {
	///////////////////////////////////////////////////////////
	// err := q.redisClient.Set(url, response, time.Hour)	 //
	// 														 //
	// if err != nil {										 //
	// 	log.Fatal(err)									 //
	// }													 //
	///////////////////////////////////////////////////////////

	return true
}

func (q *QuickCongressRedisClient) GetCacheValue(url string) (bool, string) {
	/////////////////////////////////////////////////////
	// value, err := q.redisClient.Get("").Result()	   //
	// 												   //
	// if err != nil {								   //
	// 	log.Fatal(err)							   //
	// }											   //
	// 												   //
	// println(value)								   //
	/////////////////////////////////////////////////////
	return true, "true"
}
