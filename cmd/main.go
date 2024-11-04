package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/christianluer/golang-backend-hex/config"
	"github.com/christianluer/golang-backend-hex/domain"
	"github.com/christianluer/golang-backend-hex/handler"
	"github.com/christianluer/golang-backend-hex/infrastructure/repository"
	"github.com/christianluer/golang-backend-hex/router"
	"github.com/christianluer/golang-backend-hex/service"
)

func main() {
	db, err := config.ConnectDB()
	if err != nil {
		log.Fatalf("Could not connect to MySQL: %v\n", err)
	}
	defer db.Close()
	userRepo := repository.NewMySQLUserRepo(db)
	_ = userRepo.Save(&domain.User{Username: "testuser", Password: "password"}) // Add a test user
	authService := service.NewAuthService(userRepo)
	authHandler := handler.NewAuthHandler(authService)
	r := router.SetupRouter(authHandler)
	port := ":8080"
	fmt.Printf("Server running on %s\n", port)
	log.Fatal(http.ListenAndServe(port, r))
}
