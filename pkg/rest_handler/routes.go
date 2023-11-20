package rest_handler

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "task/docs"
	"task/pkg/service"
)

type Controller struct {
	Task *TaskHandler
}

func NewController(service *service.Service) *Controller {
	return &Controller{
		Task: NewTaskHandler(service.Task),
	}
}

func (c *Controller) InitRoutes() *gin.Engine {
	router := gin.New()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	task := router.Group("/task")
	{
		task.GET("/:id", c.Task.GetTask)
		task.POST("", c.Task.CreateTask)
		task.DELETE("/:id", c.Task.DeleteTask)
		task.PUT("/:id", c.Task.UpdateTask)

		list := task.Group("/list")
		{
			list.GET("", c.Task.GetTasksList)
			list.GET("/date", c.Task.GetTasksByDate)
			list.GET("/page", c.Task.GetTasksByPage)
		}
	}

	return router
}
