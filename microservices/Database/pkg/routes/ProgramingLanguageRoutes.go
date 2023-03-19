package routes

import (
	"database/pkg/controllers"
	"github.com/labstack/echo/v4"
)

func ProgramingLanguageRoutes(g *echo.Group, handler *controllers.Handler) {
	g.GET("/programing-language", handler.GetProgramingLanguages)
	g.GET("/programing-language/:id", handler.GetProgramingLanguageById)
	g.POST("/programing-language", handler.CreateProgramingLanguage)
	g.PUT("/programing-language/:id", handler.UpdateProgramingLanguage)
}
