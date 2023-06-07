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
	data, err := s.repo.GetAll()
	return data, err
}
func (s *PostService) Create(UserId int, data TonWork.Post) error {
	err := s.repo.Create(UserId, data)
	return err
}
