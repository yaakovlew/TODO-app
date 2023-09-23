package repository

import (
	"github.com/jmoiron/sqlx"
	"task/pkg/models"
	"time"
)

type Task interface {
	GetTask(id int) (models.Task, error)
	CreateTask(task models.TaskInput) (int, error)
	DeleteTask(id int) (int, error)
	UpdateTask(id int, task models.TaskInput) (models.TaskInput, error)
	GetTasksList() ([]models.Task, error)
	GetTasksByPage(isDone bool, limit, page int) ([]models.Task, error)
	GetTasksByDate(date time.Time, isDone bool) ([]models.Task, error)
}

type Repo struct {
	Task
}

func NewRepository(db *sqlx.DB) *Repo {
	return &Repo{
		Task: NewTaskRepository(db),
	}
}
