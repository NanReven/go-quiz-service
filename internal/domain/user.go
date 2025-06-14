package domain

import "errors"

var (
	ErrUserNotFound      = errors.New("user not found")
	ErrUserAlreadyExists = errors.New("user is already exists")
)

type User struct {
	ID         int    `json:"id" db:"id"`
	Email      string `json:"email" binding:"required" db:"email"`
	FirstName  string `json:"first_name" binding:"required" db:"first_name"`
	SecondName string `json:"second_name" binding:"required" db:"second_name"`
	Password   string `json:"password" binding:"required" db:"password_hash"`
}

type UserLogin struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
