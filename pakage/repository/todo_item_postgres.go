package repository

import (
	"fmt"
	"strings"

	"github.com/bearury/go-rest-api-postgres-todo/entitys"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type TodoItemPostgres struct {
	db *sqlx.DB
}

func NewTodoItemPostgres(db *sqlx.DB) *TodoItemPostgres {
	return &TodoItemPostgres{db: db}
}

func (r *TodoItemPostgres) Create(listId string, item entitys.TodoItem) (string, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return "", err
	}

	var id string
	createItemQuery := fmt.Sprintf("INSERT INTO %s (title, description) VALUES ($1, $2) RETURNING id", todoItemsTable)
	row := tx.QueryRow(createItemQuery, item.Title, item.Description)
	if err := row.Scan(&id); err != nil {
		err := tx.Rollback()
		if err != nil {
			return "", err
		}
		return "", err
	}

	createListItemQuery := fmt.Sprintf("INSERT INTO %s (list_id, item_id) VALUES ($1, $2)", listItemsTable)
	_, err = tx.Exec(createListItemQuery, listId, id)
	if err != nil {
		err := tx.Rollback()
		if err != nil {
			return "", err
		}
		return "", err
	}

	return id, tx.Commit()
}

func (r *TodoItemPostgres) GetAllItems(userId, listId string) ([]entitys.TodoItem, error) {
	var todoItems []entitys.TodoItem

	query := fmt.Sprintf("SELECT ti.id, ti.title, ti.description, ti.complete FROM %s AS ti INNER JOIN %s AS li ON ti.id = li.item_id INNER JOIN %s AS ul ON ul.list_id = li.list_id WHERE ul.user_id = $1 AND li.list_id = $2", todoItemsTable, listItemsTable, usersListTable)

	err := r.db.Select(&todoItems, query, userId, listId)

	return todoItems, err
}

func (r *TodoItemPostgres) GetItemById(userId, itemId string) (entitys.TodoItem, error) {
	var todoList entitys.TodoItem

	query := fmt.Sprintf(`SELECT ti.id, ti.title, ti.description, ti.complete FROM %s ti INNER JOIN %s li on li.item_id = ti.id
									INNER JOIN %s ul on ul.list_id = li.list_id WHERE ti.id = $1 AND ul.user_id = $2`,
		todoItemsTable, listItemsTable, usersListTable)

	if err := r.db.Get(&todoList, query, itemId, userId); err != nil {
		return todoList, err
	}

	return todoList, nil
}

func (r *TodoItemPostgres) UpdateItem(userId, itemId string, input entitys.UpdateItemInput) error {
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

	if input.Complete != nil {
		setValues = append(setValues, fmt.Sprintf("complete=$%d", argId))
		args = append(args, *input.Complete)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf(`UPDATE %s ti SET %s FROM %s li, %s ul
									WHERE ti.id = li.item_id AND li.list_id = ul.list_id AND ul.user_id = $%d AND ti.id = $%d`,
		todoItemsTable, setQuery, listItemsTable, usersListTable, argId, argId+1)
	args = append(args, userId, itemId)

	logrus.Info("Аргуцменты", args)

	_, err := r.db.Exec(query, args...)
	return err
}

func (r *TodoItemPostgres) DeleteItem(userId, itemId string) error {
	query := fmt.Sprintf(`DELETE FROM %s ti USING %s li, %s ul 
									WHERE ti.id = li.item_id AND li.list_id = ul.list_id AND ul.user_id = $1 AND ti.id = $2`,
		todoItemsTable, listItemsTable, usersListTable)
	_, err := r.db.Exec(query, userId, itemId)
	return err
}
