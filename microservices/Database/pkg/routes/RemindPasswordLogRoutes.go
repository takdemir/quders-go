package routes

import (
	"database/pkg/controllers"
	"github.com/labstack/echo/v4"
)

func RemindPasswordLogRoutes(g *echo.Group, handler *controllers.Handler) {
	g.GET("/check-one-time-code", handler.CheckOneTimeCode)
	g.POST("/one-time-code", handler.CreateOneTimeCode)
	g.PUT("/one-time-code", handler.UpdateOneTimeCode)
}
