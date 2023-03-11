package routes

import (
	"github.com/labstack/echo/v4"
	"notificaiton/pkg/controllers"
)

func SmsRoutes(g *echo.Group, h *controllers.Handler) {
	g.POST("/sms", h.SendSms).Name = "send-sms"
}
