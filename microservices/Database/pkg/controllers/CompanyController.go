package controllers

import (
	"database/pkg/models"
	"database/pkg/utils/common"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"gopkg.in/go-playground/validator.v9"
	"net/http"
	"strconv"
	"strings"
)

func (h *Handler) GetCompanies(c echo.Context) error {
	companies, err := h.CompanyRepository.FetchCompanies()
	if err != nil {
		response := common.ReplyUtil(false, "", "error on fetch companies: "+err.Error())
		return c.JSON(http.StatusBadRequest, response)
	}
	response := common.ReplyUtil(true, companies, "success")
	return c.JSON(http.StatusOK, response)
}

func (h *Handler) GetCompanyById(c echo.Context) error {
	companyIdParam := c.Param("id")
	if strings.TrimSpace(companyIdParam) == "" {
		response := common.ReplyUtil(false, "", "no valid company id")
		return c.JSON(http.StatusBadRequest, response)
	}
	companyId, _ := strconv.Atoi(companyIdParam)
	companyDetail, err := h.CompanyRepository.GetCompanyById(companyId)
	if err != nil {
		response := common.ReplyUtil(false, "", "error on fetch company:"+err.Error())
		return c.JSON(http.StatusBadRequest, response)
	}
	response := common.ReplyUtil(true, companyDetail, "success")
	return c.JSON(http.StatusOK, response)
}

func (h *Handler) CreateCompany(c echo.Context) error {
	/*var currency models.Currency
	if err := c.Bind(&currency); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	fmt.Println(currency)*/
	jsonMap := models.Company{}
	err := json.NewDecoder(c.Request().Body).Decode(&jsonMap)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	//fmt.Println(jsonMap)
	v := validator.New()
	err = v.Struct(jsonMap)
	if err != nil {
		for _, valError := range err.(validator.ValidationErrors) {
			response := common.ReplyUtil(false, "", "No valid field "+valError.Field())
			return c.JSON(http.StatusBadRequest, response)
		}
	}
	err = h.CompanyRepository.CreateCompany(jsonMap)
	if err != nil {
		response := common.ReplyUtil(false, "", err.Error())
		return c.JSON(http.StatusBadRequest, response)
	}
	response := common.ReplyUtil(true, "", "success")
	return c.JSON(http.StatusOK, response)
}

func (h *Handler) UpdateCompany(c echo.Context) error {
	jsonMap := models.Company{}
	err := json.NewDecoder(c.Request().Body).Decode(&jsonMap)
	if err != nil {
		response := common.ReplyUtil(false, "", "company json parse error"+err.Error())
		return c.JSON(http.StatusBadRequest, response)
	}
	v := validator.New()
	err = v.Struct(jsonMap)
	if err != nil {
		for _, valError := range err.(validator.ValidationErrors) {
			response := common.ReplyUtil(false, "", "no valid field "+valError.Field())
			return c.JSON(http.StatusBadRequest, response)
		}
	}
	companyIdParam := c.Param("id")
	if strings.TrimSpace(companyIdParam) == "" {
		response := common.ReplyUtil(false, "", "no valid company id")
		return c.JSON(http.StatusBadRequest, response)
	}
	companyId, _ := strconv.Atoi(companyIdParam)
	err = h.CompanyRepository.UpdateCompany(companyId, jsonMap)
	if err != nil {
		response := common.ReplyUtil(false, "", "update error: "+err.Error())
		return c.JSON(http.StatusBadRequest, response)
	}
	response := common.ReplyUtil(true, "", "success")
	return c.JSON(http.StatusOK, response)
}
