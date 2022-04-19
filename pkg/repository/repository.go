package repository

import "github.com/jmoiron/sqlx"

type Authorization interface {
}

type ToDoList interface {
}

type ToDoItem interface {
}

type Repository struct {
	Authorization
	ToDoItem
	ToDoList
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{}
}
