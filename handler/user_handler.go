package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/christianluer/golang-backend-hex/domain/model"
	"github.com/christianluer/golang-backend-hex/service"
	"github.com/gorilla/mux"
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
	w.Header().Set("Content-Type", "application/json")
	response := map[string]string{"username": createUser.Username}
	json.NewEncoder(w).Encode(response)
}

func (handler *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, errParsing := strconv.Atoi(idStr)
	if errParsing != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}
	user, err := handler.userService.GetUser(id)
	if err != nil {
		http.Error(w, "Failed to get user", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func (handler *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
}

func (handler *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
}
