package service

import (
	"crypto/sha256"
	"errors"
	"log"
	"zip-app/internal/repository"

	"github.com/google/uuid"
)

type ZipService struct {
	repo repository.ZipRepository
}

func NewService(repository repository.ZipRepository) ZipService {
	return ZipService{
		repo: repository,
	}
}

func (s *ZipService) CreateTask(userIP, userAgent string) (uuid.UUID, error) {
	hasher := sha256.New()
	hasher.Write([]byte(userIP + userAgent))
	userID := string(hasher.Sum(nil))

	return s.repo.CreateTask(userID)
}

func (s *ZipService) UpdateTask(taskID uuid.UUID, files []string) error {
	task, err := s.repo.DB.GetTask(taskID)
	if err != nil {
		return err
	}
	log.Println(len(task.Files))
	log.Println(task.Files)
	if len(task.Files) >= 3 {
		return errors.New("zip progress...")
	}
	s.repo.UpdateTask(taskID, files)
	return nil
}

func (s *ZipService) CheckStatus(taskID uuid.UUID) string {
	return s.repo.GetTaskStatus(taskID)
}