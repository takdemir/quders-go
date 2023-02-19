package main

import (
	"github.com/go-playground/validator"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"net/http"
	"quders/pkg/routes"
)

type (
	CustomValidator struct {
		validator *validator.Validate
	}
)

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		return err
	}
	return nil
}

func main() {
	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}
	err := godotenv.Load()
	if err != nil {
		return
	}
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Welcome to the main page :)")
	})
	routes.CompanyRoutes(e)
	routes.CurrencyRoutes(e)
	e.Logger.Fatal(e.Start(":9031"))

}
