package handler

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	_ "github.com/lib/pq"
	"net/http/httptest"
	"task/pkg/models"
	service2 "task/pkg/service"
	mock_service "task/pkg/service/mocks"
	"testing"
	"time"
)

func TestTaskHandler_CreateTask(t *testing.T) {
	type mockBehaviour func(s *mock_service.MockTask, input models.TaskInput)

	isDone := false
	date, err := time.Parse("2006-01-02", "2023-09-23")
	if err != nil {
		t.Errorf("can't prepare data for testing: %s", err.Error())
	}

	testsTable := []struct {
		name                string
		inputBody           string
		task                models.TaskInput
		mockBehaviour       mockBehaviour
		expectedStatusCode  int
		expectedRequestBody string
	}{
		{
			name:      "OK",
			inputBody: `{"is_done": false, "header": "убраться в комнате", "description": "помыть посуду", "date": "2023-09-23T00:00:00Z"}`,
			task: models.TaskInput{
				IsDone:      &isDone,
				Header:      "убраться в комнате",
				Description: "помыть посуду",
				Date:        date,
			},
			mockBehaviour: func(s *mock_service.MockTask, task models.TaskInput) {
				s.EXPECT().CreateTask(task).Return(1, nil)
			},
			expectedStatusCode:  200,
			expectedRequestBody: `{"task_id":1}`,
		},
		{
			name:      "Empty fields",
			inputBody: `{"is_done": false, "description": "помыть посуду", "date": "2023-09-23T00:00:00Z"}`,
			task: models.TaskInput{
				IsDone:      &isDone,
				Description: "помыть посуду",
				Date:        date,
			},
			mockBehaviour:       func(s *mock_service.MockTask, task models.TaskInput) {},
			expectedStatusCode:  400,
			expectedRequestBody: `{"message":"incorrect data: Key: 'TaskInput.Header' Error:Field validation for 'Header' failed on the 'required' tag"}`,
		},
	}

	for _, testCase := range testsTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			task := mock_service.NewMockTask(c)
			testCase.mockBehaviour(task, testCase.task)

			service := service2.Service{Task: task}
			handler := NewTaskHandler(service)

			r := gin.New()
			r.POST("/task", handler.CreateTask)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/task", bytes.NewBufferString(testCase.inputBody))

			r.ServeHTTP(w, req)

			if testCase.expectedStatusCode != w.Code {
				t.Error("status codes is not equal")
			}

			if testCase.expectedRequestBody != w.Body.String() {
				t.Error("requests is not equal")
			}
		})
	}
}

func TestTaskHandler_DeleteTask(t *testing.T) {
	type mockBehaviour func(s *mock_service.MockTask, taskId int)

	testsTable := []struct {
		name                string
		TaskId              int
		mockBehaviour       mockBehaviour
		expectedStatusCode  int
		expectedRequestBody string
	}{
		{
			name:   "OK",
			TaskId: 1,
			mockBehaviour: func(s *mock_service.MockTask, taskId int) {
				s.EXPECT().DeleteTask(taskId).Return(1, nil)
			},
			expectedStatusCode:  200,
			expectedRequestBody: `{"task_id":1}`,
		},
		{
			name:   "Internal Server Error",
			TaskId: 0,
			mockBehaviour: func(s *mock_service.MockTask, taskId int) {
				s.EXPECT().DeleteTask(taskId).Return(0, errors.New("not found"))
			},
			expectedStatusCode:  500,
			expectedRequestBody: `{"message":"server error: not found"}`,
		},
	}

	for _, testCase := range testsTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			task := mock_service.NewMockTask(c)
			testCase.mockBehaviour(task, testCase.TaskId)

			service := service2.Service{Task: task}
			handler := NewTaskHandler(service)

			r := gin.New()
			r.DELETE("/task/:id", handler.DeleteTask)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("DELETE", fmt.Sprintf("/task/%d", testCase.TaskId), nil)

			r.ServeHTTP(w, req)

			if testCase.expectedStatusCode != w.Code {
				t.Error("status codes is not equal")
			}
			if testCase.expectedRequestBody != w.Body.String() {
				t.Error("requests is not equal")
			}
		})
	}
}

