package main

import (
	golang_first_steps "go-first-steps"
	"go-first-steps/pkg/handler"
	"log"
)

func main() {
	srv := new(golang_first_steps.Server)
	handlers := new(handler.Handler)

	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
