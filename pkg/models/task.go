package models

import "time"

type TaskInput struct {
	IsDone      *bool     `json:"is_done" binding:"required"`
	Header      string    `json:"header" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Date        time.Time `json:"date" binding:"required"`
}

type Task struct {
	IsDone      bool      `json:"is_done" db:"is_done"`
	Id          int       `json:"id" db:"id"`
	Header      string    `json:"header" db:"header"`
	Description string    `json:"description" db:"description"`
	Date        time.Time `json:"date" db:"date"`
}
