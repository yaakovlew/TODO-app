package repository

import "github.com/jmoiron/sqlx"

type Task interface {
}

type Repo struct {
	Task
}

func NewRepository(db *sqlx.DB) *Repo {
	return &Repo{
		Task: NewTaskRepository(db),
	}
}
