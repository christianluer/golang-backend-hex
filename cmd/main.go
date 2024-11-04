package main

import (
	"fmt"

	"log"
	"net/http"

	"github.com/christianluer/golang-backend-hex/config"
	"github.com/christianluer/golang-backend-hex/handler"
	"github.com/christianluer/golang-backend-hex/infrastructure/persistence"
	"github.com/christianluer/golang-backend-hex/router"
	"github.com/christianluer/golang-backend-hex/service"
)

func main() {
	// Connect to the database
	db, err := config.ConnectDB()
	if err != nil {
		log.Fatalf("Could not connect to MySQL: %v\n", err)
	}
	defer db.Close()
	userRepo := persistence.NewMySQLUserRepo(db)
	userService := service.NewUserService(userRepo)
	authService := service.NewAuthService(userRepo)
	userHandler := handler.NewUserHandler(userService)
	authHandler := handler.NewAuthHandler(authService)
	r := router.SetupRouter(userHandler, authHandler)
	port := ":8080"
	fmt.Printf("Server running on %s\n", port)
	log.Fatal(http.ListenAndServe(port, r))
}
