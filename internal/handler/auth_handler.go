package handler

import (
	"QuizService/internal/domain"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func (h *Handler) Register(c *gin.Context) {
	var input domain.User

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.usecase.User.Register(&input)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id": id,
	})
}

func (h *Handler) Login(c *gin.Context) {
	var input domain.UserLogin

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	accessToken, refreshToken, err := h.usecase.User.Login(&input)
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
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
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	claims, err := h.jwtService.ParseToken(refreshToken)
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	accessToken, refreshToken, err := h.usecase.User.Refresh(claims.ID)
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	c.SetCookie(viper.GetString("refresh_cookie_name"), refreshToken, viper.GetInt("cookie_max_age"), "/", "", true, true)

	c.JSON(http.StatusOK, gin.H{
		"access_token": accessToken,
	})
}
