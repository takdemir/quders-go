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

func (h *Handler) FetchCompanyDetailByCompanyId(c echo.Context) error {
	companyIdParam := c.Param("id")
	if strings.TrimSpace(companyIdParam) == "" {
		response := common.ReplyUtil(false, "", "no valid company id")
		return c.JSON(http.StatusBadRequest, response)
	}
	companyId, _ := strconv.Atoi(companyIdParam)
	companyDetail, err := h.CompanyDetailRepository.FetchCompanyDetailByCompanyId(uint(companyId))
	if err != nil {
		response := common.ReplyUtil(false, "", "error on fetch company detail:"+err.Error())
		return c.JSON(http.StatusBadRequest, response)
	}
	if companyDetail.ID == 0 {
		response := common.ReplyUtil(false, "", "no company detail with that company ID")
		return c.JSON(http.StatusNotFound, response)
	}
	response := common.ReplyUtil(true, companyDetail, "success")
	return c.JSON(http.StatusOK, response)
}

func (h *Handler) CreateCompanyDetail(c echo.Context) error {
	jsonMap := models.CompanyDetail{}
	err := json.NewDecoder(c.Request().Body).Decode(&jsonMap)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	var company models.Company
	var companyCount int64
	var country models.Country
	var countryCount int64
	var companyDetail models.CompanyDetail
	var companyDetailCount int64

	h.DB.Model(&companyDetail).Where("company_id=?", jsonMap.CompanyId).Find(&companyDetail).Count(&companyDetailCount)
	if companyDetailCount > 0 {
		response := common.ReplyUtil(false, "", "company has already detail")
		return c.JSON(http.StatusBadRequest, response)
	}

	h.DB.Model(&company).Where("id=?", jsonMap.CompanyId).Find(&company).Count(&companyCount)
	if companyCount == 0 {
		response := common.ReplyUtil(false, "", "no company found with that ID")
		return c.JSON(http.StatusBadRequest, response)
	}

	h.DB.Model(&country).Where("id=?", jsonMap.CountryId).Preload("Currency").Find(&country).Count(&countryCount)
	if countryCount == 0 {
		response := common.ReplyUtil(false, "", "no country found with that ID")
		return c.JSON(http.StatusBadRequest, response)
	}
	jsonMap.Company = company
	jsonMap.Country = country
	//fmt.Println(jsonMap)
	v := validator.New()
	err = v.Struct(jsonMap)
	if err != nil {
		for _, valError := range err.(validator.ValidationErrors) {
			response := common.ReplyUtil(false, "", "No valid field "+fmt.Sprint(valError))
			return c.JSON(http.StatusBadRequest, response)
		}
	}
	err = h.CompanyDetailRepository.CreateCompanyDetail(jsonMap)
	if err != nil {
		response := common.ReplyUtil(false, "", err.Error())
		return c.JSON(http.StatusBadRequest, response)
	}
	response := common.ReplyUtil(true, "", "success")
	return c.JSON(http.StatusOK, response)
}

func (h *Handler) UpdateCompanyDetail(c echo.Context) error {
	jsonMap := models.CompanyDetail{}
	err := json.NewDecoder(c.Request().Body).Decode(&jsonMap)
	if err != nil {
		response := common.ReplyUtil(false, "", "company json parse error"+err.Error())
		return c.JSON(http.StatusBadRequest, response)
	}

	companyIdParam := c.Param("id")
	if strings.TrimSpace(companyIdParam) == "" {
		response := common.ReplyUtil(false, "", "no valid company id")
		return c.JSON(http.StatusBadRequest, response)
	}
	companyId, _ := strconv.Atoi(companyIdParam)

	var company models.Company
	var companyCount int64
	var country models.Country
	var countryCount int64
	var companyDetail models.CompanyDetail
	var companyDetailCount int64

	h.DB.Model(&companyDetail).Where("company_id=?", jsonMap.CompanyId).Find(&companyDetail).Count(&companyDetailCount)
	if companyDetailCount > 0 && companyDetail.ID != uint(companyId) {
		response := common.ReplyUtil(false, "", "company has already detail")
		return c.JSON(http.StatusBadRequest, response)
	}

	h.DB.Model(&company).Where("id=?", jsonMap.CompanyId).Find(&company).Count(&companyCount)
	if companyCount == 0 {
		response := common.ReplyUtil(false, "", "no company found with that ID")
		return c.JSON(http.StatusBadRequest, response)
	}

	h.DB.Model(&country).Where("id=?", jsonMap.CountryId).Preload("Currency").Find(&country).Count(&countryCount)
	if countryCount == 0 {
		response := common.ReplyUtil(false, "", "no country found with that ID")
		return c.JSON(http.StatusBadRequest, response)
	}
	jsonMap.Company = company
	jsonMap.Country = country

	v := validator.New()
	err = v.Struct(jsonMap)
	if err != nil {
		for _, valError := range err.(validator.ValidationErrors) {
			response := common.ReplyUtil(false, "", "no valid field "+valError.Field())
			return c.JSON(http.StatusBadRequest, response)
		}
	}

	err = h.CompanyDetailRepository.UpdateCompanyDetail(uint(companyId), jsonMap)
	if err != nil {
		response := common.ReplyUtil(false, "", "update error: "+err.Error())
		return c.JSON(http.StatusBadRequest, response)
	}
	response := common.ReplyUtil(true, "", "success")
	return c.JSON(http.StatusOK, response)
}
