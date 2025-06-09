package usecase

import (
	"QuizService/internal/domain"
	"QuizService/internal/repository"
	"crypto/sha1"
	"fmt"
	"os"
)

type AuthUsecase struct {
	repo repository.User
}

func NewAuthService(repo repository.User) *AuthUsecase {
	return &AuthUsecase{repo: repo}
}

func (uc *AuthUsecase) Register(input *domain.User) (int, error) {
	input.Password = generatePasswordHash(input.Password)
	return uc.repo.Register(input)
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(os.Getenv("PASSWORD_SALT"))))
}
