package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func newErrorResponse(c *gin.Context, statusCode int, message string) {
	logrus.Println(message)
	c.AbortWithStatusJSON(statusCode, gin.H{"error_message": message})
}
