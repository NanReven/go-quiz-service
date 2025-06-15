package usecase

import (
	"QuizService/internal/domain"
	"QuizService/internal/repository"
)

type QuizUsecase struct {
	repo repository.Quiz
}

func NewQuizUsecase(repo repository.Quiz) *QuizUsecase {
	return &QuizUsecase{repo: repo}
}

func (uc *QuizUsecase) CreateQuiz(input *domain.CreateQuiz) (int, error) {
	return uc.repo.CreateQuiz(input)
}