func TestTaskHandler_UpdateTask(t *testing.T) {
	type mockBehaviour func(s *mock_service.MockTask, taskId int, task models.TaskInput)

	isDone := false
	date, err := time.Parse("2006-01-02", "2023-09-23")
	if err != nil {
		t.Errorf("can't prepare data for testing: %s", err.Error())
	}

	testsTable := []struct {
		name                string
		TaskId              int
		inputBody           string
		TaskInput           models.TaskInput
		mockBehaviour       mockBehaviour
		expectedStatusCode  int
		expectedRequestBody string
	}{
		{
			name:      "OK",
			TaskId:    1,
			inputBody: `{"is_done": false, "header": "убраться в комнате", "description": "помыть посуду", "date": "2023-09-23T00:00:00Z"}`,
			mockBehaviour: func(s *mock_service.MockTask, taskId int, task models.TaskInput) {
				s.EXPECT().UpdateTask(taskId, task).Return(models.Task{
					Id:          taskId,
					Header:      task.Header,
					Description: task.Description,
					Date:        task.Date,
					IsDone:      *task.IsDone,
				}, nil)
			},
			TaskInput: models.TaskInput{
				IsDone:      &isDone,
				Header:      "убраться в комнате",
				Description: "помыть посуду",
				Date:        date,
			},
			expectedStatusCode:  200,
			expectedRequestBody: `{"is_done":false,"id":1,"header":"убраться в комнате","description":"помыть посуду","date":"2023-09-23T00:00:00Z"}`,
		},
		{
			name:                "Empty Fields",
			TaskId:              1,
			inputBody:           `{"header": "убраться в комнате", "description": "помыть посуду", "date": "2023-09-23T00:00:00Z"}`,
			mockBehaviour:       func(s *mock_service.MockTask, taskId int, task models.TaskInput) {},
			expectedStatusCode:  400,
			expectedRequestBody: `{"message":"incorrect data: Key: 'TaskInput.IsDone' Error:Field validation for 'IsDone' failed on the 'required' tag"}`,
		},
	}

	for _, testCase := range testsTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			task := mock_service.NewMockTask(c)
			testCase.mockBehaviour(task, testCase.TaskId, testCase.TaskInput)

			service := service2.Service{Task: task}
			handler := NewTaskHandler(service)

			r := gin.New()
			r.PUT("/task/:id", handler.UpdateTask)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("PUT", fmt.Sprintf("/task/%d", testCase.TaskId), bytes.NewBufferString(testCase.inputBody))

			r.ServeHTTP(w, req)

			if testCase.expectedStatusCode != w.Code {
				t.Error("status codes is not equal")
			}
			if testCase.expectedRequestBody != w.Body.String() {
				t.Error("requests is not equal")
			}
		})
	}
}

func TestTaskHandler_GetTask(t *testing.T) {
	type mockBehaviour func(s *mock_service.MockTask, taskId int)

	date, err := time.Parse("2006-01-02", "2023-09-23")
	if err != nil {
		t.Errorf("can't prepare data for testing: %s", err.Error())
	}

	testsTable := []struct {
		name                string
		inputBody           string
		taskId              int
		mockBehaviour       mockBehaviour
		expectedStatusCode  int
		expectedRequestBody string
	}{
		{
			name:   "OK",
			taskId: 1,
			mockBehaviour: func(s *mock_service.MockTask, taskId int) {
				s.EXPECT().GetTask(taskId).Return(models.Task{
					Id:          1,
					IsDone:      false,
					Header:      "убраться в комнате",
					Description: "помыть посуду",
					Date:        date,
				}, nil)
			},
			expectedStatusCode:  200,
			expectedRequestBody: `{"is_done":false,"id":1,"header":"убраться в комнате","description":"помыть посуду","date":"2023-09-23T00:00:00Z"}`,
		},
		{
			name:   "Not Valid id",
			taskId: 0,
			mockBehaviour: func(s *mock_service.MockTask, taskId int) {
				s.EXPECT().GetTask(taskId).Return(models.Task{}, fmt.Errorf("not found"))
			},
			expectedStatusCode:  500,
			expectedRequestBody: `{"message":"server error: not found"}`,
		},
	}

	for _, testCase := range testsTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			task := mock_service.NewMockTask(c)
			testCase.mockBehaviour(task, testCase.taskId)

			service := service2.Service{Task: task}
			handler := NewTaskHandler(service)

			r := gin.New()
			r.GET("/task/:id", handler.GetTask)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", fmt.Sprintf("/task/%d", testCase.taskId), nil)

			r.ServeHTTP(w, req)

			if testCase.expectedStatusCode != w.Code {
				t.Error("status codes is not equal")
			}
			if testCase.expectedRequestBody != w.Body.String() {
				t.Error("requests is not equal")
			}
		})
	}
}

