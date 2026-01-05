package repository

import (
	"github.com/bearury/go-rest-api-postgres-todo/entitys"
	"github.com/jmoiron/sqlx"
)

type AuthorizationRepository interface {
	CreateUser(user entitys.User) (string, error)
	GetUser(username, password string) (entitys.User, error)
}

type TodoListRepository interface {
	Create(userId string, list entitys.Todo) (string, error)
	GetAllLists(userId string) ([]entitys.Todo, error)
	GetListById(userId, listId string) (entitys.Todo, error)
}

type TodoItemRepository interface {
	Create(listId string, item entitys.TodoItem) (string, error)
	GetAllItems(userId, listId string) ([]entitys.TodoItem, error)
	GetItemById(listId, itemId string) (entitys.TodoItem, error)
}

type Repository struct {
	AuthorizationRepository
	TodoListRepository
	TodoItemRepository
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		AuthorizationRepository: NewAuthPostgres(db),
		TodoListRepository:      NewTodoListPostgres(db),
		TodoItemRepository:      NewTodoItemPostgres(db),
	}
}
