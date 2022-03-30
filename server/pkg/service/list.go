package service

import (
	"github.com/Argento-prg/NoteBook/server/models"
	"github.com/Argento-prg/NoteBook/server/pkg/repository"
)

type ListService struct {
	repo repository.List
}

func NewListService(repo repository.List) *ListService {
	return &ListService{repo: repo}
}

func (s *ListService) Create(userId int, list models.List) (int, error) {
	return s.repo.Create(userId, list)
}

func (s *ListService) GetAll(userId int) ([]models.List, error) {
	return s.repo.GetAll(userId)
}

func (s *ListService) GetById(userId, listId int) (models.List, error) {
	return s.repo.GetById(userId, listId)
}

func (s *ListService) Update(userId, listId int, input models.UpdateListInput) error {
	if err := input.Validate(); err != nil {
		return err
	}
	return s.repo.Update(userId, listId, input)
}

func (s *ListService) Delete(userId, listId int) error {
	return s.repo.Delete(userId, listId)
}
