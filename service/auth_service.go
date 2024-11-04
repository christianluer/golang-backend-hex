package service

import (
	"errors"

	"github.com/christianluer/golang-backend-hex/domain/repository"
	"github.com/christianluer/golang-backend-hex/infrastructure/security"
)

type AuthService interface {
	Authenticate(username, password string) (string, error)
}

type authService struct {
	userRepo repository.UserRepository
}

func NewAuthService(userRepo repository.UserRepository) AuthService {
	return &authService{userRepo: userRepo}
}

func (service *authService) Authenticate(username, password string) (string, error) {
	user, err := service.userRepo.GetByUsername(username)
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	if user.Password != password {
		return "", errors.New("invalid credentials")
	}
	token, err := security.GenerateToken(username)
	if err != nil {
		return "", err
	}

	return token, nil
}
