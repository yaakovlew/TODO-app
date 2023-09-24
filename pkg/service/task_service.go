package service

import (
	"fmt"
	"task/pkg/models"
	"task/pkg/repository"
	"time"
)

type TaskService struct {
	repo repository.Task
}

func NewTaskService(repo repository.Task) *TaskService {
	return &TaskService{
		repo: repo,
	}
}

func (s *TaskService) GetTask(id int) (models.Task, error) {
	if id == 0 {
		return models.Task{}, fmt.Errorf("not found")
	}
	return s.repo.GetTask(id)
}

func (s *TaskService) CreateTask(task models.TaskInput) (int, error) {
	return s.repo.CreateTask(task)
}

func (s *TaskService) DeleteTask(id int) (int, error) {
	return s.repo.DeleteTask(id)
}

func (s *TaskService) UpdateTask(id int, task models.TaskInput) (models.Task, error) {
	return s.repo.UpdateTask(id, task)
}

func (s *TaskService) GetTasksList() ([]models.Task, error) {
	return s.repo.GetTasksList()
}

func (s *TaskService) GetTasksByDate(date time.Time, isDone bool) ([]models.Task, error) {
	return s.repo.GetTasksByDate(date, isDone)
}

func (s *TaskService) GetTasksByPage(isDone bool, page int) ([]models.Task, error) {
	limitForPage := 2
	return s.repo.GetTasksByPage(isDone, limitForPage, page)
}
