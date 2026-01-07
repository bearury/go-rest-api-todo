package service

import (
	"errors"

	"github.com/bearury/go-rest-api-postgres-todo/entitys"
	"github.com/bearury/go-rest-api-postgres-todo/pakage/repository"
)

type TodoListService struct {
	repo     repository.TodoListRepository
	repoItem repository.TodoItemRepository
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

func (service *TodoListService) UpdateList(userId, listId string, input entitys.UpdateListInput) error {
	_, err := service.repo.GetListById(userId, listId)
	if err != nil {
		return errors.New("List already exists")
	}
	return service.repo.UpdateList(userId, listId, input)
}

func (service *TodoListService) DeleteList(userId, listId string) error {
	_, err := service.repo.GetListById(userId, listId)
	if err != nil {
		return errors.New("List already exists")
	}
	return service.repo.DeleteList(userId, listId)
}
