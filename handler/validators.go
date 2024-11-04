package handler

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
	validate.RegisterStructValidation(UpdateUserRequestValidation, UpdateUserRequest{})
}

type RegisterUserRequest struct {
	Username string `json:"username" validate:"required,min=3,max=32"`
	Password string `json:"password" validate:"required,min=4"`
}

type UpdateUserRequest struct {
	Username string `json:"username" validate:"omitempty,min=3,max=32"`
	Password string `json:"password" validate:"omitempty,min=4"`
}

func jsonErrorResponse(w http.ResponseWriter, message string, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	response := map[string]string{"error": message}
	json.NewEncoder(w).Encode(response)
}

func UpdateUserRequestValidation(sl validator.StructLevel) {
	req := sl.Current().Interface().(UpdateUserRequest)
	if req.Username == "" && req.Password == "" {
		sl.ReportError(req.Username, "Username", "Username", "atleastone", "Username or Password")
	}
}

func structValidate(v interface{}) error {
	if err := validate.Struct(v); err != nil {
		return err
	}
	return nil
}
