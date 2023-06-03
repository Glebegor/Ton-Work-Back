package service

import (
	"crypto/sha1"
	"fmt"
	"os"

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
	user.Person.Password_hash = s.PasswordHash(user.Person.Password_hash)
	if err := s.repo.CreateUser(user); err != nil {
		return err
	}
	return nil
}
func (s *AuthService) PasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(os.Getenv("Secret_Key"))))
}
