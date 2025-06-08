package main

import (
	"QuizService/internal/handler"
	"QuizService/internal/infrastructure"
	"QuizService/internal/usecase"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	//repository := repository.NewRepository()
	usecase := usecase.NewUsecase()
	handler := handler.NewHandler(usecase)

	server := new(infrastructure.Server)
	log.Fatal(server.Run("8000", handler.InitRouter()))
}
