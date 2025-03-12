package api

import (
	"testEx/internal"

	"github.com/gofiber/fiber/v2"
)

func createTask(c *fiber.Ctx) error {
	var task internal.Task
	if err := c.BodyParser(&task); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Невозможно проанализировать JSON"})
	}

	id, err := internal.CreateTask(task)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Запрос не создался"})
	}

	return c.Status(201).JSON(fiber.Map{"id": id})
}

func getTasks(c *fiber.Ctx) error {
	// 1. Получаем задачи из базы данных
	tasks, err := internal.GetTasks()
	if err != nil {
		// Если произошла ошибка, возвращаем ошибку клиенту
		return c.Status(500).JSON(fiber.Map{"error": "Не удалось получить задачи"})
	}

	// 2. Возвращаем задачи клиенту в формате JSON
	return c.JSON(tasks)
}

func getTaskByID(c *fiber.Ctx) error {
	// 1. Извлекаем ID задачи из параметров URL
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Неверный идентификатор задачи"})
	}

	// 2. Получаем задачу по ID из базы данных
	task, err := internal.GetTaskByID(id)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Не удалось получить задачу"})
	}

	// 3. Возвращаем задачу в формате JSON
	return c.JSON(task)
}

func updateTask(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Неверный идентификатор задачи"})
	}

	type UpdateTask struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		Status      string `json:"status"`
	}

	var task UpdateTask
	if err := c.BodyParser(&task); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Невозможно проанализировать JSON"})
	}

	err = internal.UpdateTask(id, internal.Task{
		Title:       task.Title,
		Description: task.Description,
		Status:      task.Status,
	})
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Не удалось обновить задачу"})
	}

	return c.Status(200).JSON(fiber.Map{"message": "Задача успешно обновлена"})
}

func deleteTask(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Неверный ID"})
	}

	err = internal.DeleteTask(id)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Ошибка удаления"})
	}

	return c.Status(200).JSON(fiber.Map{"message": "Успешно удалён"})
}
