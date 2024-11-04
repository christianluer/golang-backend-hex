package service

import (
	"errors"

	"github.com/christianluer/golang-backend-hex/domain/model"
	"github.com/christianluer/golang-backend-hex/domain/repository"
)

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
	user := &model.User{Username: username, Password: password}
	founded, _ := service.userRepo.GetByUsername(username)
	if founded != nil {
		return nil, errors.New("user found")
	}
	err := service.userRepo.Create(user)
	return user, err
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
