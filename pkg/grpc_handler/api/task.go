package api

import (
	"context"
	"fmt"
	"task/pkg/models"
	"task/pkg/service"
	"time"
)

type GRPCHandlerTask struct {
	service service.Task
}

func NewGRPCHandlerTask(service service.Task) *GRPCHandlerTask {
	return &GRPCHandlerTask{
		service: service,
	}
}

func (h *GRPCHandlerTask) CreateTask(ctx context.Context, in *TaskInput) (*TaskID, error) {
	parsedTime, err := time.Parse(time.RFC3339Nano, in.Date)
	if err != nil {
		err := fmt.Errorf("bad request: %s", err.Error())
		return nil, err
	}

	taskInput := models.TaskInput{
		IsDone:      &in.IsDone,
		Header:      in.Header,
		Description: in.Description,
		Date:        parsedTime,
	}

	taskId, err := h.service.CreateTask(taskInput)
	if err != nil {
		err := fmt.Errorf("server error: %s", err.Error())
		return nil, err
	}

	return &TaskID{ID: int32(taskId)}, nil
}

func (h *GRPCHandlerTask) GetTask(ctx context.Context, in *TaskID) (*Task, error) {
	task, err := h.service.GetTask(int(in.ID))
	if err != nil {
		err := fmt.Errorf("server error: %s", err.Error())
		return nil, err
	}

	return &Task{
		Id:          int32(task.Id),
		IsDone:      task.IsDone,
		Header:      task.Header,
		Description: task.Description,
		Date:        task.Date.String(),
	}, nil
}

func (h *GRPCHandlerTask) UpdateTask(ctx context.Context, in *Task) (*Task, error) {
	parsedTime, err := time.Parse(time.RFC3339Nano, in.Date)
	if err != nil {
		err := fmt.Errorf("bad request: %s", err.Error())
		return nil, err
	}

	taskInput := models.TaskInput{
		IsDone:      &in.IsDone,
		Header:      in.Header,
		Description: in.Description,
		Date:        parsedTime,
	}

	task, err := h.service.UpdateTask(int(in.Id), taskInput)
	if err != nil {
		err = fmt.Errorf("server error: %s", err.Error())
		return nil, err
	}

	return &Task{
		Id:          int32(task.Id),
		IsDone:      task.IsDone,
		Header:      task.Header,
		Description: task.Description,
		Date:        task.Date.String(),
	}, nil
}

func (h *GRPCHandlerTask) DeleteTask(ctx context.Context, in *TaskID) (*TaskID, error) {
	deletedTaskId, err := h.service.DeleteTask(int(in.ID))
	if err != nil {
		err = fmt.Errorf("server error: %s", err.Error())
		return nil, err
	}

	return &TaskID{
		ID: int32(deletedTaskId),
	}, nil
}

func (h *GRPCHandlerTask) GetTaskByDate(ctx context.Context, in *TaskByDate) (*TasksSlice, error) {
	parsedTime, err := time.Parse(time.RFC3339Nano, in.Date)
	if err != nil {
		err := fmt.Errorf("bad request: %s", err.Error())
		return nil, err
	}

	tasks, err := h.service.GetTasksByDate(parsedTime, in.IsDone)
	if err != nil {
		err = fmt.Errorf("server error: %s", err.Error())
		return nil, err
	}

	var tasksSlice *TasksSlice
	for _, task := range tasks {
		tasksSlice.Tasks = append(tasksSlice.Tasks, &Task{
			Id:          int32(task.Id),
			IsDone:      task.IsDone,
			Header:      task.Header,
			Description: task.Description,
			Date:        task.Date.String(),
		})
	}

	return tasksSlice, nil
}

func (h *GRPCHandlerTask) GetTasksList(ctx context.Context, in *EmptyRequest) (*TasksSlice, error) {
	tasks, err := h.service.GetTasksList()
	if err != nil {
		err = fmt.Errorf("server error: %s", err.Error())
		return nil, err
	}

	var tasksSlice *TasksSlice
	for _, task := range tasks {
		tasksSlice.Tasks = append(tasksSlice.Tasks, &Task{
			Id:          int32(task.Id),
			IsDone:      task.IsDone,
			Header:      task.Header,
			Description: task.Description,
			Date:        task.Date.String(),
		})
	}

	return tasksSlice, nil
}

func (h *GRPCHandlerTask) GetTasksByPage(ctx context.Context, in *PageRequest) (*TasksSlice, error) {
	tasks, err := h.service.GetTasksByPage(in.IsDone, int(in.Page))
	if err != nil {
		err = fmt.Errorf("server error: %s", err.Error())
		return nil, err
	}

	var tasksSlice *TasksSlice
	for _, task := range tasks {
		tasksSlice.Tasks = append(tasksSlice.Tasks, &Task{
			Id:          int32(task.Id),
			IsDone:      task.IsDone,
			Header:      task.Header,
			Description: task.Description,
			Date:        task.Date.String(),
		})
	}

	return tasksSlice, nil
}

func (h *GRPCHandlerTask) mustEmbedUnimplementedTaskServer() {}
