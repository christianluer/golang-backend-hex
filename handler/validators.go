package handler

import "github.com/go-playground/validator"

var Validate = validator.New()

type RegisterUserRequest struct {
	Username string `json:"username" validate:"required,min=3,max=32"`
	Password string `json:"password" validate:"required,min=6"`
}
