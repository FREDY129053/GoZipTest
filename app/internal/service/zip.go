package service

import (
	"crypto/sha256"
	"errors"
	"log"
	"zip-app/internal/repository"
	"zip-app/pkg/helpers"

	"github.com/google/uuid"
)

type Answer struct {

}

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
	task, err := s.repo.GetTask(taskID)
	if err != nil {
		return err
	}
	
	if len(task.Files) >= 3 {
		return errors.New("zip progress...")
	}

	s.repo.UpdateTask(taskID, files)
	return nil
}

func (s *ZipService) CheckStatus(taskID uuid.UUID) (*string, error) {
	task, err := s.repo.GetTask(taskID)
	if err != nil {
		return nil, err
	}

	if len(task.Files) == 3 {
		_, archive := helpers.CreateArchive(task.Files, "archive")
		return &archive, nil
	}

	// tFiles := []string{
	// 	"https://static1.makeuseofimages.com/wordpress/wp-content/uploads/2023/02/golang-logo-on-a-code-background.jpg",
	// 	"https://static2.makeuseofimages.com/wordpress/wp-content/uploads/2023/02/golang-logo-on-a-code-background.jpg",
	// 	"https://static3999.makeuseofimages.com/wordpress/wp-content/uploads/2023/02/golang-logo-on-a-code-background.jpg",
	// }
	
	status := task.Status.String()
	return &status, nil
}