package schemas

type APIError struct {
	Error string `json:"error"`
}

type CreatedTask struct {
	TaskID string `json:"task_id" example:"b841515d-476e-4eb7-956e-5a6976d7f026"`
}

type MessageAnswer struct {
	Message string `json:"message"`
}

type TaskStatus struct {
	TaskStatus  string   `json:"task_status" example:"in_progress" enums:"in_progress,completed"`
	FailedFiles []string `json:"failed_files" format:"nullable" example:"file1.jpg,file2.pdf"`
	ArchiveLink string   `json:"archive_link" swaggertype:"string" format:"nullable"`
}
