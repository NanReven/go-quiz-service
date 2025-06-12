package usecase

import (
	"QuizService/internal/domain"
	"QuizService/internal/repository"
	"QuizService/internal/service"
)

type Quiz interface {
}

type Question interface {
}

type User interface {
	Register(input *domain.User) (int, error)
	Login(input *domain.UserLogin) (string, string, error)
	Refresh(userId int) (string, string, error)
}

type Usecase struct {
	Quiz
	Question
	User
}

func NewUsecase(repo *repository.Repository, jwtService *service.JWTService) *Usecase {
	return &Usecase{
		User: NewAuthService(repo.User, jwtService),
	}
}
