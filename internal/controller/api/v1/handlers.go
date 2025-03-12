package v1

import (
	"log"

	"github.com/Chaykaman/testEx/entity"
	"github.com/Chaykaman/testEx/internal"
	"github.com/gofiber/fiber/v2"
)

func createTask(c *fiber.Ctx) error {
	// Логируем начало обработки запроса
	log.Println("Обработка запроса на создание задачи")

	var task internal.Task
	if err := c.BodyParser(&task); err != nil {
		log.Printf("Ошибка при парсинге JSON: %v\n", err)
		return c.Status(400).JSON(fiber.Map{"error": entity.ErrJSONParseFailed.Error()})
	}

	// Логируем данные задачи
	log.Printf("Данные задачи: Title=%s, Description=%s\n", task.Title, task.Description)

	id, err := internal.CreateTask(task)
	if err != nil {
		log.Printf("Ошибка при создании задачи: %v\n", err)
		return c.Status(500).JSON(fiber.Map{"error": entity.ErrTaskCreationFailed.Error()})
	}

	// Логируем успешное создание задачи
	log.Printf("Задача успешно создана: ID=%d\n", id)

	return c.Status(201).JSON(fiber.Map{"id": id})
}

func getTasks(c *fiber.Ctx) error {
	// Логируем начало обработки запроса
	log.Println("Обработка запроса на получение списка задач")

	// 1. Получаем задачи из базы данных
	tasks, err := internal.GetTasks()
	if err != nil {
		log.Printf("Ошибка при получении задач: %v\n", err)
		return c.Status(500).JSON(fiber.Map{"error": entity.ErrTaskFetchFailed.Error()})
	}

	// Логируем количество полученных задач
	log.Printf("Получено задач: %d\n", len(tasks))

	// 2. Возвращаем задачи клиенту в формате JSON
	return c.JSON(tasks)
}

func getTaskByID(c *fiber.Ctx) error {
	// Логируем начало обработки запроса
	log.Println("Обработка запроса на получение задачи по ID")

	// 1. Извлекаем ID задачи из параметров URL
	id, err := c.ParamsInt("id")
	if err != nil {
		log.Printf("Неверный идентификатор задачи: %v\n", err)
		return c.Status(400).JSON(fiber.Map{"error": entity.ErrInvalidTaskID.Error()})
	}

	// Логируем ID задачи
	log.Printf("Запрошена задача с ID=%d\n", id)

	// 2. Получаем задачу по ID из базы данных
	task, err := internal.GetTaskByID(id)
	if err != nil {
		log.Printf("Ошибка при получении задачи с ID=%d: %v\n", id, err)
		return c.Status(500).JSON(fiber.Map{"error": entity.ErrTaskFetchFailed.Error()})
	}

	// Логируем успешное получение задачи
	log.Printf("Задача с ID=%d успешно получена: %+v\n", id, task)

	// 3. Возвращаем задачу в формате JSON
	return c.JSON(task)
}

func updateTask(c *fiber.Ctx) error {
	// Логируем начало обработки запроса
	log.Println("Обработка запроса на обновление задачи")

	// 1. Извлекаем ID задачи из параметров URL
	id, err := c.ParamsInt("id")
	if err != nil {
		log.Printf("Неверный идентификатор задачи: %v\n", err)
		return c.Status(400).JSON(fiber.Map{"error": entity.ErrInvalidTaskID.Error()})
	}

	// Логируем ID задачи
	log.Printf("Обновление задачи с ID=%d\n", id)

	type UpdateTask struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		Status      string `json:"status"`
	}

	var task UpdateTask
	if err := c.BodyParser(&task); err != nil {
		log.Printf("Ошибка при парсинге JSON: %v\n", err)
		return c.Status(400).JSON(fiber.Map{"error": entity.ErrJSONParseFailed.Error()})
	}

	// Логируем данные для обновления
	log.Printf("Данные для обновления: Title=%s, Description=%s, Status=%s\n", task.Title, task.Description, task.Status)

	err = internal.UpdateTask(id, internal.Task{
		Title:       task.Title,
		Description: task.Description,
		Status:      task.Status,
	})
	if err != nil {
		log.Printf("Ошибка при обновлении задачи с ID=%d: %v\n", id, err)
		return c.Status(500).JSON(fiber.Map{"error": entity.ErrTaskUpdateFailed.Error()})
	}

	// Логируем успешное обновление задачи
	log.Printf("Задача с ID=%d успешно обновлена\n", id)

	return c.Status(200).JSON(fiber.Map{"message": "Задача успешно обновлена"})
}

func deleteTask(c *fiber.Ctx) error {
	// Логируем начало обработки запроса
	log.Println("Обработка запроса на удаление задачи")

	// 1. Извлекаем ID задачи из параметров URL
	id, err := c.ParamsInt("id")
	if err != nil {
		log.Printf("Неверный идентификатор задачи: %v\n", err)
		return c.Status(400).JSON(fiber.Map{"error": entity.ErrInvalidTaskID.Error()})
	}

	// Логируем ID задачи
	log.Printf("Удаление задачи с ID=%d\n", id)

	err = internal.DeleteTask(id)
	if err != nil {
		log.Printf("Ошибка при удалении задачи с ID=%d: %v\n", id, err)
		return c.Status(500).JSON(fiber.Map{"error": entity.ErrTaskDeleteFailed.Error()})
	}

	// Логируем успешное удаление задачи
	log.Printf("Задача с ID=%d успешно удалена\n", id)

	return c.Status(200).JSON(fiber.Map{"message": "Успешно удалён"})
}
