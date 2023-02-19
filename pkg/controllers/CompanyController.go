package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"quders/pkg/utils"
)

func validate() {

}

func GetCompanies(c echo.Context) error {
	response := utils.ReplyUtil(true, "", "success")
	return c.JSON(http.StatusOK, response)
}

func GetCompanyById(c echo.Context) error {
	response := utils.ReplyUtil(true, "", "success")
	return c.JSON(http.StatusOK, response)
}

func CreateCompany(c echo.Context) error {
	/*var currency models.Currency
	if err := c.Bind(&currency); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	fmt.Println(currency)*/
	jsonMap := make(map[string]interface{})
	err := json.NewDecoder(c.Request().Body).Decode(&jsonMap)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	fmt.Println(jsonMap)
	response := utils.ReplyUtil(true, "", "success")
	return c.JSON(http.StatusOK, response)
}

func UpdateCompany(c echo.Context) error {
	response := utils.ReplyUtil(true, "", "success")
	return c.JSON(http.StatusOK, response)
}
