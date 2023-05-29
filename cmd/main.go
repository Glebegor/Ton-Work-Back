package main

import (
	TonWork "github.com/TonWork/back"
	handlers "github.com/TonWork/back/pkg/handler"
	repositoryes "github.com/TonWork/back/pkg/repository"
	services "github.com/TonWork/back/pkg/service"
)

func main() {
	repository := repositoryes.Repository(db)
	service := services.Service(repository)
	handler := handlers.Handler(service)
	server := new(TonWork.Server)
	if err := server.Run("8000", handler); err != nil {

	}
	return
}
