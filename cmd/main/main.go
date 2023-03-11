package main

import (
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"net/http"
	"quders/pkg/controllers"
	"quders/pkg/routes"
)

func main() {
	e := controllers.CreateNewRouter()
	err := godotenv.Load()
	if err != nil {
		return
	}
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Welcome to the main page :)")
	})
	h := controllers.NewHandler()
	//routes.CompanyRoutes(e,h)
	routes.CurrencyRoutes(e, h)
	e.Logger.Fatal(e.Start(":9031"))

}
