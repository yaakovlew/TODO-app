package service

import (
	"task/pkg/models"
	"task/pkg/repository"
	"time"
)

type Task interface {
	GetTask(id int) (models.Task, error)
	CreateTask(task models.TaskInput) (int, error)
	DeleteTask(id int) (int, error)
	UpdateTask(id int, task models.TaskInput) (models.TaskInput, error)
	GetTasksList() ([]models.Task, error)
	GetTasksByDate(date time.Time, isDone bool) ([]models.Task, error)
	GetTasksByPage(isDone bool, page int) ([]models.Task, error)
}

type Service struct {
	Task
}

func NewService(repo *repository.Repo) *Service {
	return &Service{
		Task: NewTaskService(repo.Task),
	}
}
