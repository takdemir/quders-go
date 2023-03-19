package controllers

import (
	"database/pkg/models"
	"database/pkg/utils/common"
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"gopkg.in/go-playground/validator.v9"
	"net/http"
	"strconv"
	"strings"
)

func (h *Handler) GetCountries(c echo.Context) error {
	companies, err := h.CountryRepository.FetchCountries()
	if err != nil {
		response := common.ReplyUtil(false, "", "error on fetch countries: "+err.Error())
		return c.JSON(http.StatusBadRequest, response)
	}
	response := common.ReplyUtil(true, companies, "success")
	return c.JSON(http.StatusOK, response)
}

func (h *Handler) GetCountryById(c echo.Context) error {
	countryIdParam := c.Param("id")
	if strings.TrimSpace(countryIdParam) == "" {
		response := common.ReplyUtil(false, "", "no valid country id")
		return c.JSON(http.StatusBadRequest, response)
	}
	countryId, _ := strconv.Atoi(countryIdParam)
	countryDetail, err := h.CountryRepository.FetchCountryById(countryId)
	if err != nil {
		response := common.ReplyUtil(false, "", "error on fetch country:"+err.Error())
		return c.JSON(http.StatusBadRequest, response)
	}
	response := common.ReplyUtil(true, countryDetail, "success")
	return c.JSON(http.StatusOK, response)
}

func (h *Handler) CreateCountry(c echo.Context) error {
	/*var currency models.Currency
	if err := c.Bind(&currency); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	fmt.Println(currency)*/
	jsonMap := models.Country{}
	checkCountry := models.Country{}
	var checkCountryCount int64
	err := json.NewDecoder(c.Request().Body).Decode(&jsonMap)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	h.DB.Table("country").Where("code=?", jsonMap.Code).Find(&checkCountry).Count(&checkCountryCount)
	if checkCountryCount > 0 {
		response := common.ReplyUtil(false, "", "country code is already exist")
		return c.JSON(http.StatusBadRequest, response)
	}

	var currency models.Currency
	h.DB.Table("currency").Where("id=?", jsonMap.CurrencyId).Find(&currency)

	if err != nil {
		return err
	}
	if currency.ID == 0 {
		response := common.ReplyUtil(false, "", "invalid currency")
		return c.JSON(http.StatusBadRequest, response)
	}
	jsonMap.Currency = currency
	fmt.Println(jsonMap)
	v := validator.New()
	err = v.Struct(jsonMap)
	if err != nil {
		for _, valError := range err.(validator.ValidationErrors) {
			response := common.ReplyUtil(false, "", "No valid field "+fmt.Sprint(valError))
			return c.JSON(http.StatusBadRequest, response)
		}
	}
	err = h.CountryRepository.CreateCountry(jsonMap)
	if err != nil {
		response := common.ReplyUtil(false, "", err.Error())
		return c.JSON(http.StatusBadRequest, response)
	}
	response := common.ReplyUtil(true, "", "success")
	return c.JSON(http.StatusOK, response)
}

func (h *Handler) UpdateCountry(c echo.Context) error {
	jsonMap := models.Country{}
	err := json.NewDecoder(c.Request().Body).Decode(&jsonMap)
	if err != nil {
		response := common.ReplyUtil(false, "", "country json parse error"+err.Error())
		return c.JSON(http.StatusBadRequest, response)
	}

	countryIdParam := c.Param("id")
	if strings.TrimSpace(countryIdParam) == "" {
		response := common.ReplyUtil(false, "", "no valid country id")
		return c.JSON(http.StatusBadRequest, response)
	}
	countryId, _ := strconv.Atoi(countryIdParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	var currency models.Currency
	h.DB.Table("currency").Where("id=?", jsonMap.CurrencyId).Find(&currency)
	if currency.ID == 0 {
		response := common.ReplyUtil(false, "", "invalid currency")
		return c.JSON(http.StatusBadRequest, response)
	}
	jsonMap.Currency = currency

	v := validator.New()
	err = v.Struct(jsonMap)
	if err != nil {
		for _, valError := range err.(validator.ValidationErrors) {
			response := common.ReplyUtil(false, "", "no valid field "+fmt.Sprint(valError))
			return c.JSON(http.StatusBadRequest, response)
		}
	}
	var checkCountry models.Country
	h.DB.Table("country").Where("code=?", jsonMap.Code).Find(&checkCountry)
	if checkCountry.ID != uint(countryId) {
		response := common.ReplyUtil(false, "", "country code is already exist")
		return c.JSON(http.StatusBadRequest, response)
	}

	err = h.CountryRepository.UpdateCountry(countryId, jsonMap)
	if err != nil {
		response := common.ReplyUtil(false, "", "update error: "+err.Error())
		return c.JSON(http.StatusBadRequest, response)
	}
	response := common.ReplyUtil(true, "", "success")
	return c.JSON(http.StatusOK, response)
}
