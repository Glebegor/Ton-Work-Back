package service

import (
	"strconv"

	"github.com/Glebegor/Ton-Work-Back/pkg/repository"
	TonWork "github.com/Glebegor/Ton-Work-Back/structint"
)

type PostService struct {
	repo repository.Post
}

func NewPostService(repo repository.Post) *PostService {
	return &PostService{repo: repo}
}

func (s *PostService) GetAll() ([]TonWork.Post, error) {
	return s.repo.GetAll()
}
func (s *PostService) Create(UserId int, data TonWork.Post) error {
	return s.repo.Create(UserId, data)
}
func (s *PostService) GetById(id int) (TonWork.Post, error) {
	return s.repo.GetById(id)
}
func (s *PostService) Update(id int, input TonWork.PostUpdate) error {
	idStr := strconv.Itoa(id)
	err := s.repo.Update(idStr, input)
	return err
}
func (s *PostService) Delete(id int) error {
	idStr := strconv.Itoa(id)
	err := s.repo.Delete(idStr)
	return err
}
