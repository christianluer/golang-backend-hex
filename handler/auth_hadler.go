package handler

import (
	"encoding/json"
	"net/http"

	"github.com/christianluer/golang-backend-hex/service"
)

type AuthHandler struct {
	authService *service.AuthService
}

func NewAuthHandler(authService *service.AuthService) *AuthHandler {
	return &AuthHandler{authService: authService}
}

func (handler *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var creds struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		http.Error(w, "Invalid Body", http.StatusBadRequest)
		return
	}
	token, err := handler.authService.Authenticate(creds.Username, creds.Password)
	if err != nil {
		http.Error(w, "Authentication failed", http.StatusUnauthorized)
		return
	}
	json.NewEncoder(w).Encode(map[string]string{token: token})
}
