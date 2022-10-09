package service

import (
	goapp "github.com/QANTBEK/Go"
	"github.com/QANTBEK/Go/pkg/repository"
)

type Autorization interface {
	CreateUser(user goapp.User) (int, error)
}

type Todolist interface {
}

type TodoItem interface {
}

type Service struct {
	Autorization
	Todolist
	TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Autorization: NewAuthService(repos.Autorization),
	}
}
