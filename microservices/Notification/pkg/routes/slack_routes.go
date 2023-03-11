package routes

import (
	"github.com/labstack/echo/v4"
	"notificaiton/pkg/controllers"
)

func SlackRoutes(g *echo.Group, h *controllers.Handler) {
	g.POST("/slack", h.SendSlackMessage).Name = "send-slack-message"
}
