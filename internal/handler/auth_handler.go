package handler

import (
	"QuizService/internal/domain"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func (h *Handler) Register(c *gin.Context) {
	var input domain.User

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input", err.Error())
		return
	}

	id, err := h.usecase.User.Register(&input)
	if err != nil {
		if errors.Is(err, domain.ErrUserAlreadyExists) {
			newErrorResponse(c, http.StatusConflict, "email is already used", err.Error())
		} else {
			newErrorResponse(c, http.StatusInternalServerError, "internal server error", err.Error())
		}
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id": id,
	})
}

func (h *Handler) Login(c *gin.Context) {
	var input domain.UserLogin

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input", err.Error())
		return
	}

	accessToken, refreshToken, err := h.usecase.User.Login(&input)
	if err != nil {
		if errors.Is(err, domain.ErrUserNotFound) {
			newErrorResponse(c, http.StatusUnauthorized, "wrong email or password", err.Error())
		} else {
			newErrorResponse(c, http.StatusInternalServerError, "internal server error", err.Error())
		}
		return
	}

	c.SetCookie(viper.GetString("refresh_cookie_name"), refreshToken, viper.GetInt("cookie_max_age"), "/", "", true, true)

	c.JSON(http.StatusOK, gin.H{
		"access_token": accessToken,
	})
}
func (h *Handler) Refresh(c *gin.Context) {
	refreshToken, err := c.Cookie(viper.GetString("refresh_cookie_name"))
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, "refresh token isn`t set", err.Error())
		return
	}

	claims, err := h.jwtService.ParseToken(refreshToken)
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, "invalid refresh token", err.Error())
		return
	}

	accessToken, refreshToken, err := h.usecase.User.Refresh(claims.ID)
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, "invalid access token", err.Error())
		return
	}

	c.SetCookie(viper.GetString("refresh_cookie_name"), refreshToken, viper.GetInt("cookie_max_age"), "/", "", true, true)

	c.JSON(http.StatusOK, gin.H{
		"access_token": accessToken,
	})
}
