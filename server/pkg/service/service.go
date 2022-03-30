package service

import (
	"github.com/Argento-prg/NoteBook/server/models"
	repository "github.com/Argento-prg/NoteBook/server/pkg/repository"
)

type Authorization interface {
	CreateUser(user models.User) (int, error)
	GenerateToken(login, password string) (string, string, error)
	ParseToken(token string) (int, error)
}

type List interface {
	Create(userId int, list models.List) (int, error)
	GetAll(userId int) ([]models.List, error)
	GetById(userId, listId int) (models.List, error)
	Update(userId, listId int, input models.UpdateListInput) error
	Delete(userId, listId int) error
}

type Todo interface {
	Create(userId, listId int, todo models.Todo) (int, error)
	GetAll(userId, listId int) ([]models.Todo, error)
	GetById(userId, todoId int) (models.Todo, error)
	Update(userId, todoId int, input models.UpdateTodoInput) error
	Delete(userId, todoId int) error
}

type Service struct {
	Authorization
	List
	Todo
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		List:          NewListService(repos.List),
		Todo:          NewTodoService(repos.Todo, repos.List),
	}
}
