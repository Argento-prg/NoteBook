package repository

import (
	"fmt"
	"strings"

	"github.com/Argento-prg/NoteBook/server/models"
	"github.com/jmoiron/sqlx"
)

type TodoPostgres struct {
	db *sqlx.DB
}

func NewTodoPostgres(db *sqlx.DB) *TodoPostgres {
	return &TodoPostgres{db: db}
}

func (r *TodoPostgres) Create(listId int, todo models.Todo) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}
	var todoId int
	createItemQuery := fmt.Sprintf(
		"INSERT INTO %s (title, description, importance, done) VALUES ($1, $2, $3, $4) RETURNING id",
		todosTable,
	)
	row := tx.QueryRow(createItemQuery, todo.Title, todo.Description, todo.Importance, todo.Done)
	err = row.Scan(&todoId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	createListItemsQuery := fmt.Sprintf(
		"INSERT INTO %s (listId, todoId) VALUES ($1, $2)",
		listsTodosTable,
	)
	_, err = tx.Exec(createListItemsQuery, listId, todoId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	return todoId, tx.Commit()
}

func (r *TodoPostgres) GetAll(userId, listId int) ([]models.Todo, error) {
	var todos []models.Todo
	query := fmt.Sprintf(
		`SELECT t.id, t.title, t.importance, t.done FROM %s t 
		INNER JOIN %s lt ON lt.todoId = t.id 
		INNER JOIN %s ul ON ul.listId = lt.listId 
		WHERE lt.listId = $1 
		AND ul.userId = $2 ORDER BY t.id`,
		todosTable,
		listsTodosTable,
		usersListsTable,
	)
	if err := r.db.Select(&todos, query, listId, userId); err != nil {
		return nil, err
	}
	return todos, nil
}

func (r *TodoPostgres) GetById(userId, todoId int) (models.Todo, error) {
	var todo models.Todo
	query := fmt.Sprintf(
		`SELECT t.id, t.title, t.description, t.importance, t.done FROM %s t 
		INNER JOIN %s lt ON lt.todoId = t.id 
		INNER JOIN %s ul ON ul.listId = lt.listId 
		WHERE t.id = $1 
		AND ul.userId = $2`,
		todosTable,
		listsTodosTable,
		usersListsTable,
	)
	if err := r.db.Get(&todo, query, todoId, userId); err != nil {
		return todo, err
	}
	return todo, nil
}

func (r *TodoPostgres) Update(userId, todoId int, input models.UpdateTodoInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1
	if input.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title = $%d", argId))
		args = append(args, *input.Title)
		argId++
	}
	if input.Description != nil {
		setValues = append(setValues, fmt.Sprintf("description = $%d", argId))
		args = append(args, *input.Description)
		argId++
	}
	if input.Importance != nil {
		setValues = append(setValues, fmt.Sprintf("importance = $%d", argId))
		args = append(args, *input.Importance)
		argId++
	}
	if input.Done != nil {
		setValues = append(setValues, fmt.Sprintf("done = $%d", argId))
		args = append(args, *input.Done)
		argId++
	}
	setQuery := strings.Join(setValues, ", ")
	query := fmt.Sprintf(
		`UPDATE %s t SET %s FROM %s lt, %s ul 
		WHERE t.id = lt.todoId 
		AND lt.listId = ul.listId 
		AND ul.userId = $%d 
		AND t.id = $%d`,
		todosTable,
		setQuery,
		listsTodosTable,
		usersListsTable,
		argId,
		argId+1,
	)
	args = append(args, userId, todoId)
	_, err := r.db.Exec(query, args...)
	return err
}

func (r *TodoPostgres) Delete(userId, todoId int) error {
	query := fmt.Sprintf(
		`DELETE FROM %s t USING %s lt, %s ul 
		WHERE t.id = lt.todoId 
		AND lt.listId = ul.listId 
		AND ul.userId = $1 
		AND t.id = $2`,
		todosTable,
		listsTodosTable,
		usersListsTable,
	)
	_, err := r.db.Exec(query, userId, todoId)
	return err
}
