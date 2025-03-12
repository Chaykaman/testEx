package main

import (
	"log"

	"github.com/Chaykaman/testEx/config"
	"github.com/Chaykaman/testEx/internal"
	v1 "github.com/Chaykaman/testEx/internal/controller/api/v1"
	"github.com/gofiber/fiber/v2"
)

func main() {
	// Загрузка конфигурации
	config.LoadConfig()

	// Инициализация базы данных
	internal.InitDB(config.GetDBURL())

	app := fiber.New()

	v1.SetupRoutes(app)

	// Запуск сервера
	log.Fatal(app.Listen(":3000"))
}
