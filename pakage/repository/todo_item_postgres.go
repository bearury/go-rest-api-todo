package repository

import (
	"fmt"

	"github.com/bearury/go-rest-api-postgres-todo/entitys"
	"github.com/jmoiron/sqlx"
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
		tx.Rollback()
		return "", err
	}

	createListItemQuery := fmt.Sprintf("INSERT INTO %s (list_id, item_id) VALUES ($1, $2)", listItemsTable)
	_, err = tx.Exec(createListItemQuery, listId, id)
	if err != nil {
		tx.Rollback()
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

func (r *TodoItemPostgres) GetItemById(userId, listId string) (entitys.TodoItem, error) {
	var todoList entitys.TodoItem

	query := fmt.Sprintf("SELECT tl.id, tl.title, tl.description FROM %s AS tl INNER JOIN %s AS ul ON tl.id = ul.list_id WHERE ul.user_id = $1 AND ul.list_id = $2", todoListsTable, usersListTable)

	err := r.db.Get(&todoList, query, userId, listId)

	return todoList, err
}
