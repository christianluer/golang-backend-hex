package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

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
	var req RegisterUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		jsonErrorResponse(w, "Error: "+err.Error(), http.StatusBadRequest)
		return
	}
	if err := validate.Struct(req); err != nil {
		jsonErrorResponse(w, "Validation failed: "+err.Error(), http.StatusBadRequest)
		return
	}
	createUser, err := handler.userService.RegisterUser(req.Username, req.Password)
	if err != nil {
		jsonErrorResponse(w, "Error: "+err.Error(), http.StatusInternalServerError)
		return
	}
	response := map[string]string{"username": createUser.Username}
	w.Header().Set("Content-Type", "application/json")
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
