package routes

import (
	"database/pkg/controllers"
	"github.com/labstack/echo/v4"
)

func RegistrationQuestionRoutes(g *echo.Group, handler *controllers.Handler) {
	g.GET("/registration-question", handler.GetRegistrationQuestions)
	g.GET("/registration-question/:id", handler.GetRegistrationQuestionById)
	g.POST("/registration-question", handler.CreateRegistrationQuestion)
	g.PUT("/registration-question/:id", handler.UpdateRegistrationQuestion)
}
