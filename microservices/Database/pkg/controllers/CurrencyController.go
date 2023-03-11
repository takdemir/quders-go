package controllers

import (
	"database/pkg/utils/common"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"gopkg.in/validator.v2"
	"net/http"
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

	if !common.RegexpString("^[a-zA-Z]{2,10}$", v.(string)) {
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

func (h *Handler) GetCurrencies(c echo.Context) error {
	currencies := h.CurrencyRepository.GetCurrencies()
	response := common.ReplyUtil(true, currencies, "success")
	return c.JSON(http.StatusOK, response)
}

func (h *Handler) GetCurrencyById(c echo.Context) error {
	response := common.ReplyUtil(true, "", "success")
	return c.JSON(http.StatusOK, response)
}

func (h *Handler) CreateCurrency(c echo.Context) error {
	var newCurrency = new(Currency)
	err := json.NewDecoder(c.Request().Body).Decode(&newCurrency)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	errValidation := validator.Validate(newCurrency)
	fmt.Println(errValidation)
	if errValidation != nil {
		response := common.ReplyUtil(false, errValidation.Error(), "success")
		return c.JSON(http.StatusOK, response)
	}
	fmt.Println(newCurrency)
	response := common.ReplyUtil(true, "", "success")
	return c.JSON(http.StatusOK, response)
}

func (h *Handler) UpdateCurrency(c echo.Context) error {
	response := common.ReplyUtil(true, "", "success")
	return c.JSON(http.StatusOK, response)
}
