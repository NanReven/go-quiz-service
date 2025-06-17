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

func (uc *QuizUsecase) CreateQuiz(input *domain.CreateQuizInput) (int, error) {
	return uc.repo.CreateQuiz(input)
}

func (uc *QuizUsecase) GetQuizById(quizID, userID int) (domain.Quiz, error) {
	return uc.repo.GetQuizById(quizID, userID)
}

func (uc *QuizUsecase) GetAllQuizes(userID int) ([]domain.Quiz, error) {
	return uc.repo.GetAllQuizes(userID)
}

func (uc *QuizUsecase) DeleteQuiz(quizID, userID int) (domain.Quiz, error) {
	isActive, err := uc.repo.CheckQuizStatus(quizID, userID)
	if err != nil {
		return domain.Quiz{}, err
	}

	if !isActive {
		return domain.Quiz{}, domain.ErrQuizAlreadyDeleted
	}

	if err := uc.repo.DeleteQuiz(quizID, userID); err != nil {
		return domain.Quiz{}, err
	}
	return uc.repo.GetQuizById(quizID, userID)
}

func (uc *QuizUsecase) UpdateQuiz(input *domain.UpdateQuizInput) (domain.Quiz, error) {
	if input.Title == "" && input.Description == "" {
		return domain.Quiz{}, domain.ErrInvalidUpdateInput
	}

	if input.Title != "" {
		if err := uc.repo.UpdateQuizTitle(input); err != nil {
			return domain.Quiz{}, err
		}
	}

	if input.Description != "" {
		if err := uc.repo.UpdateQuizDescription(input); err != nil {
			return domain.Quiz{}, err
		}
	}

	quiz, err := uc.repo.GetQuizById(input.QuizID, input.AuthorID)
	if err != nil {
		return quiz, err
	}
	return quiz, nil
}
