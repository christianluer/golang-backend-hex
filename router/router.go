package router

import (
	"github.com/christianluer/golang-backend-hex/handler"
	"github.com/gorilla/mux"
)

func SetupRouter(userHandler *handler.UserHandler, authHandler *handler.AuthHandler) *mux.Router {
	r := mux.NewRouter()

	// User CRUD routes
	r.HandleFunc("/users", userHandler.Register).Methods("POST")
	r.HandleFunc("/users", userHandler.GetUser).Methods("GET")
	r.HandleFunc("/users", userHandler.UpdateUser).Methods("PUT")
	r.HandleFunc("/users", userHandler.DeleteUser).Methods("DELETE")
	r.HandleFunc("/login", authHandler.Login).Methods("POST")
	return r
}
