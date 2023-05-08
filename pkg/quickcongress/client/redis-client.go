package client

import (
	"github.com/go-redis/redis"
)

type QuickCongressRedisClient struct {
	redisHost     string
	redisPassword string
	redisClient   *redis.Client
}

func (q *QuickCongressRedisClient) NewClient(redisHost string, redisPassword string) *QuickCongressRedisClient {

	return &QuickCongressRedisClient{
		redisHost:     redisHost,
		redisPassword: redisPassword,
		redisClient: redis.NewClient(&redis.Options{
			Addr:     redisHost,
			Password: redisPassword,
			DB:       0,
		}),
	}
}

// func ReconnectRedis(client *QuickCongressRedisClient) {
// 	client.redisClient.Close()
// 	client.redisClient = NewClient()
// 	return
// }

// func SetValue(client *QuickCongressRedisClient, key string, value *congresses.CongressRes) {
// 	client.redisClient.SetValue()
// }

// func GetValue(client *QuickCongQuickCongressRedisClient, key string) interface{}
