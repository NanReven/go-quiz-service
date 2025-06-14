package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func newErrorResponse(c *gin.Context, statusCode int, userMessage string, logMessage string) {
	logrus.WithFields(logrus.Fields{
		"method":      c.Request.Method,
		"path":        c.Request.URL.Path,
		"status_code": statusCode,
		"message":     logMessage,
	}).Error("error")

	c.AbortWithStatusJSON(statusCode, gin.H{"error_message": userMessage})
}
