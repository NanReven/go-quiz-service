package usecase

import (
	"QuizService/internal/domain"
	"QuizService/internal/repository"
	"QuizService/internal/service"
)

type Quiz interface {
	CreateQuiz(input *domain.CreateQuizInput) (int, error)
	GetQuizById(quizID, userID int) (domain.Quiz, error)
	GetAllQuizes(userID int) ([]domain.Quiz, error)
	UpdateQuiz(input *domain.UpdateQuizInput) (domain.Quiz, error)
	DeleteQuiz(quizID, userID int) (domain.Quiz, error)
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
		Quiz: NewQuizUsecase(repo.Quiz),
	}
}
