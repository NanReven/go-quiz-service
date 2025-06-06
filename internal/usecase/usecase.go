package usecase

type Quiz interface {
}

type Question interface {
}

type User interface {
}

type Usecase struct {
	Quiz
	Question
	User
}

func NewUsecase() *Usecase {
	return &Usecase{}
}
