package main

import (
	"log"
	"zip-app/internal/api/handlers"
	"zip-app/internal/api/routers"
	"zip-app/internal/database"
	"zip-app/internal/repository"
	"zip-app/internal/service"
)

func main() {
	db := database.NewDatabase()
	repo := repository.NewRepository(db)
	service := service.NewService(repo)
	zipHandler := handlers.NewHandler(service)
	router := routers.SetupRouter(zipHandler)

	if err := router.Run("localhost:8080"); err != nil {
		log.Panicf("error start server: %v", err)
	}
}