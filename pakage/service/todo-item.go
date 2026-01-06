package service

import (
	"errors"

	"github.com/bearury/go-rest-api-postgres-todo/entitys"
	"github.com/bearury/go-rest-api-postgres-todo/pakage/repository"
)

type TodoItemService struct {
	repo     repository.TodoItemRepository
	repoList repository.TodoListRepository
}

func NewTodoItemService(repo repository.TodoItemRepository, repoList repository.TodoListRepository) *TodoItemService {
	return &TodoItemService{repo: repo, repoList: repoList}
}

func (service *TodoItemService) CreateItem(userId, listId string, item entitys.TodoItem) (string, error) {
	_, err := service.repoList.GetListById(userId, listId)
	if err != nil {
		return "", errors.New("item already exists")
	}

	return service.repo.Create(listId, item)
}

func (service *TodoItemService) GetAllItems(userId, listId string) ([]entitys.TodoItem, error) {
	_, err := service.repoList.GetListById(userId, listId)
	if err != nil {
		return nil, errors.New("item already exists")
	}
	return service.repo.GetAllItems(userId, listId)
}

func (service *TodoItemService) GetItemById(listId, itemId string) (entitys.TodoItem, error) {
	return service.repo.GetItemById(listId, itemId)
}

func (s *TodoItemService) UpdateItem(userId, itemId string, input entitys.UpdateItemInput) error {
	return s.repo.UpdateItem(userId, itemId, input)
}
