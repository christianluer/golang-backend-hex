package service

import (
	"errors"

	"github.com/christianluer/golang-backend-hex/domain"
	"github.com/christianluer/golang-backend-hex/infrastructure/security"
)

type AuthService struct {
	repo domain.UserRepository
}

func NewAuthService(repo domain.UserRepository) *AuthService {
	return &AuthService{repo: repo}
}

func (authService *AuthService) Authenticate(username, password string) (string, error) {
	user, error := authService.repo.GetByUsername(username)
	if error != nil || user.Password != password {
		return "", errors.New("invalid_credentials")
	}
	return security.GenerateToken(username)
}
