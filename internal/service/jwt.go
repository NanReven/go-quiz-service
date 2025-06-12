package service

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
)

type JWT interface {
	GenerateAccessToken(userID int) (string, error)
	GenerateRefreshToken(userID int) (string, error)
	ParseToken(tokenString string) (*UserClaims, error)
}

type UserClaims struct {
	ID int
	jwt.RegisteredClaims
}

type JWTService struct {
	secret []byte
}

func NewJWTService() *JWTService {
	secret := []byte(os.Getenv("JWT_SECRET"))
	return &JWTService{secret: secret}
}

func (s *JWTService) GenerateAccessToken(userID int) (string, error) {
	return generateToken(userID, viper.GetDuration("access_token_ttl"), s.secret)
}

func (s *JWTService) GenerateRefreshToken(userID int) (string, error) {
	return generateToken(userID, viper.GetDuration("refresh_token_ttl"), s.secret)
}

func (s *JWTService) ParseToken(tokenString string) (*UserClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return s.secret, nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*UserClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}

func generateToken(userID int, duration time.Duration, secret []byte) (string, error) {
	claims := UserClaims{
		ID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(duration)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secret)
}
