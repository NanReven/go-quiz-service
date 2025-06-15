package domain

import "time"

type Quiz struct {
	ID          int       `json:"id" db:"id"`
	AuthorID    int       `json:"author_id" db:"author_id"`
	Title       string    `json:"title" binding:"required" db:"title"`
	Description string    `json:"description" binding:"required" db:"description"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	DeletedAt   time.Time `json:"deleted_at" db:"deleted_at"`
	IsActive    bool      `json:"is_active" db:"is_active"`
}

type CreateQuiz struct {
	AuthorID    int    `json:"author_id" db:"author_id"`
	Title       string `json:"title" binding:"required" db:"title"`
	Description string `json:"description" binding:"required" db:"description"`
}