func TestTaskHandler_GetTasksList(t *testing.T) {
	type mockBehaviour func(s *mock_service.MockTask)

	date, err := time.Parse("2006-01-02", "2023-09-23")
	if err != nil {
		t.Errorf("can't prepare data for testing: %s", err.Error())
	}

	testsTable := []struct {
		name                string
		mockBehaviour       mockBehaviour
		expectedStatusCode  int
		expectedRequestBody string
	}{
		{
			name: "OK",
			mockBehaviour: func(s *mock_service.MockTask) {
				s.EXPECT().GetTasksList().Return([]models.Task{
					{
						Id:          1,
						IsDone:      false,
						Header:      "убраться в комнате",
						Description: "помыть посуду",
						Date:        date,
					},
				}, nil)
			},
			expectedStatusCode:  200,
			expectedRequestBody: `[{"is_done":false,"id":1,"header":"убраться в комнате","description":"помыть посуду","date":"2023-09-23T00:00:00Z"}]`,
		},
		{
			name: "OK NIL",
			mockBehaviour: func(s *mock_service.MockTask) {
				s.EXPECT().GetTasksList().Return([]models.Task{}, nil)
			},
			expectedStatusCode:  200,
			expectedRequestBody: `[]`,
		},
		{
			name: "Server Error",
			mockBehaviour: func(s *mock_service.MockTask) {
				s.EXPECT().GetTasksList().Return([]models.Task{}, fmt.Errorf("some errors"))
			},
			expectedStatusCode:  500,
			expectedRequestBody: `{"message":"server error: some errors"}`,
		},
	}

	for _, testCase := range testsTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			task := mock_service.NewMockTask(c)
			testCase.mockBehaviour(task)

			service := service2.Service{Task: task}
			handler := NewTaskHandler(service)

			r := gin.New()
			r.GET("/task/list", handler.GetTasksList)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", "/task/list", nil)

			r.ServeHTTP(w, req)

			if testCase.expectedStatusCode != w.Code {
				t.Error("status codes is not equal")
			}
			if testCase.expectedRequestBody != w.Body.String() {
				t.Error("requests is not equal")
			}
		})
	}
}

