package routes

import (
	"github.com/labstack/echo/v4"
	"quders/pkg/controllers"
)

func CompanyRoutes(e *echo.Echo) {
	g := e.Group("/api/v1")
	g.GET("/company", controllers.GetCompanies)
	g.GET("/company/:id", controllers.GetCompanyById)
	g.POST("/company", controllers.CreateCompany)
	g.PUT("/company", controllers.UpdateCompany)
}
