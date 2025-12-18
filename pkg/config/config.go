package config

import (
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB
var redisClient *redis.Client

func init() {
	dsn := "host=localhost user=postgres password=Qwerty123$ dbname=todo port=5432 sslmode=disable"
	d, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db = d

	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	redisClient = client
}
func GetDB() *gorm.DB {
	return db
}

func GetRedis() *redis.Client {
	return redisClient
}
