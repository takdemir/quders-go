package routes

import (
	"github.com/labstack/echo/v4"
	"quders/pkg/controllers"
)

func CompanyRoutes(e *echo.Echo) {
	e.GET("/api/v1/company", controllers.GetCompanies)
	e.GET("/api/v1/company/:id", controllers.GetCompanyById)
	e.POST("/api/v1/company", controllers.CreateCompany)
	e.PUT("/api/v1/company", controllers.UpdateCompany)
}
