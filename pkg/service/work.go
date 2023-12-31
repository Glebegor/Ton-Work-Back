package service

import (
	"strconv"

	"github.com/Glebegor/Ton-Work-Back/pkg/repository"
	TonWork "github.com/Glebegor/Ton-Work-Back/structint"
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
func (s *WorkService) Update(id int, input TonWork.WorkUpdate) error {
	idStr := strconv.Itoa(id)
	err := s.repo.Update(idStr, input)
	return err
}
func (s *WorkService) Delete(id int) error {
	idStr := strconv.Itoa(id)
	err := s.repo.Delete(idStr)
	return err
}
