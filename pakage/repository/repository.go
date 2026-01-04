package repository

import (
	"github.com/bearury/go-rest-api-postgres-todo/entitys"
	"github.com/jmoiron/sqlx"
)

type AuthorizationRepository interface {
	CreateUser(user entitys.User) (int, error)
	GetUser(username, password string) (entitys.User, error)
}

type TodoListRepository interface {
	Create(userId int, list entitys.Todo) (int, error)
	GetAllLists(userId int) ([]entitys.Todo, error)
	GetListById(userId, listId int) (entitys.Todo, error)
}

type TodoItemRepository interface{}

type Repository struct {
	AuthorizationRepository
	TodoListRepository
	TodoItemRepository
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		AuthorizationRepository: NewAuthPostgres(db),
		TodoListRepository:      NewTodoListPostgres(db),
	}
}
