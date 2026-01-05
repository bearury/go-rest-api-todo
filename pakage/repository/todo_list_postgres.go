package repository

import (
	"fmt"

	"github.com/bearury/go-rest-api-postgres-todo/entitys"
	"github.com/jmoiron/sqlx"
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
