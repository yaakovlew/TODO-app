package service

import "task/pkg/repository"

type Task interface{}

type Service struct {
	Task
}

func NewService(repo *repository.Repo) *Service {
	return &Service{
		Task: NewTaskService(repo.Task),
	}
}
