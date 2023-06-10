package service

import (
	TonWork "github.com/TonWork/back"
	"github.com/TonWork/back/pkg/repository"
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
