package service

import (
	"github.com/bearury/go-rest-api-postgres-todo/entitys"
	"github.com/bearury/go-rest-api-postgres-todo/pakage/repository"
)

type AuthorizationService interface {
	CreateUser(user entitys.User) (int, error)
}

type TodoListService interface{}

type TodoItemService interface{}

type Service struct {
	AuthorizationService
	TodoListService
	TodoItemService
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		AuthorizationService: NewAuthService(repo.AuthorizationRepository),
	}
}
