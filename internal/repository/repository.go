package repository

import (
	"QuizService/internal/domain"

	"github.com/jmoiron/sqlx"
)

type Quiz interface {
}

type Question interface {
}

type User interface {
	Register(input *domain.User) (int, error)
	GetUser(input *domain.UserLogin) (*domain.User, error)
}

type Repository struct {
	Quiz
	Question
	User
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		User: NewUserRepository(db),
	}
}
