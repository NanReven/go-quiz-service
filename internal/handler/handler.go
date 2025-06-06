package handler

import (
	"QuizService/internal/usecase"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	usecase *usecase.Usecase
}

func NewHandler(usecase *usecase.Usecase) *Handler {
	return &Handler{usecase: usecase}
}

func (h *Handler) InitRouter() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/register", h.Register)
		auth.POST("/login", h.Login)
		auth.POST("/logout", h.Logout)
	}

	api := router.Group("/api")
	{
		quiz := api.Group("/quiz")
		{
			quiz.GET("/", h.GetAllQuizes)
			quiz.GET("/:id", h.GetQuizById)
			quiz.POST("/", h.CreateQuiz)
			quiz.PUT("/:id", h.UpdateQuiz)
			quiz.DELETE("/:id", h.DeleteQuiz)
		}

		question := api.Group("/question")
		{
			question.GET("/", h.GetAllQuestions)
			question.GET("/:id", h.GetQuestionById)
			question.POST("/", h.CreateQuestion)
			question.PUT("/:id", h.UpdateQuestion)
			question.DELETE("/:id", h.DeleteQuestion)
		}

	}

	return router
}
