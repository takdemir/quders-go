package controllers

import (
	"database/pkg/models"
	"database/pkg/utils/common"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"gopkg.in/go-playground/validator.v9"
	"net/http"
	"time"
)

type Currency struct {
	Name      string    `json:"name"`
	Code      string    `json:"code"`
	IsActive  bool      `json:"isActive"`
	CreatedAt time.Time `json:"createdAt"`
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
	var newCurrency = models.Currency{}
	err := json.NewDecoder(c.Request().Body).Decode(&newCurrency)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	v := validator.New()
	validateError := v.Struct(newCurrency)
	if validateError != nil {
		for _, v := range validateError.(validator.ValidationErrors) {
			response := common.ReplyUtil(false, "", "invalid field "+v.Field())
			return c.JSON(http.StatusBadRequest, response)
		}
	}
	err = h.CurrencyRepository.CreateCurrency(newCurrency)
	if err != nil {
		response := common.ReplyUtil(false, "", err.Error())
		return c.JSON(http.StatusBadRequest, response)
	}

	response := common.ReplyUtil(true, "", "success")
	return c.JSON(http.StatusOK, response)
}

func (h *Handler) UpdateCurrency(c echo.Context) error {
	response := common.ReplyUtil(true, "", "success")
	return c.JSON(http.StatusOK, response)
}
