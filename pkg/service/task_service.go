package service

import "task/pkg/repository"

type TaskService struct {
	repo repository.Task
}

func NewTaskService(repo repository.Task) *TaskService {
	return &TaskService{
		repo: repo,
	}
}
