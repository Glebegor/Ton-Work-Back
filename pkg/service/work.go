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
	data, err := s.repo.GetAll()
	return data, err
}
func (s *WorkService) Create(UserId int, data TonWork.Work) error {
	err := s.repo.Create(UserId, data)
	return err
}
