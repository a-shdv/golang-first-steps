package main

import (
	golang_first_steps "go-first-steps"
	"go-first-steps/pkg/handler"
	"go-first-steps/pkg/repository"
	"go-first-steps/pkg/service"
	"log"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(golang_first_steps.Server)
	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
