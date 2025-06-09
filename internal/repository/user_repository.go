package repository

import (
	"QuizService/internal/domain"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (repo *UserRepository) Register(input *domain.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (email, first_name, second_name, password_hash) VALUES ($1, $2, $3, $4) RETURNING id", usersTable)
	row := repo.db.QueryRow(query, input.Email, input.FirstName, input.SecondName, input.Password)
	if err := row.Scan(&id); err != nil {
		return 0, fmt.Errorf("failed to create user with email %s: %w", input.Email, err)
	}
	return id, nil
}