func TestTaskHandler_GetTasksByDate(t *testing.T) {
	type mockBehaviour func(s *mock_service.MockTask, date time.Time, isDone bool)

	date, err := time.Parse("2006-01-02", "2023-09-23")
	if err != nil {
		t.Errorf("can't prepare data for testing: %s", err.Error())
	}

	testsTable := []struct {
		name                string
		Date                time.Time
		IsDone              string
		mockBehaviour       mockBehaviour
		expectedStatusCode  int
		expectedRequestBody string
	}{
		{
			name:   "OK",
			Date:   date,
			IsDone: "false",
			mockBehaviour: func(s *mock_service.MockTask, date time.Time, isDone bool) {
				s.EXPECT().GetTasksByDate(date, isDone).Return([]models.Task{
					{
						Id:          1,
						IsDone:      false,
						Header:      "убраться в комнате",
						Description: "помыть посуду",
						Date:        date,
					},
				}, nil)
			},
			expectedStatusCode:  200,
			expectedRequestBody: `[{"is_done":false,"id":1,"header":"убраться в комнате","description":"помыть посуду","date":"2023-09-23T00:00:00Z"}]`,
		},
		{
			name:                "OK",
			Date:                date,
			IsDone:              "123",
			mockBehaviour:       func(s *mock_service.MockTask, date time.Time, isDone bool) {},
			expectedStatusCode:  400,
			expectedRequestBody: `{"message":"incorrect flag"}`,
		},
	}

	for _, testCase := range testsTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			var boolCount bool
			if testCase.IsDone == "1" {
				boolCount = true
			} else {
				boolCount = false
			}

			task := mock_service.NewMockTask(c)
			testCase.mockBehaviour(task, testCase.Date, boolCount)

			service := service2.Service{Task: task}
			handler := NewTaskHandler(service)

			r := gin.New()
			r.GET("/task/list/date", handler.GetTasksByDate)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", fmt.Sprintf("/task/list/date?date=%s&is_done=%s", date.Format("2006-01-02"), testCase.IsDone), nil)

			r.ServeHTTP(w, req)

			if testCase.expectedStatusCode != w.Code {
				t.Error("status codes is not equal")
			}
			if testCase.expectedRequestBody != w.Body.String() {
				t.Error("requests is not equal")
			}
		})
	}
}

func TestTaskHandler_GetTasksByPage(t *testing.T) {
	type mockBehaviour func(s *mock_service.MockTask, isDone bool, page int)

	date, err := time.Parse("2006-01-02", "2023-09-23")
	if err != nil {
		t.Errorf("can't prepare data for testing: %s", err.Error())
	}

	testsTable := []struct {
		name                string
		inputBody           string
		Page                int
		IsDone              string
		mockBehaviour       mockBehaviour
		expectedStatusCode  int
		expectedRequestBody string
	}{
		{
			name:   "OK",
			Page:   1,
			IsDone: "false",
			mockBehaviour: func(s *mock_service.MockTask, isDone bool, page int) {
				s.EXPECT().GetTasksByPage(isDone, page).Return([]models.Task{
					{
						Id:          1,
						IsDone:      false,
						Header:      "убраться в комнате",
						Description: "помыть посуду",
						Date:        date,
					},
				}, nil)
			},
			expectedStatusCode:  200,
			expectedRequestBody: `[{"is_done":false,"id":1,"header":"убраться в комнате","description":"помыть посуду","date":"2023-09-23T00:00:00Z"}]`,
		},
		{
			name:   "Empty Page",
			Page:   -1,
			IsDone: "false",
			mockBehaviour: func(s *mock_service.MockTask, isDone bool, page int) {
				s.EXPECT().GetTasksByPage(isDone, page).Return([]models.Task{}, fmt.Errorf("server error: pq: OFFSET must not be negative"))
			},
			expectedStatusCode:  500,
			expectedRequestBody: `{"message":"server error: server error: pq: OFFSET must not be negative"}`,
		},
		{
			name:                "Not Valid IsDone",
			Page:                4,
			IsDone:              "123",
			mockBehaviour:       func(s *mock_service.MockTask, isDone bool, page int) {},
			expectedStatusCode:  400,
			expectedRequestBody: `{"message":"incorrect flag"}`,
		},
	}

	for _, testCase := range testsTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			var boolCount bool
			if testCase.IsDone == "1" {
				boolCount = true
			} else {
				boolCount = false
			}

			task := mock_service.NewMockTask(c)
			testCase.mockBehaviour(task, boolCount, testCase.Page)

			service := service2.Service{Task: task}
			handler := NewTaskHandler(service)

			r := gin.New()
			r.GET("/task/list/page", handler.GetTasksByPage)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", fmt.Sprintf("/task/list/page?is_done=%s&page=%d", testCase.IsDone, testCase.Page), nil)

			r.ServeHTTP(w, req)

			if testCase.expectedStatusCode != w.Code {
				t.Error("status codes is not equal")
			}
			if testCase.expectedRequestBody != w.Body.String() {
				t.Error("requests is not equal")
			}
		})
	}
}
