package service

import (
	"github.com/bearury/go-rest-api-postgres-todo/entitys"
	"github.com/bearury/go-rest-api-postgres-todo/pakage/repository"
)

type Authorization interface {
	CreateUser(user entitys.User) (string, error)
	GenerateToken(username string, password string) (string, error)
	ParseToken(token string) (string, error)
}

type TodoList interface {
	CreateList(userId string, list entitys.Todo) (string, error)
	GetAllLists(userId string) ([]entitys.Todo, error)
	GetListById(userId, listId string) (entitys.Todo, error)
	UpdateList(userId, listId string, input entitys.UpdateListInput) error
	DeleteList(userId, listId string) error
}

type TodoItem interface {
	CreateItem(userId, listId string, list entitys.TodoItem) (string, error)
	GetAllItems(userId, listId string) ([]entitys.TodoItem, error)
	GetItemById(userId, itemId string) (entitys.TodoItem, error)
	UpdateItem(userId, itemId string, input entitys.UpdateItemInput) error
	DeleteItem(userId, itemId string) error
}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repo.AuthorizationRepository),
		TodoList:      NewTodoListService(repo.TodoListRepository),
		TodoItem:      NewTodoItemService(repo.TodoItemRepository, repo.TodoListRepository),
	}
}
