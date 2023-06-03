package service

import (
	TonWork "github.com/TonWork/back"
	"github.com/TonWork/back/pkg/repository"
)

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user TonWork.User) error {
	if err := s.repo.CreateUser(user); err != nil {
		return err
	}

	return nil
}
