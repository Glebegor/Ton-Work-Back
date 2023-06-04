package service

import (
	"crypto/sha1"
	"fmt"
	"os"
	"time"

	TonWork "github.com/TonWork/back"
	"github.com/TonWork/back/pkg/repository"
	"github.com/golang-jwt/jwt"
)

type AuthService struct {
	repo repository.Authorization
}
type TokenClaims struct {
	jwt.StandardClaims
	UserId       int
	UserUsername string `json:"user_username"`
	UserName     string `json:"user_name"`
	UserSurname  string `json:"user_surname"`
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user TonWork.User) error {
	user.Password_hash = s.PasswordHash(user.Password_hash)
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
func (s *AuthService) GenerateToken(username, password string) (string, error) {
	user, err := s.repo.GetUser(username, s.PasswordHash(password))
	if err != nil {
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &TokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.Id,
		username,
		user.Name,
		user.Surname,
	})
	return token.SignedString([]byte(os.Getenv("Secret_Key")))
}
