package usecase

import (
	"QuizService/internal/domain"
	"QuizService/internal/repository"
)

type Quiz interface {
}

type Question interface {
}

type User interface {
	Register(input *domain.User) (int, error)
}

type Usecase struct {
	Quiz
	Question
	User
}

func NewUsecase(repo *repository.Repository) *Usecase {
	return &Usecase{
		User: NewAuthService(repo.User),
	}
}
