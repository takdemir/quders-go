package routes

import (
	"database/pkg/controllers"
	"github.com/labstack/echo/v4"
)

func CountryRoutes(g *echo.Group, handler *controllers.Handler) {
	g.GET("/country", handler.GetCountries)
	g.GET("/country/:id", handler.GetCountryById)
	g.POST("/country", handler.CreateCountry)
	g.PUT("/country/:id", handler.UpdateCountry)
}
