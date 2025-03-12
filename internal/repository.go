package internal

import (
	"context"
	"log"

	"github.com/Chaykaman/testEx/entity"
	"github.com/jackc/pgx/v4/pgxpool"
)

var db *pgxpool.Pool

func InitDB(dbURL string) {
	var err error
	db, err = pgxpool.Connect(context.Background(), dbURL)
	if err != nil {
		log.Fatalf("Ошибка при подключении к базе данных: %v\n", err)
	}
	log.Println("Подключение к базе данных успешно установлено")
}

func GetTaskByID(id int) (Task, error) {
	var task Task

	// Логируем начало выполнения метода
	log.Printf("Попытка получить задачу с ID=%d\n", id)

	// Выполняем SQL-запрос для получения задачи по ID
	err := db.QueryRow(context.Background(), "SELECT id, title, description, status, created_at, updated_at FROM tasks WHERE id = $1", id).Scan(
		&task.ID, &task.Title, &task.Description, &task.Status, &task.CreatedAt, &task.UpdatedAt,
	)

	if err != nil {
		log.Printf("Ошибка при получении задачи с ID=%d: %v\n", id, err)
		return Task{}, entity.ErrNotFound
	}

	// Логируем успешное выполнение
	log.Printf("Задача с ID=%d успешно получена: %+v\n", id, task)

	return task, nil
}

func GetTasks() ([]Task, error) {
	// Логируем начало выполнения метода
	log.Println("Попытка получить список всех задач")

	rows, err := db.Query(context.Background(), "SELECT id, title, description, status, created_at, updated_at FROM tasks")
	if err != nil {
		log.Println("Ошибка при выполнении запроса:", err)
		return nil, entity.ErrRequestFailed
	}
	defer rows.Close()

	var tasks []Task
	for rows.Next() {
		var task Task
		err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.Status, &task.CreatedAt, &task.UpdatedAt)
		if err != nil {
			log.Println("Ошибка при сканировании строки:", err)
			return nil, entity.ErrRequestFailed
		}

		// Логируем каждую задачу
		log.Printf("Задача: ID=%d, Title=%s, Description=%s, Status=%s, CreatedAt=%s, UpdatedAt=%s\n",
			task.ID, task.Title, task.Description, task.Status, task.CreatedAt, task.UpdatedAt)

		tasks = append(tasks, task)
	}

	// Логируем общее количество задач
	log.Printf("Всего задач извлечено: %d\n", len(tasks))

	return tasks, nil
}

func CreateTask(task Task) (int, error) {
	// Логируем начало выполнения метода
	log.Printf("Попытка создать задачу: Title=%s, Description=%s\n", task.Title, task.Description)

	var id int
	err := db.QueryRow(context.Background(), "INSERT INTO tasks (title, description) VALUES ($1, $2) RETURNING id", task.Title, task.Description).Scan(&id)
	if err != nil {
		log.Printf("Ошибка при создании задачи: %v\n", err)
		return 0, entity.ErrTaskCreationFailed
	}

	// Логируем успешное создание задачи
	log.Printf("Задача успешно создана: ID=%d, Title=%s, Description=%s\n", id, task.Title, task.Description)

	return id, nil
}

func UpdateTask(id int, task Task) error {
	// Логируем начало выполнения метода
	log.Printf("Попытка обновить задачу с ID=%d: Title=%s, Description=%s, Status=%s\n", id, task.Title, task.Description, task.Status)

	_, err := db.Exec(context.Background(), "UPDATE tasks SET title = $1, description = $2, status = $3, updated_at = now() WHERE id = $4", task.Title, task.Description, task.Status, id)
	if err != nil {
		log.Printf("Ошибка при обновлении задачи с ID=%d: %v\n", id, err)
		return entity.ErrTaskUpdateFailed
	}

	// Логируем успешное обновление задачи
	log.Printf("Задача с ID=%d успешно обновлена\n", id)

	return nil
}

func DeleteTask(id int) error {
	// Логируем начало выполнения метода
	log.Printf("Попытка удалить задачу с ID=%d\n", id)

	_, err := db.Exec(context.Background(), "DELETE FROM tasks WHERE id = $1", id)
	if err != nil {
		log.Printf("Ошибка при удалении задачи с ID=%d: %v\n", id, err)
		return entity.ErrTaskDeleteFailed
	}

	// Логируем успешное удаление задачи
	log.Printf("Задача с ID=%d успешно удалена\n", id)

	return nil
}
