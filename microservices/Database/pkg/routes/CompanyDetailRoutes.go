package routes

import (
	"database/pkg/controllers"
	"github.com/labstack/echo/v4"
)

func CompanyDetailRoutes(g *echo.Group, handler *controllers.Handler) {
	g.GET("/company-detail/:id", handler.FetchCompanyDetailByCompanyId)
	g.POST("/company-detail", handler.CreateCompanyDetail)
	g.PUT("/company-detail/:id", handler.UpdateCompanyDetail)
}
