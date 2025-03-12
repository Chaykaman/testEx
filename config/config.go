package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var dbURL string

func LoadConfig() {
	// Загрузка переменных окружения из .env файла
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Ошибка загрузки файла .env")
	}

	dbURL = os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("Не задан файл .env")
	}

	// Отладочный вывод
	log.Println("DATABASE_URL загружен:", dbURL)
}

func GetDBURL() string {
	return dbURL
}
