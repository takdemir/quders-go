package routes

import (
	"github.com/labstack/echo/v4"
	"notificaiton/pkg/controllers"
)

func TeamsRoutes(g *echo.Group, h *controllers.Handler) {
	g.POST("/teams", h.SendTeamsMessage).Name = "send-teams-message"
}
