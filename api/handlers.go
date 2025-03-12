package api

import (
	"testEx/internal"

	"github.com/gofiber/fiber/v2"
)

func createTask(c *fiber.Ctx) error {
	var task internal.Task
	if err := c.BodyParser(&task); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	id, err := internal.CreateTask(task)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to create task"})
	}

	return c.Status(201).JSON(fiber.Map{"id": id})
}

func getTasks(c *fiber.Ctx) error {
	// 1. Получаем задачи из базы данных
	tasks, err := internal.GetTasks()
	if err != nil {
		// Если произошла ошибка, возвращаем ошибку клиенту
		return c.Status(500).JSON(fiber.Map{"error": "Failed to fetch tasks"})
	}

	// 2. Возвращаем задачи клиенту в формате JSON
	return c.JSON(tasks)
}

func updateTask(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid task ID"})
	}

	type UpdateTask struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		Status      string `json:"status"`
	}

	var task UpdateTask
	if err := c.BodyParser(&task); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	err = internal.UpdateTask(id, internal.Task{
		Title:       task.Title,
		Description: task.Description,
		Status:      task.Status,
	})
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to update task"})
	}

	return c.Status(200).JSON(fiber.Map{"message": "Task updated successfully"})
}

func deleteTask(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid task ID"})
	}

	err = internal.DeleteTask(id)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to delete task"})
	}

	return c.Status(200).JSON(fiber.Map{"message": "Task deleted successfully"})
}
