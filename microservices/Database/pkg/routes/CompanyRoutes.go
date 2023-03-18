package routes

import (
	"database/pkg/controllers"
	"github.com/labstack/echo/v4"
)

func CompanyRoutes(g *echo.Group, handler *controllers.Handler) {
	g.GET("/company", handler.GetCompanies)
	g.GET("/company/:id", handler.GetCompanyById)
	g.POST("/company", handler.CreateCompany)
	g.PUT("/company/:id", handler.UpdateCompany)
}
