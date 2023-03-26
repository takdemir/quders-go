package main

import (
	"database/pkg/config"
	"database/pkg/controllers"
	"database/pkg/controllers/middleware"
	"database/pkg/repository"
	"database/pkg/routes"
	utils "database/pkg/utils/jwt"
	"database/pkg/utils/redis"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
)

var mysqlDB *gorm.DB
var qudersRedis *redis.QudersRedis

func init() {
	mysqlDB = config.Connect()
	err := mysqlDB.AutoMigrate()
	if err != nil {
		panic("DB auto migrate error: " + err.Error())
	}
	qudersRedis.Connect("REDIS_9034")
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
	currencyRepository := repository.NewCurrencyStore(mysqlDB, qudersRedis)
	companyRepository := repository.NewCompanyStore(mysqlDB, qudersRedis)
	countryRepository := repository.NewCountryStore(mysqlDB, qudersRedis)
	companyDetailRepository := repository.NewCompanyDetailStore(mysqlDB, qudersRedis)
	agreementAndTermsRepository := repository.NewAgreementAndTermsStore(mysqlDB, qudersRedis)
	notificationEventRepository := repository.NewNotificationEventStore(mysqlDB, qudersRedis)
	notificationRepository := repository.NewNotificationStore(mysqlDB, qudersRedis)
	programingLanguageRepository := repository.NewProgramingLanguageStore(mysqlDB, qudersRedis)
	frameworkRepository := repository.NewFrameworkStore(mysqlDB, qudersRedis)
	registrationQuestionRepository := repository.NewRegistrationQuestionStore(mysqlDB, qudersRedis)
	questionAnswerRepository := repository.NewQuestionAnswerStore(mysqlDB, qudersRedis)
	handler := controllers.NewHandler(
		mysqlDB,
		currencyRepository,
		userRepository,
		companyRepository,
		countryRepository,
		companyDetailRepository,
		agreementAndTermsRepository,
		notificationEventRepository,
		notificationRepository,
		programingLanguageRepository,
		frameworkRepository,
		registrationQuestionRepository,
		questionAnswerRepository,
	)
	g := e.Group("/api/v1", middleware.JWTWithConfig(utils.JWTSecretKey, handler))
	routes.CurrencyRoutes(g, handler)
	routes.CompanyRoutes(g, handler)
	routes.CountryRoutes(g, handler)
	routes.CompanyDetailRoutes(g, handler)
	routes.AgreementAndTermsRoutes(g, handler)
	routes.NotificationRoutes(g, handler)
	routes.ProgramingLanguageRoutes(g, handler)
	routes.FrameworkRoutes(g, handler)
	routes.RegistrationQuestionRoutes(g, handler)
	routes.QuestionAnswerRoutes(g, handler)
	e.Logger.Fatal(e.Start(":9029"))
}
