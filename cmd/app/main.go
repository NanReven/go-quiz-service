package main

import (
	"QuizService/internal/handler"
	"QuizService/internal/infrastructure"
	"QuizService/internal/repository"
	"QuizService/internal/service"
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

	jwtService := service.NewJWTService()

	repository := repository.NewRepository(db)
	usecase := usecase.NewUsecase(repository, jwtService)
	handler := handler.NewHandler(usecase, jwtService)

	server := new(infrastructure.Server)
	log.Fatal(server.Run(viper.GetString("port"), handler.InitRouter()))
}

func initConfig() error {
	viper.AddConfigPath("config")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
