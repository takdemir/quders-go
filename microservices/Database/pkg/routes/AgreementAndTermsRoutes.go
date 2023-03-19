package routes

import (
	"database/pkg/controllers"
	"github.com/labstack/echo/v4"
)

func AgreementAndTermsRoutes(g *echo.Group, handler *controllers.Handler) {
	g.GET("/agreement-and-terms", handler.FetchAgreementAndTerms)
}
