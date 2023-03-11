package routes

import (
	"database/pkg/controllers"
	"github.com/labstack/echo/v4"
)

func CurrencyRoutes(g *echo.Group, handler *controllers.Handler) {
	g.GET("/currency", handler.GetCurrencies)
	g.GET("/currency/:id", handler.GetCurrencyById)
	g.POST("/currency", handler.CreateCurrency)
	g.PUT("/currency", handler.UpdateCurrency)
}
