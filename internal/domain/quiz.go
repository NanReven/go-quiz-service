package domain

import (
	"database/sql"
	"errors"
	"time"
)

var (
	ErrQuizNotFound       = errors.New("quiz not found")
	ErrInvalidUpdateInput = errors.New("invalid input")
	ErrQuizAlreadyDeleted = errors.New("quiz is deleted already")
)

type Quiz struct {
	ID          int          `json:"id" db:"id"`
	AuthorID    int          `json:"author_id" db:"author_id"`
	Title       string       `json:"title" db:"title"`
	Description string       `json:"description" db:"description"`
	CreatedAt   time.Time    `json:"created_at" db:"created_at"`
	DeletedAt   sql.NullTime `json:"deleted_at" db:"deleted_at"`
	IsActive    bool         `json:"is_active" db:"is_active"`
}

type CreateQuizInput struct {
	AuthorID    int    `db:"author_id"`
	Title       string `json:"title" binding:"required" db:"title"`
	Description string `json:"description" binding:"required" db:"description"`
}

type UpdateQuizInput struct {
	QuizID      int    `db:"id"`
	AuthorID    int    `db:"author_id"`
	Title       string `json:"title" db:"title"`
	Description string `json:"description" db:"description"`
}
