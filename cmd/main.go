package main

import (
	"log"
	"testEx/api"
	"testEx/configs"
	"testEx/internal"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Загрузка конфигурации
	configs.LoadConfig()

	// Инициализация базы данных
	internal.InitDB(configs.GetDBURL())

	app := fiber.New()

	// Настройка маршрутов
	api.SetupRoutes(app)

	// Запуск сервера
	log.Fatal(app.Listen(":3000"))
}
