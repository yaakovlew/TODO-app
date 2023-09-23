package repository

import (
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"task/pkg/models"
	"time"
)

type TaskRepository struct {
	db *sqlx.DB
}

func NewTaskRepository(db *sqlx.DB) *TaskRepository {
	return &TaskRepository{
		db: db,
	}
}

func (r *TaskRepository) GetTask(id int) (models.Task, error) {
	var task models.Task
	query := fmt.Sprintf("SELECT * FROM %s WHERE id = $1 ",
		taskTable)
	if err := r.db.Get(&task, query, id); err != nil {
		err := errors.New("there isn't this task")
		return models.Task{}, err
	}

	return task, nil
}

func (r *TaskRepository) CreateTask(task models.TaskInput) (int, error) {
	query := fmt.Sprintf("INSERT INTO %s (header, description, date, is_done) VALUES ($1, $2, $3, $4) RETURNING id",
		taskTable)
	var id int
	if err := r.db.QueryRow(query, task.Header, task.Description, task.Date, task.IsDone).Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *TaskRepository) DeleteTask(id int) (int, error) {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1",
		taskTable)
	if _, err := r.db.Exec(query, id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *TaskRepository) UpdateTask(id int, task models.TaskInput) (models.TaskInput, error) {
	query := fmt.Sprintf("UPDATE %s SET header = $1, description = $2, date = $3, is_done = $4 WHERE id = $5",
		taskTable)
	if _, err := r.db.Exec(query, task.Header, task.Description, task.Date, task.IsDone, id); err != nil {
		return models.TaskInput{}, err
	}

	return task, nil
}

func (r *TaskRepository) GetTasksList() ([]models.Task, error) {
	var tasks []models.Task
	query := fmt.Sprintf("SELECT * FROM %s ORDER BY date, id",
		taskTable)
	if err := r.db.Select(&tasks, query); err != nil {
		return nil, err
	}

	return tasks, nil
}

func (r *TaskRepository) GetTasksByPage(isDone bool, limit, page int) ([]models.Task, error) {
	offset := limit * (page - 1)
	var tasks []models.Task
	query := fmt.Sprintf("SELECT * FROM %s WHERE is_done = $1 ORDER BY id LIMIT $2 OFFSET $3",
		taskTable)
	if err := r.db.Select(&tasks, query, isDone, limit, offset); err != nil {
		return nil, err
	}

	return tasks, nil
}

func (r *TaskRepository) GetTasksByDate(date time.Time, isDone bool) ([]models.Task, error) {
	var tasks []models.Task
	query := fmt.Sprintf("SELECT * FROM %s WHERE is_done = $1 AND date = $2 ORDER BY id",
		taskTable)
	if err := r.db.Select(&tasks, query, isDone, date); err != nil {
		return nil, err
	}

	return tasks, nil
}
