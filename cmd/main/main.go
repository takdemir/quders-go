package main

import (
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"net/http"
	"os"
	"quders/pkg/controllers"
	"quders/pkg/repository"
	"quders/pkg/routes"
)

var mysqlqudersDB *gorm.DB

func ConnectToMySQLQuders() *gorm.DB {
	db, err := gorm.Open(mysql.New(mysql.Config{DSN: os.Getenv("MYSQL_DB_URL")}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		panic(err)
	}
	return db
}

func main() {
	e := controllers.CreateNewRouter()
	err := godotenv.Load()
	mysqlqudersDB = ConnectToMySQLQuders()
	if err != nil {
		return
	}
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Welcome to the main page :)")
	})
	currencyRepository := repository.NewCurrencyStore(mysqlqudersDB)
	h := controllers.NewHandler(currencyRepository)
	//routes.CompanyRoutes(e,h)
	routes.CurrencyRoutes(e, h)
	e.Logger.Fatal(e.Start(":9031"))

}
