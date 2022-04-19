package repository

import (
	"github.com/alamansson/todo"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user todo.User) (int, error)
	GetUser(username, password string) (todo.User, error)
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
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
