package handler

import (
	"QuizService/internal/domain"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetAllQuizes(c *gin.Context) {

}

func (h *Handler) GetQuizById(c *gin.Context) {

}

func (h *Handler) CreateQuiz(c *gin.Context) {
	var input domain.CreateQuiz
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input data", err.Error())
		return
	}

	userID, ok := h.GetUserId(c)
	if !ok {
		newErrorResponse(c, http.StatusUnauthorized, "unauthorized", "invalid user id:"+fmt.Sprint(userID))
		return
	}
	input.AuthorID = userID

	quizID, err := h.usecase.Quiz.CreateQuiz(&input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "can not create quiz", err.Error())
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"quiz_id": quizID,
	})
}

func (h *Handler) UpdateQuiz(c *gin.Context) {

}

func (h *Handler) DeleteQuiz(c *gin.Context) {

}
