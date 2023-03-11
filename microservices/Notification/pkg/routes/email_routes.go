package routes

import (
	"github.com/labstack/echo/v4"
	"notificaiton/pkg/controllers"
)

func EmailRoutes(g *echo.Group, h *controllers.Handler) {
	g.POST("/email", h.SendEmail).Name = "send-email"
}
