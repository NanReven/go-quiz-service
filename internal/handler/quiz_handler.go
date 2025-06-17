package handler

import (
	"QuizService/internal/domain"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetAllQuizes(c *gin.Context) {
	userID, ok := h.GetUserID(c)
	if !ok {
		return
	}

	quizzes, err := h.usecase.GetAllQuizes(userID)
	if err != nil {
		if errors.Is(err, domain.ErrQuizNotFound) {
			newErrorResponse(c, http.StatusNotFound, "quizzes not found", fmt.Sprintf("no quizzes found for user %d", userID))
		} else {
			newErrorResponse(c, http.StatusInternalServerError, "failed to get quizzes", err.Error())
		}
		return
	}

	c.JSON(http.StatusOK, quizzes)
}

func (h *Handler) GetQuizById(c *gin.Context) {
	quizID, ok := GetQuizID(c)
	if !ok {
		return
	}

	userID, ok := h.GetUserID(c)
	if !ok {
		return
	}

	quiz, err := h.usecase.GetQuizById(quizID, userID)
	if err != nil {
		if errors.Is(err, domain.ErrQuizNotFound) {
			newErrorResponse(c, http.StatusNotFound, "quiz not found", fmt.Sprintf("quiz %d not found for user %d", quizID, userID))
		} else {
			newErrorResponse(c, http.StatusInternalServerError, "failed to get quiz", err.Error())
		}
		return
	}

	c.JSON(http.StatusOK, quiz)
}

func (h *Handler) CreateQuiz(c *gin.Context) {
	var input domain.CreateQuizInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input data", err.Error())
		return
	}

	userID, ok := h.GetUserID(c)
	if !ok {
		return
	}
	input.AuthorID = userID

	quizID, err := h.usecase.Quiz.CreateQuiz(&input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "failed to create quiz", err.Error())
		return
	}

	c.JSON(http.StatusCreated, gin.H{"quiz_id": quizID})
}

func (h *Handler) UpdateQuiz(c *gin.Context) {
	var input domain.UpdateQuizInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input data", err.Error())
		return
	}

	quizID, ok := GetQuizID(c)
	if !ok {
		return
	}
	input.QuizID = quizID

	userID, ok := h.GetUserID(c)
	if !ok {
		return
	}
	input.AuthorID = userID

	updatedQuiz, err := h.usecase.Quiz.UpdateQuiz(&input)
	if err != nil {
		switch {
		case errors.Is(err, domain.ErrQuizNotFound):
			newErrorResponse(c, http.StatusNotFound, "quiz not found", fmt.Sprintf("quiz %d not found", quizID))
		case errors.Is(err, domain.ErrInvalidUpdateInput):
			newErrorResponse(c, http.StatusBadRequest, "invalid update input", "title and description are empty")
		default:
			newErrorResponse(c, http.StatusInternalServerError, "failed to update quiz", err.Error())
		}
		return
	}

	c.JSON(http.StatusOK, updatedQuiz)
}

func (h *Handler) DeleteQuiz(c *gin.Context) {
	quizID, ok := GetQuizID(c)
	if !ok {
		return
	}

	userID, ok := h.GetUserID(c)
	if !ok {
		return
	}

	quiz, err := h.usecase.Quiz.DeleteQuiz(quizID, userID)
	if err != nil {
		switch {
		case errors.Is(err, domain.ErrQuizNotFound):
			newErrorResponse(c, http.StatusNotFound, "quiz not found", fmt.Sprintf("quiz %d not found", quizID))
		case errors.Is(err, domain.ErrQuizAlreadyDeleted):
			newErrorResponse(c, http.StatusBadRequest, "quiz already deleted", fmt.Sprintf("quiz %d is already deleted", quizID))
		default:
			newErrorResponse(c, http.StatusInternalServerError, "failed to delete quiz", err.Error())
		}
		return
	}

	c.JSON(http.StatusOK, quiz)
}
