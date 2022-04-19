package service

import "github.com/alamansson/todo/pkg/repository"

type Authorization interface {
}

type ToDoList interface {
}

type ToDoItem interface {
}

type Service struct {
	Authorization
	ToDoItem
	ToDoList
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}
