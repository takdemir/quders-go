package routes

import (
	"database/pkg/controllers"
	"github.com/labstack/echo/v4"
)

func FrameworkRoutes(g *echo.Group, handler *controllers.Handler) {
	g.GET("/framework", handler.GetFrameworks)
	g.GET("/framework/:id", handler.GetFrameworkById)
	g.POST("/framework", handler.CreateFramework)
	g.PUT("/framework/:id", handler.UpdateFramework)
}
