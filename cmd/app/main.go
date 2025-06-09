package main

import (
	"QuizService/internal/handler"
	"QuizService/internal/infrastructure"
	"QuizService/internal/repository"
	"QuizService/internal/usecase"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	if err := initConfig(); err != nil {
		logrus.Fatal("failed to parse config", err.Error())
	}

	if err := godotenv.Load(".env"); err != nil {
		logrus.Fatal("failed to load env file", err.Error())
	}

	db, err := repository.NewPostgresDB(&repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: os.Getenv("DB_PASSWORD"),
	})

	if err != nil {
		logrus.Fatal("failed to connect to db", err.Error())
	}

	repository := repository.NewRepository(db)
	usecase := usecase.NewUsecase(repository)
	handler := handler.NewHandler(usecase)

	server := new(infrastructure.Server)
	log.Fatal(server.Run("8000", handler.InitRouter()))
}

func initConfig() error {
	viper.AddConfigPath("config")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
