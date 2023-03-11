package main

import (
	"database/pkg/config"
	"database/pkg/controllers"
	"database/pkg/controllers/middleware"
	"database/pkg/repository"
	"database/pkg/routes"
	utils "database/pkg/utils/jwt"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
)

var mysqlDB *gorm.DB

func init() {
	mysqlDB = config.Connect()
}
func main() {
	e := controllers.CreateNewRouter()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Welcome to main page of Database API!")
	})
	err := godotenv.Load()
	if err != nil {
		return
	}
	userRepository := repository.NewUserStore(mysqlDB)
	currencyRepository := repository.NewCurrencyStore(mysqlDB)
	handler := controllers.NewHandler(
		currencyRepository,
		userRepository,
	)
	g := e.Group("/api/v1", middleware.JWTWithConfig(utils.JWTSecretKey, handler))
	routes.CurrencyRoutes(g, handler)
	routes.CompanyRoutes(g, handler)
	e.Logger.Fatal(e.Start(":9029"))
}
