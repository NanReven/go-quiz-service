package repository

import (
	"QuizService/internal/domain"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type QuizRepository struct {
	db *sqlx.DB
}

func NewQuizRepository(db *sqlx.DB) *QuizRepository {
	return &QuizRepository{db: db}
}

func (repo *QuizRepository) CreateQuiz(input *domain.CreateQuiz) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (author_id, title, description) VALUES ($1, $2, $3) RETURNING id", quizzesTable)
	row := repo.db.QueryRow(query, input.AuthorID, input.Title, input.Description)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}
