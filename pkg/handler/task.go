package handler

import (
	"github.com/gin-gonic/gin"
	"task/pkg/service"
)

type TaskHandler struct {
	service service.Task
}

func NewTaskHandler(service service.Task) *TaskHandler {
	return &TaskHandler{
		service: service,
	}
}

// CreateTask @Summary create task
// @Tags task
// @Description create task
// @Id create-task
// @Params
// @Produce json
// @Success 200
// @Failure 400 {object} errorWeb
// @Failure 404 {object} errorWeb
// @Failure 500 {object} errorWeb
// @Failure default {object} errorWeb
// @Router /task [post]
func (h *TaskHandler) CreateTask(c *gin.Context) {}

// GetTask @Summary get task
// @Tags task
// @Description get task
// @Id get-task
// @Params
// @Produce json
// @Success 200
// @Failure 400 {object} errorWeb
// @Failure 404 {object} errorWeb
// @Failure 500 {object} errorWeb
// @Failure default {object} errorWeb
// @Router /task/:id [get]
func (h *TaskHandler) GetTask(c *gin.Context) {}

// UpdateTask @Summary update task
// @Tags task
// @Description update task
// @Id update-task
// @Params
// @Produce json
// @Success 200
// @Failure 400 {object} errorWeb
// @Failure 404 {object} errorWeb
// @Failure 500 {object} errorWeb
// @Failure default {object} errorWeb
// @Router /task [put]
func (h *TaskHandler) UpdateTask(c *gin.Context) {}

// DeleteTask @Summary delete task
// @Tags task
// @Description delete task
// @Id delete-task
// @Params
// @Produce json
// @Success 200
// @Failure 400 {object} errorWeb
// @Failure 404 {object} errorWeb
// @Failure 500 {object} errorWeb
// @Failure default {object} errorWeb
// @Router /task [delete]
func (h *TaskHandler) DeleteTask(c *gin.Context) {}

// GetTasksList @Summary get tasks list
// @Tags task
// @Description get tasks list
// @Id get-tasks-list
// @Params
// @Produce json
// @Success 200
// @Failure 400 {object} .errorWeb
// @Failure 404 {object} errorWeb
// @Failure 500 {object} errorWeb
// @Failure default {object} errorWeb
// @Router /task/list [get]
func (h *TaskHandler) GetTasksList(c *gin.Context) {}
