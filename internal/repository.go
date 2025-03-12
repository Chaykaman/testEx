package internal

import (
	"context"
	"log"

	"github.com/jackc/pgx/v4/pgxpool"
)

var db *pgxpool.Pool

func InitDB(dbURL string) {
	var err error
	db, err = pgxpool.Connect(context.Background(), dbURL)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
}

func GetTasks() ([]Task, error) {
	rows, err := db.Query(context.Background(), "SELECT id, title, description, status, created_at, updated_at FROM tasks")
	if err != nil {
		log.Println("Ошибка при выполнении запроса:", err)
		return nil, err
	}
	defer rows.Close()

	var tasks []Task
	for rows.Next() {
		var task Task
		err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.Status, &task.CreatedAt, &task.UpdatedAt)
		if err != nil {
			log.Println("Ошибка при сканировании строки:", err)
			return nil, err
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
	var id int
	err := db.QueryRow(context.Background(), "INSERT INTO tasks (title, description) VALUES ($1, $2) RETURNING id", task.Title, task.Description).Scan(&id)
	if err != nil {
		return 0, err
	}
	// Логируем созданную задачу
	log.Printf("Создана задача: ID=%d, Title=%s, Description=%s\n", id, task.Title, task.Description)
	return id, nil
}

func UpdateTask(id int, task Task) error {
	_, err := db.Exec(context.Background(), "UPDATE tasks SET title = $1, description = $2, status = $3, updated_at = now() WHERE id = $4", task.Title, task.Description, task.Status, id)
	return err
}

func DeleteTask(id int) error {
	_, err := db.Exec(context.Background(), "DELETE FROM tasks WHERE id = $1", id)
	return err
}
