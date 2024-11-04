package handler

import (
	"encoding/json"
	"fmt"
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
	if err := structValidate(req); err != nil {
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

func (handler *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	id, errParsing := getIdFromParams(mux.Vars(r))
	var req UpdateUserRequest
	if errParsing != nil {
		jsonErrorResponse(w, "Error: "+errParsing.Error(), http.StatusInternalServerError)
		return
	}
	_, err := handler.userService.GetUser(id)
	if err != nil {
		jsonErrorResponse(w, "Error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		jsonErrorResponse(w, "Error: "+err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Println(req)
	if err := structValidate(req); err != nil {
		jsonErrorResponse(w, "Error: "+err.Error(), http.StatusInternalServerError)
		return
	}
	err = handler.userService.UpdateUser(id, req.Username, req.Password)
	if err != nil {
		jsonErrorResponse(w, "Error updating user: "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	updateUser := map[string]string{"message": "user updated successfully"}
	json.NewEncoder(w).Encode(updateUser)
}

func getIdFromParams(params map[string]string) (int, error) {
	idStr := params["id"]
	id, errParsing := strconv.Atoi(idStr)
	if errParsing != nil {
		return 0, errParsing
	}
	return id, nil
}

func (handler *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	id, errParsing := getIdFromParams(mux.Vars(r))
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

func (handler *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
}
