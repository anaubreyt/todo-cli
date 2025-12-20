package config

import (
	"os"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)
// объявляем переменную приватную, в которой хранится бд с типом gorm.DB
var db *gorm.DB

// инициализируем и записываем бд в переменную
func init() {
	err := godotenv.Load(".env")
	if err != nil {
		panic("Ошибка загрузки .env файла")
	}
	
	dsn := os.Getenv("DATABASE_URL")

	d, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db = d
}

// публичный метод для получения указателя на бд типа gorm.DB
func GetDB() *gorm.DB {
	return db
}