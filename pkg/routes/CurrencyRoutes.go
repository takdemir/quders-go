package routes

import (
	"github.com/labstack/echo/v4"
	"quders/pkg/controllers"
)

func CurrencyRoutes(e *echo.Echo) {
	e.GET("/api/v1/currency", controllers.GetCurrencies)
	e.GET("/api/v1/currency/:id", controllers.GetCurrencyById)
	e.POST("/api/v1/currency", controllers.CreateCurrency)
	e.PUT("/api/v1/currency", controllers.UpdateCurrency)
}
