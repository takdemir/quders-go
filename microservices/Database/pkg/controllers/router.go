package controllers

import (
	"database/pkg/repository"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

type Handler struct {
	CurrencyRepository repository.CurrencyRepository
	UserRepository     repository.UserRepository
}

func CreateNewRouter() *echo.Echo {
	e := echo.New()
	e.Logger.SetLevel(log.DEBUG)
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Logger())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))
	return e
}

func NewHandler(
	currencyRepo repository.CurrencyRepository,
	userRepo repository.UserRepository,
) *Handler {
	return &Handler{
		CurrencyRepository: currencyRepo,
		UserRepository:     userRepo,
	}
}
