package routes

import (
	"database/pkg/controllers"
	"github.com/labstack/echo/v4"
)

func QuestionAnswerRoutes(g *echo.Group, handler *controllers.Handler) {
	g.GET("/question-answer", handler.GetQuestionAnswers)
	g.GET("/question-answer/:id", handler.GetQuestionAnswerById)
	g.POST("/question-answer", handler.CreateQuestionAnswer)
	g.PUT("/question-answer/:id", handler.UpdateQuestionAnswer)
}
