package handler

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func (h *Handler) AuthMiddleware(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if !strings.HasPrefix(authHeader, "Bearer ") {
		newErrorResponse(c, http.StatusUnauthorized, "invalid auth header")
		return
	}

	authHeaderData := strings.Split(authHeader, " ")
	if len(authHeaderData) != 2 {
		newErrorResponse(c, http.StatusUnauthorized, "invalid auth header format")
		return
	}

	claims, err := h.jwtService.ParseToken(authHeaderData[1])
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	c.Set("userId", claims.ID)
	c.Next()
}

func (h *Handler) GetUserId(c *gin.Context) (int, bool) {
	id, ok := c.Get("userId")
	if !ok {
		newErrorResponse(c, http.StatusUnauthorized, "user is unauthorized")
		return 0, false
	}

	userId, ok := id.(int)
	if !ok {
		newErrorResponse(c, http.StatusBadRequest, "invalid user id")
		return 0, false
	}

	return userId, true
}
