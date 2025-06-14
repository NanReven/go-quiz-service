package repository

import (
	"QuizService/internal/domain"
	"database/sql"
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (repo *UserRepository) CheckUserExists(email string) bool {
	var result bool
	query := fmt.Sprintf("SELECT 1 FROM %s WHERE email=$1", usersTable)
	row := repo.db.QueryRow(query, email)
	if err := row.Scan(&result); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false
		}
	}
	return true
}

func (repo *UserRepository) Register(input *domain.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (email, first_name, second_name, password_hash) VALUES ($1, $2, $3, $4) RETURNING id", usersTable)
	row := repo.db.QueryRow(query, input.Email, input.FirstName, input.SecondName, input.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (repo *UserRepository) GetUser(input *domain.UserLogin) (*domain.User, error) {
	var user domain.User
	query := fmt.Sprintf("SELECT * FROM %s WHERE email=$1 AND password_hash=$2", usersTable)
	if err := repo.db.Get(&user, query, input.Email, input.Password); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return &user, domain.ErrUserNotFound
		} else {
			return &user, err
		}
	}
	return &user, nil
}
