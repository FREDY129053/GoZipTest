package repository

import (
	"errors"
	"log"
	"zip-app/internal/database"
	"zip-app/internal/models"

	"github.com/google/uuid"
)

type ZipRepository struct {
	DB database.LocalDB
}

func NewRepository(db database.LocalDB) ZipRepository {
	return ZipRepository{
		DB: db,
	}
}

func (r *ZipRepository) CreateTask(userID string) (uuid.UUID, error) {
	counter := 0
	for _, tasks := range r.DB.Records {
		for _, task := range tasks {
			if task.Status == models.InProgress {
				counter++
			}
			if counter == 3 {
				log.Println("3 in progress!")
				return uuid.Nil, errors.New("3 tasks in progress")
			}
		}
	}

	// Создаем задачу с дефолтными значениями
	task := models.Task{}
	task.BeforeCreate()

	// Добавляем в БД запись
	r.DB.Records[userID] = append(r.DB.Records[userID], &task)

	return task.ID, nil
}

func (r *ZipRepository) UpdateTaskFiles(id uuid.UUID, files []string) error {
	for _, tasks := range r.DB.Records {
		for _, task := range tasks {
			if id == task.ID {
				task.Files = append(task.Files, files...)
				return nil
			}
		}
	}

	return errors.New("task not found")
}

func (r *ZipRepository) GetTask(id uuid.UUID) (*models.Task, error) {
	return r.DB.GetTask(id)
}