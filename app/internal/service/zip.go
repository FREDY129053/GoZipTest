package service

import (
	"crypto/sha256"
	"errors"
	"net/http"
	"zip-app/internal/models"
	"zip-app/internal/repository"
	"zip-app/pkg/helpers"

	"github.com/google/uuid"
)

type Answer struct {
	Code int
	Message any
	Err error
}

type ZipService struct {
	repo repository.ZipRepository
}

func NewService(repository repository.ZipRepository) ZipService {
	return ZipService{
		repo: repository,
	}
}

func (s *ZipService) CreateTask(userIP, userAgent string) Answer {
	hasher := sha256.New()
	hasher.Write([]byte(userIP + userAgent))
	userID := string(hasher.Sum(nil))

	resID, err := s.repo.CreateTask(userID)
	if err != nil {
		return Answer{
			Code: http.StatusLocked,
			Message: map[string]string{"error": err.Error()},
			Err: err,
		}
	}

	return Answer{
		Code: http.StatusCreated,
		Message: map[string]string{"task_id": resID.String()},
		Err: nil,
	}
}

func (s *ZipService) UpdateTask(taskID uuid.UUID, files []string) Answer {
	task, err := s.repo.GetTask(taskID)
	if err != nil {
		return Answer{
			Code: http.StatusNotFound,
			Message: map[string]string{"error": err.Error()},
			Err: err,
		}
	}
	
	// Замена файла
	for i, taskFile := range task.Files {
		for _, file := range files {
			if file == taskFile {
				task.Files[i] = file

				return Answer{
					Code: http.StatusOK,
					Message: map[string]string{"message": "file changed"},
					Err: nil,
				}
			}
		}
	}

	if len(task.Files) >= 3 {
		return Answer{
			Code: http.StatusUnprocessableEntity,
			Message: map[string]string{"error": "maximum count of files (3 files) in task reached and u can get archive"},
			Err: errors.New("max files count reached"),
		}
	}

	filesToAddCount := 3 - len(task.Files)
	if len(files) > filesToAddCount {
		files = files[:filesToAddCount]
	} 

	s.repo.UpdateTaskFiles(taskID, files)
	return Answer{
		Code: http.StatusOK,
		Message: map[string]string{"message": "files added"},
		Err: nil,
	}
}

func (s *ZipService) CheckStatus(taskID uuid.UUID) Answer {
	task, err := s.repo.GetTask(taskID)
	if err != nil {
		return Answer{
			Code: http.StatusNotFound,
			Message: map[string]string{"error": err.Error()},
			Err: err,
		}
	}

	if len(task.Files) == 3 && task.Status != models.Completed {
		task.Status = models.Completed
		failedFiles, archive := helpers.CreateArchive(task.Files, taskID.String())
		
		response := map[string]any{}
		response["archive_link"] = "http://localhost:8080/api/v1/zip_task/download/" + archive
		response["task_status"] = task.Status.String()

		if len(failedFiles) != 0 {
			response["failed_files"] = failedFiles
		}

		return Answer{
			Code: http.StatusOK,
			Message: response,
			Err: nil,
		}
	}

	// tFiles := []string{
	// 	"https://static1.makeuseofimages.com/wordpress/wp-content/uploads/2023/02/golang-logo-on-a-code-background.jpg",
	// 	"https://static2.makeuseofimages.com/wordpress/wp-content/uploads/2023/02/golang-logo-on-a-code-background.jpg",
	// 	"https://static3999.makeuseofimages.com/wordpress/wp-content/uploads/2023/02/golang-logo-on-a-code-background.jpg",
	// }
	
	return Answer{
		Code: http.StatusOK,
		Message: map[string]string{"task_status": task.Status.String()},
		Err: nil,
	}
}