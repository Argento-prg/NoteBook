package repository

import (
	"github.com/Argento-prg/NoteBook/server/models"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user models.User) (int, error)
	GetUser(login, password string) (models.User, error)
}

type List interface {
	Create(userId int, list models.List) (int, error)
	GetAll(userId int) ([]models.List, error)
	GetById(userId, listId int) (models.List, error)
	Update(userId, listId int, input models.UpdateListInput) error
	Delete(userId, listId int) error
}

type Todo interface {
	Create(listId int, todo models.Todo) (int, error)
	GetAll(userId, listId int) ([]models.Todo, error)
	GetById(userId, todoId int) (models.Todo, error)
	Update(userId, todoId int, input models.UpdateTodoInput) error
	Delete(userId, todoId int) error
}

type Repository struct {
	Authorization
	List
	Todo
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		List:          NewListPostgres(db),
		Todo:          NewTodoPostgres(db),
	}
}
