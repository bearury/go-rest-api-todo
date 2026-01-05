package service

import (
	"github.com/bearury/go-rest-api-postgres-todo/entitys"
	"github.com/bearury/go-rest-api-postgres-todo/pakage/repository"
)

type TodoListService struct {
	repo repository.TodoListRepository
}

func NewTodoListService(repo repository.TodoListRepository) *TodoListService {
	return &TodoListService{repo: repo}
}

func (service *TodoListService) CreateList(userId string, list entitys.Todo) (string, error) {
	return service.repo.Create(userId, list)
}

func (service *TodoListService) GetAllLists(userId string) ([]entitys.Todo, error) {
	return service.repo.GetAllLists(userId)
}

func (service *TodoListService) GetListById(userId, listId string) (entitys.Todo, error) {
	return service.repo.GetListById(userId, listId)
}
