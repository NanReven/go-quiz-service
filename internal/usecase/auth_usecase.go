package usecase

import (
	"QuizService/internal/domain"
	"QuizService/internal/repository"
	"QuizService/internal/service"
	"crypto/sha1"
	"fmt"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

type AuthUsecase struct {
	repo       repository.User
	jwtService service.JWT
}

type UserClaims struct {
	ID int
	jwt.RegisteredClaims
}

func NewAuthService(repo repository.User, jwtService service.JWT) *AuthUsecase {
	return &AuthUsecase{repo: repo, jwtService: jwtService}
}

func (uc *AuthUsecase) Register(input *domain.User) (int, error) {
	if uc.repo.CheckUserExists(input.Email) {
		return 0, domain.ErrUserAlreadyExists
	}
	input.Password = generatePasswordHash(input.Password)
	return uc.repo.Register(input)
}

func (uc *AuthUsecase) Login(input *domain.UserLogin) (string, string, error) {
	input.Password = generatePasswordHash(input.Password)

	user, err := uc.repo.GetUser(input)
	if err != nil {
		return "", "", err
	}

	accessToken, err := uc.jwtService.GenerateAccessToken(user.ID)
	if err != nil {
		return "", "", err
	}

	refreshToken, err := uc.jwtService.GenerateRefreshToken(user.ID)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

func (uc *AuthUsecase) Refresh(userId int) (string, string, error) {
	accessToken, err := uc.jwtService.GenerateAccessToken(userId)
	if err != nil {
		return "", "", err
	}

	refreshToken, err := uc.jwtService.GenerateRefreshToken(userId)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(os.Getenv("PASSWORD_SALT"))))
}
