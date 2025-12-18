package config

import (
	"log"
	"os"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"github.com/joho/godotenv"
)

var db *gorm.DB

func init() {
	godotenv.Load()

	dsn := os.Getenv("POSTGRES_DATABASE_URL") 
	// dsn := "host=localhost user=postgres dbname=todo port=5432 sslmode=disable"
	log.Print("[DEBUG dsn]", dsn)
	d, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	
	db = d
}
func GetDB() *gorm.DB {
	return db
}
