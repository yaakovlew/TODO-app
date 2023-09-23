package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"task/pkg/models"
	"task/pkg/service"
	"time"
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
// @Param input body models.TaskInput true "Task input"
// @Accept json
// @Produce json
// @Success 200 {integer} integer 1
// @Failure 400 {object} errorWeb
// @Failure 404 {object} errorWeb
// @Failure 500 {object} errorWeb
// @Failure default {object} errorWeb
// @Router /task [post]
func (h *TaskHandler) CreateTask(c *gin.Context) {
	var taskInput models.TaskInput
	if err := c.BindJSON(&taskInput); err != nil {
		err := fmt.Errorf("incorrect data: %s", err.Error())
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	taskId, err := h.service.CreateTask(taskInput)
	if err != nil {
		err := fmt.Errorf("server error: %s", err.Error())
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, models.TaskIdResponse{
		TaskId: taskId,
	})
}

// GetTask @Summary get task
// @Tags task
// @Description get task
// @Id get-task
// @Param id path int true "Task ID"
// @Produce json
// @Success 200 {object} models.Task
// @Failure 400 {object} errorWeb
// @Failure 404 {object} errorWeb
// @Failure 500 {object} errorWeb
// @Failure default {object} errorWeb
// @Router /task/{id} [get]
func (h *TaskHandler) GetTask(c *gin.Context) {
	taskId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		err := fmt.Errorf("incorrect data: %s", err.Error())
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	task, err := h.service.GetTask(taskId)
	if err != nil {
		err := fmt.Errorf("server error: %s", err.Error())
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, task)
}

// UpdateTask @Summary update task
// @Tags task
// @Description update task
// @Id update-task
// @Param id path int true "Task ID"
// @Param input body models.TaskInput true "update task input"
// @Accept json
// @Produce json
// @Success 200 {object} models.TaskInput
// @Failure 400 {object} errorWeb
// @Failure 404 {object} errorWeb
// @Failure 500 {object} errorWeb
// @Failure default {object} errorWeb
// @Router /task/{id} [put]
func (h *TaskHandler) UpdateTask(c *gin.Context) {
	taskId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		err := fmt.Errorf("incorrect data: %s", err.Error())
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	var taskInput models.TaskInput
	if err := c.BindJSON(&taskInput); err != nil {
		err := fmt.Errorf("incorrect data: %s", err.Error())
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	task, err := h.service.UpdateTask(taskId, taskInput)
	if err != nil {
		err := fmt.Errorf("server error: %s", err.Error())
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, task)
}

// DeleteTask @Summary delete task
// @Tags task
// @Description delete task
// @Id delete-task
// @Param id path int true "Task ID"
// @Produce json
// @Success 200 {int} id
// @Failure 400 {object} errorWeb
// @Failure 404 {object} errorWeb
// @Failure 500 {object} errorWeb
// @Failure default {object} errorWeb
// @Router /task/{id} [delete]
func (h *TaskHandler) DeleteTask(c *gin.Context) {
	taskId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		err := fmt.Errorf("incorrect data: %s", err.Error())
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	deletedTaskId, err := h.service.DeleteTask(taskId)
	if err != nil {
		err := fmt.Errorf("server error: %s", err.Error())
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"deleted_task_id": deletedTaskId,
	})
}

// GetTasksByDate @Summary get tasks list by date
// @Tags task
// @Description get tasks list
// @Id get-tasks-list-date
// @Param is_done query boolean true "Is_Done"
// @Param date query string true "Date"
// @Produce json
// @Success 200
// @Failure 400 {object} .errorWeb
// @Failure 404 {object} errorWeb
// @Failure 500 {object} errorWeb
// @Failure default {object} errorWeb
// @Router /task/list/date [get]
func (h *TaskHandler) GetTasksByDate(c *gin.Context) {
	dateParam := c.Query("date")
	date, err := time.Parse("2006-01-02", dateParam)
	if err != nil {
		err := fmt.Errorf("incorrect flag: %s", err.Error())
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	isDone, err := strconv.ParseBool(c.Query("is_done"))
	if err != nil {
		err := fmt.Errorf("incorrect flag: %s", err.Error())
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	tasks, err := h.service.GetTasksByDate(date, isDone)
	if err != nil {
		err := fmt.Errorf("server error: %s", err.Error())
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, tasks)
}

// GetTasksList @Summary get tasks list
// @Tags task
// @Description get tasks list
// @Id get-tasks-list
// @Produce json
// @Success 200 {object} []models.Task
// @Failure 400 {object} .errorWeb
// @Failure 404 {object} errorWeb
// @Failure 500 {object} errorWeb
// @Failure default {object} errorWeb
// @Router /task/list [get]
func (h *TaskHandler) GetTasksList(c *gin.Context) {
	tasks, err := h.service.GetTasksList()
	if err != nil {
		err := fmt.Errorf("server error: %s", err.Error())
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, tasks)
}

// GetTasksByPage @Summary get tasks list
// @Tags task
// @Description get tasks list
// @Id get-tasks-list-page
// @Param page query integer true "Page"
// @Param is_done query boolean true "Is_Done"
// @Produce json
// @Success 200
// @Failure 400 {object} .errorWeb
// @Failure 404 {object} errorWeb
// @Failure 500 {object} errorWeb
// @Failure default {object} errorWeb
// @Router /task/list/page [get]
func (h *TaskHandler) GetTasksByPage(c *gin.Context) {
	pageRead := c.Query("page")
	page, err := strconv.Atoi(pageRead)
	if err != nil {
		err := fmt.Errorf("incorrect page: %s", err.Error())
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	isDone, err := strconv.ParseBool(c.Query("is_done"))
	if err != nil {
		err := fmt.Errorf("incorrect flag: %s", err.Error())
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	tasks, err := h.service.GetTasksByPage(isDone, page)
	if err != nil {
		err := fmt.Errorf("server error: %s", err.Error())
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"tasks": tasks,
	})
}
