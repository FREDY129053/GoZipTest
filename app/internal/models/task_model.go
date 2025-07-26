package models

import "github.com/google/uuid"

type TaskState int

const (
	InProgress TaskState = iota
	Completed
)

var strState = map[TaskState]string{
	InProgress: "in_progress",
	Completed:  "completed",
}

func (state TaskState) String() string {
	return strState[state]
}

type Task struct {
	ID     uuid.UUID `json:"id"`
	Status TaskState `json:"status"`
	Files  []string  `json:"files"`
}

func (t *Task) BeforeCreate() {
	t.ID, _ = uuid.NewRandom()
	t.Status = InProgress
}
