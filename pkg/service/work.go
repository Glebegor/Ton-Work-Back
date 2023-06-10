package service

import (
	TonWork "github.com/TonWork/back"
	"github.com/TonWork/back/pkg/repository"
)

type WorkService struct {
	repo repository.Work
}

func NewWorkService(repo repository.Work) *WorkService {
	return &WorkService{repo: repo}
}

func (s *WorkService) GetAll() ([]TonWork.Work, error) {
	return s.repo.GetAll()
}
func (s *WorkService) Create(UserId int, data TonWork.Work) error {
	return s.repo.Create(UserId, data)
}
func (s *WorkService) GetById(id int) (TonWork.Work, error) {
	return s.repo.GetById(id)
}
