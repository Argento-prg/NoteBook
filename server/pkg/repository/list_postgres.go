package repository

import (
	"fmt"
	"strings"

	"github.com/Argento-prg/NoteBook/server/models"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type ListPostgres struct {
	db *sqlx.DB
}

func NewListPostgres(db *sqlx.DB) *ListPostgres {
	return &ListPostgres{db: db}
}

func (r *ListPostgres) Create(userId int, list models.List) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}
	var id int
	createListQuery := fmt.Sprintf(
		"INSERT INTO %s (title, description) VALUES ($1, $2) RETURNING id",
		listsTable,
	)
	row := tx.QueryRow(createListQuery, list.Title, list.Description)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}
	createUsersListsQuery := fmt.Sprintf(
		"INSERT INTO %s (userId, listId) VALUES($1, $2)",
		usersListsTable,
	)
	_, err = tx.Exec(createUsersListsQuery, userId, id)
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	return id, tx.Commit()
}

func (r *ListPostgres) GetAll(userId int) ([]models.List, error) {
	var lists []models.List
	query := fmt.Sprintf(
		`SELECT l.id, l.title FROM %s l 
		INNER JOIN %s ul ON l.id = ul.listId 
		WHERE ul.userId = $1 ORDER BY l.id`,
		listsTable,
		usersListsTable,
	)
	err := r.db.Select(&lists, query, userId)
	return lists, err
}

func (r *ListPostgres) GetById(userId, listId int) (models.List, error) {
	var list models.List
	query := fmt.Sprintf(
		`SELECT l.id, l.title, l.description FROM %s l
		INNER JOIN %s ul 
		ON l.id = ul.listId 
		WHERE ul.userId = $1 AND ul.listId = $2`,
		listsTable,
		usersListsTable,
	)
	err := r.db.Get(&list, query, userId, listId)
	return list, err
}

func (r *ListPostgres) Update(userId, listId int, input models.UpdateListInput) error {
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
	setQuery := strings.Join(setValues, ", ")
	query := fmt.Sprintf(
		`UPDATE %s l SET %s FROM %s ul 
		WHERE l.id = ul.listId 
		AND ul.listId = $%d 
		AND ul.userId = $%d`,
		listsTable,
		setQuery,
		usersListsTable,
		argId,
		argId+1,
	)
	args = append(args, listId, userId)
	logrus.Debugf("updateQuery: %s", query)
	logrus.Debugf("args: %s", args)
	_, err := r.db.Exec(query, args...)
	return err
}

func (r *ListPostgres) Delete(userId, listId int) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	//удаляем все задачи из списка
	queryTodos := fmt.Sprintf(
		`DELETE FROM %s t USING %s lt 
		WHERE t.id = lt.todoId 
		AND lt.listId = $1`,
		todosTable,
		listsTodosTable,
	)
	_, err = tx.Exec(queryTodos, listId)
	if err != nil {
		tx.Rollback()
		return err
	}
	queryLists := fmt.Sprintf(
		`DELETE FROM %s l USING %s ul 
		WHERE l.id = ul.listId 
		AND ul.userId = $1 
		AND ul.listId = $2`,
		listsTable,
		usersListsTable,
	)
	_, err = tx.Exec(queryLists, userId, listId)
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit()
}
