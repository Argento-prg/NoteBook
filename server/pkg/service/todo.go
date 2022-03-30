package service

import (
	"github.com/Argento-prg/NoteBook/server/models"
	"github.com/Argento-prg/NoteBook/server/pkg/repository"
)

type TodoService struct {
	repo     repository.Todo
	listRepo repository.List
}

func NewTodoService(repo repository.Todo, listRepo repository.List) *TodoService {
	return &TodoService{repo: repo, listRepo: listRepo}
}

func (s *TodoService) Create(userId, listId int, todo models.Todo) (int, error) {
	_, err := s.listRepo.GetById(userId, listId)
	if err != nil {
		//list does not exists or does not belongs to user
		return 0, err
	}
	return s.repo.Create(listId, todo)
}

func (s *TodoService) GetAll(userId, listId int) ([]models.Todo, error) {
	return s.repo.GetAll(userId, listId)
}

func (s *TodoService) GetById(userId, todoId int) (models.Todo, error) {
	return s.repo.GetById(userId, todoId)
}

func (s *TodoService) Update(userId, todoId int, input models.UpdateTodoInput) error {
	return s.repo.Update(userId, todoId, input)
}

func (s *TodoService) Delete(userId, todoId int) error {
	return s.repo.Delete(userId, todoId)
}
