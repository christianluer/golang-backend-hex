package service

import (
	"errors"
	"fmt"

	"github.com/christianluer/golang-backend-hex/domain/model"
	"github.com/christianluer/golang-backend-hex/domain/repository"
)

var ErrUserAlreadyExists = errors.New("user already exists")

type UserService interface {
	RegisterUser(username, password string) (*model.User, error)
	GetUser(id int) (*model.User, error)
	UpdateUser(id int, username, password string) error
	DeleteUser(id int) error
}

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{userRepo: repo}
}

func (service *userService) RegisterUser(username, password string) (*model.User, error) {
	founded, err := service.userRepo.GetByUsername(username)
	if err != nil {
		return nil, fmt.Errorf("error checking user existence: %w", err)
	}
	if founded != nil {
		return nil, ErrUserAlreadyExists
	}
	user := &model.User{Username: username, Password: password}
	err = service.userRepo.Create(user)
	if err != nil {
		return nil, fmt.Errorf("error creating user: %w", err)
	}
	return user, nil
}

func (service *userService) GetUser(id int) (*model.User, error) {
	user, err := service.userRepo.GetById(id)
	if err != nil {
		// error handler
		return nil, err
	}
	return user, nil
}

func (service *userService) UpdateUser(id int, username, password string) error {
	_, err := service.GetUser(id)
	if err != nil {
		return err
	}
	userInstance := &model.User{ID: id, Username: username, Password: password}
	return service.userRepo.Update(userInstance)
}

func (service *userService) DeleteUser(id int) error {
	user, err := service.GetUser(id)
	if err != nil {
		return err
	}
	return service.userRepo.Delete(user)
}
