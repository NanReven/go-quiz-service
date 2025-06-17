package repository

import (
	"QuizService/internal/domain"

	"github.com/jmoiron/sqlx"
)

type Quiz interface {
	CreateQuiz(input *domain.CreateQuizInput) (int, error)
	GetQuizById(quizID, userID int) (domain.Quiz, error)
	GetAllQuizes(userID int) ([]domain.Quiz, error)
	UpdateQuizTitle(input *domain.UpdateQuizInput) error
	UpdateQuizDescription(input *domain.UpdateQuizInput) error
	DeleteQuiz(quizID, userID int) error
	CheckQuizStatus(quizID, userID int) (bool, error)
}

type Question interface {
}

type User interface {
	Register(input *domain.User) (int, error)
	GetUser(input *domain.UserLogin) (*domain.User, error)
	CheckUserExists(email string) bool
}

type Repository struct {
	Quiz
	Question
	User
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		User: NewUserRepository(db),
		Quiz: NewQuizRepository(db),
	}
}
