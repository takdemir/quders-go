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

func (h *Handler) GetFrameworks(c echo.Context) error {
	Frameworks, err := h.FrameworkRepository.FetchFrameworks()
	if err != nil {
		response := common.ReplyUtil(false, "", "error on fetch frameworks: "+err.Error())
		return c.JSON(http.StatusBadRequest, response)
	}
	response := common.ReplyUtil(true, Frameworks, "success")
	return c.JSON(http.StatusOK, response)
}

func (h *Handler) GetFrameworkById(c echo.Context) error {
	FrameworkIdParam := c.Param("id")
	if strings.TrimSpace(FrameworkIdParam) == "" {
		response := common.ReplyUtil(false, "", "no valid framework id")
		return c.JSON(http.StatusBadRequest, response)
	}
	FrameworkId, _ := strconv.Atoi(FrameworkIdParam)
	FrameworkDetail, err := h.FrameworkRepository.FetchFrameworkById(uint(FrameworkId))
	if err != nil {
		response := common.ReplyUtil(false, "", "error on fetch framework:"+err.Error())
		return c.JSON(http.StatusBadRequest, response)
	}
	response := common.ReplyUtil(true, FrameworkDetail, "success")
	return c.JSON(http.StatusOK, response)
}

func (h *Handler) CreateFramework(c echo.Context) error {
	jsonMap := models.Framework{}
	err := json.NewDecoder(c.Request().Body).Decode(&jsonMap)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	var programingLanguage models.ProgramingLanguage
	var programingLanguageCount int64

	h.DB.Model(&programingLanguage).Where("id=?", jsonMap.ProgramingLanguageId).Find(&programingLanguage).Count(&programingLanguageCount)
	if programingLanguageCount <= 0 {
		response := common.ReplyUtil(false, "", "programing language not found")
		return c.JSON(http.StatusBadRequest, response)
	}

	jsonMap.ProgramingLanguage = programingLanguage
	jsonMap.Name = strings.ToUpper(jsonMap.Name)

	v := validator.New()
	err = v.Struct(jsonMap)
	if err != nil {
		for _, valError := range err.(validator.ValidationErrors) {
			response := common.ReplyUtil(false, "", "No valid field "+fmt.Sprint(valError))
			return c.JSON(http.StatusBadRequest, response)
		}
	}

	err = h.FrameworkRepository.CreateFramework(jsonMap)
	if err != nil {
		response := common.ReplyUtil(false, "", err.Error())
		return c.JSON(http.StatusBadRequest, response)
	}
	response := common.ReplyUtil(true, "", "success")
	return c.JSON(http.StatusOK, response)
}

func (h *Handler) UpdateFramework(c echo.Context) error {
	jsonMap := models.Framework{}
	err := json.NewDecoder(c.Request().Body).Decode(&jsonMap)
	if err != nil {
		response := common.ReplyUtil(false, "", "framework json parse error"+err.Error())
		return c.JSON(http.StatusBadRequest, response)
	}
	var programingLanguage models.ProgramingLanguage
	var programingLanguageCount int64

	h.DB.Model(&programingLanguage).Where("id=?", jsonMap.ProgramingLanguageId).Find(&programingLanguage).Count(&programingLanguageCount)
	if programingLanguageCount <= 0 {
		response := common.ReplyUtil(false, "", "programing language not found")
		return c.JSON(http.StatusBadRequest, response)
	}

	jsonMap.ProgramingLanguage = programingLanguage
	jsonMap.Name = strings.ToUpper(jsonMap.Name)

	v := validator.New()
	err = v.Struct(jsonMap)
	if err != nil {
		for _, valError := range err.(validator.ValidationErrors) {
			response := common.ReplyUtil(false, "", "No valid field "+fmt.Sprint(valError))
			return c.JSON(http.StatusBadRequest, response)
		}
	}
	FrameworkIdParam := c.Param("id")
	if strings.TrimSpace(FrameworkIdParam) == "" {
		response := common.ReplyUtil(false, "", "no valid framework id")
		return c.JSON(http.StatusBadRequest, response)
	}
	FrameworkId, _ := strconv.Atoi(FrameworkIdParam)
	err = h.FrameworkRepository.UpdateFramework(uint(FrameworkId), jsonMap)
	if err != nil {
		response := common.ReplyUtil(false, "", "update error: "+err.Error())
		return c.JSON(http.StatusBadRequest, response)
	}
	response := common.ReplyUtil(true, "", "success")
	return c.JSON(http.StatusOK, response)
}
