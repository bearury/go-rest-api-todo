package repository

import (
	"fmt"
	"strings"

	"github.com/bearury/go-rest-api-postgres-todo/entitys"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type TodoListPostgres struct {
	db *sqlx.DB
}

func NewTodoListPostgres(db *sqlx.DB) *TodoListPostgres {
	return &TodoListPostgres{db: db}
}

func (r *TodoListPostgres) Create(userId string, list entitys.Todo) (string, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return "", err
	}

	var id string
	createListQuery := fmt.Sprintf("INSERT INTO %s (title, description) VALUES ($1, $2) RETURNING id", todoListsTable)
	row := tx.QueryRow(createListQuery, list.Title, list.Description)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return "", err
	}

	createUsersListQuery := fmt.Sprintf("INSERT INTO %s (user_id, list_id) VALUES ($1, $2)", usersListTable)
	_, err = tx.Exec(createUsersListQuery, userId, id)
	if err != nil {
		tx.Rollback()
		return "", err
	}

	return id, tx.Commit()
}

func (r *TodoListPostgres) GetAllLists(userId string) ([]entitys.Todo, error) {
	var todoLists []entitys.Todo

	query := fmt.Sprintf("SELECT tl.id, tl.title, tl.description FROM %s AS tl INNER JOIN %s AS ul ON tl.id = ul.list_id WHERE ul.user_id = $1", todoListsTable, usersListTable)

	err := r.db.Select(&todoLists, query, userId)

	return todoLists, err
}

func (r *TodoListPostgres) GetListById(userId, listId string) (entitys.Todo, error) {
	var todoList entitys.Todo

	query := fmt.Sprintf("SELECT tl.id, tl.title, tl.description FROM %s AS tl INNER JOIN %s AS ul ON tl.id = ul.list_id WHERE ul.user_id = $1 AND ul.list_id = $2", todoListsTable, usersListTable)

	err := r.db.Get(&todoList, query, userId, listId)

	return todoList, err
}

func (r *TodoListPostgres) DeleteList(userId, listId string) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	queryDeleteItems := fmt.Sprintf("DELETE FROM %s WHERE id IN (SELECT item_id FROM %s WHERE list_id = $1)", todoItemsTable, listItemsTable)
	_, err = tx.Exec(queryDeleteItems, listId)
	if err != nil {
		_ = tx.Rollback()
		return fmt.Errorf("не удалось удалить элементы списка: %w", err)
	}

	queryDeleteList := fmt.Sprintf("DELETE FROM %s tl USING %s ul WHERE tl.id = ul.list_id AND ul.user_id=$1 AND ul.list_id=$2",
		todoListsTable, usersListTable)
	_, err = tx.Exec(queryDeleteList, userId, listId)
	if err != nil {
		_ = tx.Rollback()
		return fmt.Errorf("не удалось удалить список: %w", err)
	}

	return tx.Commit()
}

func (r *TodoListPostgres) UpdateList(userId, listId string, input entitys.UpdateListInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title=$%d", argId))
		args = append(args, *input.Title)
		argId++
	}

	if input.Description != nil {
		setValues = append(setValues, fmt.Sprintf("description=$%d", argId))
		args = append(args, *input.Description)
		argId++
	}

	// title=$1
	// description=$1
	// title=$1, description=$2
	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE %s tl SET %s FROM %s ul WHERE tl.id = ul.list_id AND ul.list_id=$%d AND ul.user_id=$%d",
		todoListsTable, setQuery, usersListTable, argId, argId+1)
	args = append(args, listId, userId)

	logrus.Debugf("updateQuery: %s", query)
	logrus.Debugf("args: %s", args)

	_, err := r.db.Exec(query, args...)
	return err
}
