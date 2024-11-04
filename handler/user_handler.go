package handler

import (
	"encoding/json"
	"net/http"

	"github.com/christianluer/golang-backend-hex/domain/model"
	"github.com/christianluer/golang-backend-hex/service"
)

type UserHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func (handler *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	var user model.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	createUser, err := handler.userService.RegisterUser(user.Username, user.Password)
	if err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(createUser)
}

func (handler *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
}

func (handler *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
}

func (handler *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
}
