package redisCache

import (
	"context"
	_ "fmt"
	"os"
	"time"
	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
)

var client *redis.Client
var ctx context.Context

func init() {
	godotenv.Load()

	client = redis.NewClient(&redis.Options{
		Addr: os.Getenv("REDIS_DATABASE_URL"),
	})

	ctx = context.Background()
}

func GetRedisClient() *redis.Client {
	return client
}

func GetKey(key string) (val any, err error) {
	val, err = client.Get(ctx, key).Result()
	return
}

func SetKey(key string, val any, exp_time time.Duration) (err error) {
	err = client.Set(ctx, key, val, exp_time).Err()
	return
}

