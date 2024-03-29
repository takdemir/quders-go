package controllers

import (
	"database/pkg/repository"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type Handler struct {
	DB                             *gorm.DB
	CurrencyRepository             repository.CurrencyRepository
	UserRepository                 repository.UserRepository
	CompanyRepository              repository.CompanyRepository
	CountryRepository              repository.CountryRepository
	CompanyDetailRepository        repository.CompanyDetailRepository
	AgreementAndTermsRepository    repository.AgreementAndTermsRepository
	NotificationEventRepository    repository.NotificationEventRepository
	NotificationRepository         repository.NotificationRepository
	ProgramingLanguageRepository   repository.ProgramingLanguageRepository
	FrameworkRepository            repository.FrameworkRepository
	RegistrationQuestionRepository repository.RegistrationQuestionRepository
	QuestionAnswerRepository       repository.QuestionAnswerRepository
	RemindPasswordLogRepository    repository.RemindPasswordLogRepository
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
	db *gorm.DB,
	currencyRepo repository.CurrencyRepository,
	userRepo repository.UserRepository,
	companyRepo repository.CompanyRepository,
	countryRepo repository.CountryRepository,
	companyDetailRepo repository.CompanyDetailRepository,
	agreementAndTermsRepo repository.AgreementAndTermsRepository,
	notificationEventRepo repository.NotificationEventRepository,
	notificationRepo repository.NotificationRepository,
	programingLanguageRepo repository.ProgramingLanguageRepository,
	frameworkRepo repository.FrameworkRepository,
	registrationQuestionRepo repository.RegistrationQuestionRepository,
	questionAnswerRepo repository.QuestionAnswerRepository,
	remindPasswordLogRepo repository.RemindPasswordLogRepository,
) *Handler {
	return &Handler{
		DB:                             db,
		CurrencyRepository:             currencyRepo,
		UserRepository:                 userRepo,
		CompanyRepository:              companyRepo,
		CountryRepository:              countryRepo,
		CompanyDetailRepository:        companyDetailRepo,
		AgreementAndTermsRepository:    agreementAndTermsRepo,
		NotificationEventRepository:    notificationEventRepo,
		NotificationRepository:         notificationRepo,
		ProgramingLanguageRepository:   programingLanguageRepo,
		FrameworkRepository:            frameworkRepo,
		RegistrationQuestionRepository: registrationQuestionRepo,
		QuestionAnswerRepository:       questionAnswerRepo,
		RemindPasswordLogRepository:    remindPasswordLogRepo,
	}
}
