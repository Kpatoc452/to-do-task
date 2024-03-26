package service

import "github.com/Kpatoc452/to-do-task/models"

type Service struct {
	Task
	Auth
}

type Task interface {
	CreateTask(models.Task) (int, error)
	UpdateTask(models.Task) error
	DeleteTask(int) error
	GetTask(int) models.Task
}

type Auth interface {
	CreateUser(models.User) (int, error)
}
