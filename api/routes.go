package api

import (
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Post("/tasks", createTask)
	app.Get("/tasks", getTasks)
	app.Get("/tasks/:id", getTaskByID)
	app.Put("/tasks/:id", updateTask)
	app.Delete("/tasks/:id", deleteTask)
}
