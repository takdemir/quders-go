package routes

import (
	"database/pkg/controllers"
	"github.com/labstack/echo/v4"
)

func NotificationRoutes(g *echo.Group, h *controllers.Handler) {
	g.GET("/notification", h.GetNotifications).Name = "fetch-notifications"
	g.GET("/:id/notification", h.GetNotificationById).Name = "fetch-notification-by-id"
	g.GET("/active/notification", h.GetActiveNotifications).Name = "fetch-active-notifications"
}
