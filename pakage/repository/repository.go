package repository

import (
	"github.com/bearury/go-rest-api-postgres-todo/entitys"
	"github.com/jmoiron/sqlx"
)

type AuthorizationRepository interface {
	CreateUser(user entitys.User) (int, error)
}

type TodoListRepository interface{}

type TodoItemRepository interface{}

type Repository struct {
	AuthorizationRepository
	TodoListRepository
	TodoItemRepository
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		AuthorizationRepository: NewAuthPostgres(db),
	}
}
