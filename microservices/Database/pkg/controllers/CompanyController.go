package controllers

import (
	"database/pkg/utils/common"
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

func validate() {

}

func (h *Handler) GetCompanies(c echo.Context) error {
	response := common.ReplyUtil(true, "", "success")
	return c.JSON(http.StatusOK, response)
}

func (h *Handler) GetCompanyById(c echo.Context) error {
	response := common.ReplyUtil(true, "", "success")
	return c.JSON(http.StatusOK, response)
}

func (h *Handler) CreateCompany(c echo.Context) error {
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
	response := common.ReplyUtil(true, "", "success")
	return c.JSON(http.StatusOK, response)
}

func (h *Handler) UpdateCompany(c echo.Context) error {
	response := common.ReplyUtil(true, "", "success")
	return c.JSON(http.StatusOK, response)
}
