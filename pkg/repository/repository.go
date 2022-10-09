package repository

import (
	goapp "github.com/QANTBEK/Go"
	"github.com/jmoiron/sqlx"
)

type Autorization interface {
	CreateUser(user goapp.User) (int, error)
}

type Todolist interface {
}

type TodoItem interface {
}

type Repository struct {
	Autorization
	Todolist
	TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Autorization: NewAuthPostgres(db),
	}
}
