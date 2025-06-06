package repository

type Quiz interface {
}

type Question interface {
}

type User interface {
}

type Repository struct {
	Quiz
	Question
	User
}

func NewRepository() *Repository {
	return &Repository{}
}
