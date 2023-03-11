package main

import (
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"net/http"
	"notificaiton/pkg/controllers"
	"notificaiton/pkg/controllers/middleware"
	"notificaiton/pkg/routes"
	utils "notificaiton/pkg/utils/jwt"
)

func main() {
	e := controllers.CreateNewRouter()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Welcome to main page of Database API!")
	})
	err := godotenv.Load()
	if err != nil {
		return
	}

	handler := controllers.NewHandler()
	g := e.Group("/api/v1", middleware.JWTWithConfig(utils.JWTSecretKey, handler))
	routes.EmailRoutes(g, handler)
	routes.SlackRoutes(g, handler)
	routes.SmsRoutes(g, handler)
	routes.TeamsRoutes(g, handler)
	e.Logger.Fatal(e.Start(":9023"))
}
