package database

import (
	"errors"
	"zip-app/internal/models"

	"github.com/google/uuid"
)

type Record = map[string][]*models.Task

type LocalDB struct {
	Records Record
}

func NewDatabase() LocalDB {
	return LocalDB{Records: make(Record)}
}

func (d *LocalDB) GetTask(id uuid.UUID) (*models.Task, error) {
	for _, tasks := range d.Records {
		for _, task := range tasks {
			if task.ID == id {
				return task, nil
			}
		}
	}

	return nil, errors.New("task not found")
}