package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"gopkg.in/validator.v2"
	"net/http"
	"quders/pkg/utils"
	"reflect"
	"time"
)

type Currency struct {
	Name      string    `json:"name" validate:"validateName"`
	Code      string    `json:"code"`
	IsActive  bool      `json:"isActive"`
	CreatedAt time.Time `json:"createdAt"`
}

func ValidateName(v interface{}, param string) error {
	st := reflect.ValueOf(v)
	if st.Kind() != reflect.String {
		return errors.New("parameter type must be string")
	}

	if !utils.RegexpString("^[a-zA-Z]{2,10}$", v.(string)) {
		return errors.New("name is invalid. must be at least 2 only letter chars")
	}
	return nil
}

func init() {
	err := validator.SetValidationFunc("validateName", ValidateName)
	if err != nil {
		return
	}
}

func GetCurrencies(c echo.Context) error {
	response := utils.ReplyUtil(true, "", "success")
	return c.JSON(http.StatusOK, response)
}

func GetCurrencyById(c echo.Context) error {
	response := utils.ReplyUtil(true, "", "success")
	return c.JSON(http.StatusOK, response)
}

func CreateCurrency(c echo.Context) error {
	var newCurrency = new(Currency)
	err := json.NewDecoder(c.Request().Body).Decode(&newCurrency)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	errValidation := validator.Validate(newCurrency)
	fmt.Println(errValidation)
	if errValidation != nil {
		response := utils.ReplyUtil(false, errValidation.Error(), "success")
		return c.JSON(http.StatusOK, response)
	}
	fmt.Println(newCurrency)
	response := utils.ReplyUtil(true, "", "success")
	return c.JSON(http.StatusOK, response)
}

func UpdateCurrency(c echo.Context) error {
	response := utils.ReplyUtil(true, "", "success")
	return c.JSON(http.StatusOK, response)
}
