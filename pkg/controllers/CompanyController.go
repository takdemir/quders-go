package controllers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"quders/pkg/utils"
)

func GetCompanies(c echo.Context) error {
	response := utils.ReplyUtil(true, "", "success")
	return c.JSON(http.StatusOK, response)
}

func GetCompanyById(c echo.Context) error {
	response := utils.ReplyUtil(true, "", "success")
	return c.JSON(http.StatusOK, response)
}

func Create(c echo.Context) error {
	response := utils.ReplyUtil(true, "", "success")
	return c.JSON(http.StatusOK, response)
}

func Update(c echo.Context) error {
	response := utils.ReplyUtil(true, "", "success")
	return c.JSON(http.StatusOK, response)
}
