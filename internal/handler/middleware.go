package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func (h *Handler) AuthMiddleware(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if !strings.HasPrefix(authHeader, "Bearer ") {
		newErrorResponse(c, http.StatusUnauthorized, "authorization header is invalid", "invalid authorization header: "+authHeader)
		return
	}

	authHeaderData := strings.Split(authHeader, " ")
	if len(authHeaderData) != 2 {
		newErrorResponse(c, http.StatusUnauthorized, "authorization header is invalid", "invalid authorization header: "+authHeader)
		return
	}

	claims, err := h.jwtService.ParseToken(authHeaderData[1])
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, "access token is invalid or expired", "failed to parse access token: "+err.Error())
		return
	}

	c.Set("userId", claims.ID)
	c.Next()
}

func (h *Handler) GetUserID(c *gin.Context) (int, bool) {
	id, ok := c.Get("userId")
	if !ok {
		newErrorResponse(c, http.StatusUnauthorized, "user is not authorized", "user ID is not set")
		return 0, false
	}

	userId, ok := id.(int)
	if !ok {
		newErrorResponse(c, http.StatusBadRequest, "user ID is invalid", "user ID is invalid: "+fmt.Sprint(id))
		return 0, false
	}

	return userId, true
}

func GetQuizID(c *gin.Context) (int, bool) {
	id := c.Param("id")
	quizID, err := strconv.Atoi(id)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid quiz id", err.Error())
		return 0, false
	}
	return quizID, true
}
