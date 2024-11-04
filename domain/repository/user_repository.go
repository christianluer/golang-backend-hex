package repository

import (
	"github.com/christianluer/golang-backend-hex/domain/model"
)

type UserRepository interface {
	GetByUsername(username string) (*model.User, error)
	GetById(id int) (*model.User, error)
	Create(user *model.User) error
	Update(user *model.User) error
	Delete(user *model.User) error
}
