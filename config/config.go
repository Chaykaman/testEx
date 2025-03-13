package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	DBHost     string
	DBPort     int
	DBUser     string
	DBPassword string
	DBName     string
	AppPort    int
}

var AppConfig Config

func LoadConfig() {
	// Загрузка переменных окружения из .env файла
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Ошибка загрузки .env файла")
	}

	// Загрузка конфигурации базы данных
	AppConfig.DBHost = os.Getenv("DB_HOST")
	AppConfig.DBPort, _ = strconv.Atoi(os.Getenv("DB_PORT"))
	AppConfig.DBUser = os.Getenv("DB_USER")
	AppConfig.DBPassword = os.Getenv("DB_PASSWORD")
	AppConfig.DBName = os.Getenv("DB_NAME")

	// Загрузка порта приложения
	AppConfig.AppPort, _ = strconv.Atoi(os.Getenv("APP_PORT"))

	// Проверка обязательных переменных
	if AppConfig.DBHost == "" || AppConfig.DBUser == "" || AppConfig.DBPassword == "" || AppConfig.DBName == "" {
		log.Fatal("Не все переменные окружения для базы данных заданы")
	}
}

// GetDBURL возвращает строку подключения к базе данных
func GetDBURL() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s",
		AppConfig.DBUser,
		AppConfig.DBPassword,
		AppConfig.DBHost,
		AppConfig.DBPort,
		AppConfig.DBName,
	)
}

// GetAppPort возвращает порт приложения
func GetAppPort() string {
	return fmt.Sprintf(":%d", AppConfig.AppPort)
}
