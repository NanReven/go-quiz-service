package repository

import (
	"QuizService/internal/domain"
	"database/sql"
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type QuizRepository struct {
	db *sqlx.DB
}

func NewQuizRepository(db *sqlx.DB) *QuizRepository {
	return &QuizRepository{db: db}
}

func (repo *QuizRepository) CreateQuiz(input *domain.CreateQuizInput) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (author_id, title, description) VALUES ($1, $2, $3) RETURNING id", quizzesTable)
	row := repo.db.QueryRow(query, input.AuthorID, input.Title, input.Description)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (repo *QuizRepository) GetQuizById(quizID, userID int) (domain.Quiz, error) {
	var quiz domain.Quiz
	query := fmt.Sprintf("SELECT * FROM %s WHERE id=$1 AND author_id=$2 AND deleted_at IS NULL", quizzesTable)
	if err := repo.db.Get(&quiz, query, quizID, userID); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return quiz, domain.ErrQuizNotFound
		}
		return quiz, err
	}
	return quiz, nil
}

func (repo *QuizRepository) GetAllQuizes(userID int) ([]domain.Quiz, error) {
	var quizzes []domain.Quiz
	query := fmt.Sprintf("SELECT * FROM %s WHERE author_id=$1 AND deleted_at IS NULL", quizzesTable)
	if err := repo.db.Select(&quizzes, query, userID); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return quizzes, domain.ErrQuizNotFound
		}
		return quizzes, err
	}
	return quizzes, nil
}

func (repo *QuizRepository) UpdateQuizTitle(input *domain.UpdateQuizInput) error {
	query := fmt.Sprintf("UPDATE %s SET title=$1 WHERE id=$2 AND author_id=$3 AND deleted_at IS NULL", quizzesTable)
	res, err := repo.db.Exec(query, input.Title, input.QuizID, input.AuthorID)
	if err != nil {
		return err
	}
	rows, _ := res.RowsAffected()
	if rows == 0 {
		return domain.ErrQuizNotFound
	}
	return nil
}

func (repo *QuizRepository) UpdateQuizDescription(input *domain.UpdateQuizInput) error {
	query := fmt.Sprintf("UPDATE %s SET description=$1 WHERE id=$2 AND author_id=$3 AND deleted_at IS NULL", quizzesTable)
	res, err := repo.db.Exec(query, input.Description, input.QuizID, input.AuthorID)
	if err != nil {
		return err
	}
	rows, _ := res.RowsAffected()
	if rows == 0 {
		return domain.ErrQuizNotFound
	}
	return nil
}

func (repo *QuizRepository) CheckQuizStatus(quizID, userID int) (bool, error) {
	var status bool
	query := fmt.Sprintf("SELECT is_active FROM %s WHERE id=$1 AND author_id=$2 AND deleted_at IS NULL", quizzesTable)
	row := repo.db.QueryRow(query, quizID, userID)
	if err := row.Scan(&status); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return status, domain.ErrQuizNotFound
		}
		return status, err
	}
	return status, nil
}

func (repo *QuizRepository) DeleteQuiz(quizID, userID int) error {
	query := fmt.Sprintf("UPDATE %s SET deleted_at=CURRENT_TIMESTAMP, is_active=false WHERE id=$1 AND author_id=$2", quizzesTable)
	if _, err := repo.db.Exec(query, quizID, userID); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.ErrQuizNotFound
		}
		return err
	}
	return nil
}
