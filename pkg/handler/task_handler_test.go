package handler


import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"net/http/httptest"
	"os"
	"task/pkg/models"
	"task/pkg/repository"
	"task/pkg/service"
	"testing"
	"time"
)

func initConfig() error {
	viper.AddConfigPath("../../config")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

func initializeHandler() (*Controller, error) {
	if err := initConfig(); err != nil {
		return nil, err
	}

	if err := godotenv.Load(); err != nil {
		return nil, err
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})
	if err != nil {
		return nil, err
	}
	repos := repository.NewRepository(db)
	service := service.NewService(repos)
	handler := NewController(service)

	return handler, nil
}

func TestCreateTask(t *testing.T) {
	//handler, err := initializeHandler()
	/*if err != nil {
		t.Errorf("can't initialize handler: %s", err.Error())
	}*/
	/*body := `
	{
        "is_done": true,
        "header": "уборка по дому",
        "description": "убрать стол",
        "date": "2022-10-25T00:00:00Z"
    }`*/
	isDone := true
	body := models.TaskInput{
		Header: "уборка по дому",
		IsDone: &isDone,
		Description: "убрать стол",
		Date: time.Now(),
	}
	data, err := json.Marshal(body){
		t.Errorf("can't prepare data: %s", err.Error())
	}
	r := httptest.NewRequest("GET", "http://localhost:80/user?id=42", bytes.NewBuffer(data))
	w := httptest.NewRecorder()
	gin.
		handler.Task.CreateTask(gin.Context{Request: r, Writer: w})
	user := User{}
	json.Unmarshal(w.Body.Bytes(), &user)
	if user.Id != 42 {
		t.Errorf("Invalid user id %d expected %d", user.Id, 42)
	}
}

func TestUserHandler(t *testing.T) {
	// Создаем объект *gin.Context для тестирования
	router := gin.Default()
	req := httptest.NewRequest("GET", "/user?id=42", nil)
	w := httptest.NewRecorder()
	c := router.NewContext(req, w)
	c.Set("id", "42")

	// Вызываем обработчик с созданным объектом *gin.Context
	userHandler(c)

	// Проверяем результаты работы обработчика
	user := User{}
	json.Unmarshal(w.Body.Bytes(), &user)
	if user.Id != 42 {
		t.Errorf("Invalid user id %d expected %d", user.Id, 42)
	}
}
*/