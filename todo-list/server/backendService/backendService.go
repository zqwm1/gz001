package backend

import (
	"time"
)

type Task struct {
	ID     string    `json:"_id"`
	Title  string    `json:"title"`
	Solved bool      `json:"solved"`
	Date   time.Time `json:"date"`
}

// genezio: deploy
type TaskService struct {
}

func New() TaskService {
	return TaskService{}
}

func (b TaskService) GetAllTasks() (Task, error) {

	return Task{
		ID:     "1",
		Title:  "2",
		Solved: false,
		Date:   time.Time{},
	}, nil

}
