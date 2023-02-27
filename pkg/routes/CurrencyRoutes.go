package routes

import (
	"github.com/labstack/echo/v4"
	"quders/pkg/controllers"
)

func CurrencyRoutes(e *echo.Echo, handler *controllers.Handler) {
	g := e.Group("/api/v1")
	g.GET("/currency", handler.GetCurrencies)
	g.GET("/currency/:id", handler.GetCurrencyById)
	g.POST("/currency", handler.CreateCurrency)
	g.PUT("/currency", handler.UpdateCurrency)
}
