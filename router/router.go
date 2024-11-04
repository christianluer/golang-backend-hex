package router

import (
	"github.com/christianluer/golang-backend-hex/handler"
	"github.com/gorilla/mux"
)

func SetupRouter(authHandler *handler.AuthHandler) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/login", authHandler.Login).Methods("POST")
	return r
}
