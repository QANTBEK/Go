package service

import (
	"crypto/sha1"
	"fmt"
	goapp "github.com/QANTBEK/Go"
	"github.com/QANTBEK/Go/pkg/repository"
)

const salt = "csdubcuEWFIOw26"

type AuthService struct {
	repo repository.Autorization
}

func NewAuthService(repo repository.Autorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user goapp.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
