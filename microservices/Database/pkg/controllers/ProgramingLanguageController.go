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

func (h *Handler) GetProgramingLanguages(c echo.Context) error {
	programingLanguages, err := h.ProgramingLanguageRepository.FetchProgramingLanguages()
	if err != nil {
		response := common.ReplyUtil(false, "", "error on fetch programing languages: "+err.Error())
		return c.JSON(http.StatusBadRequest, response)
	}
	response := common.ReplyUtil(true, programingLanguages, "success")
	return c.JSON(http.StatusOK, response)
}

func (h *Handler) GetProgramingLanguageById(c echo.Context) error {
	programingLanguageIdParam := c.Param("id")
	if strings.TrimSpace(programingLanguageIdParam) == "" {
		response := common.ReplyUtil(false, "", "no valid programing language id")
		return c.JSON(http.StatusBadRequest, response)
	}
	programingLanguageId, _ := strconv.Atoi(programingLanguageIdParam)
	programingLanguageDetail, err := h.ProgramingLanguageRepository.FetchProgramingLanguageById(uint(programingLanguageId))
	if err != nil {
		response := common.ReplyUtil(false, "", "error on fetch programing language:"+err.Error())
		return c.JSON(http.StatusBadRequest, response)
	}
	response := common.ReplyUtil(true, programingLanguageDetail, "success")
	return c.JSON(http.StatusOK, response)
}

func (h *Handler) CreateProgramingLanguage(c echo.Context) error {
	jsonMap := models.ProgramingLanguage{}
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
	jsonMap.Name = strings.ToUpper(jsonMap.Name)
	jsonMap.Icon = strings.ToLower(jsonMap.Icon)
	err = h.ProgramingLanguageRepository.CreateProgramingLanguage(jsonMap)
	if err != nil {
		response := common.ReplyUtil(false, "", err.Error())
		return c.JSON(http.StatusBadRequest, response)
	}
	response := common.ReplyUtil(true, "", "success")
	return c.JSON(http.StatusOK, response)
}

func (h *Handler) UpdateProgramingLanguage(c echo.Context) error {
	jsonMap := models.ProgramingLanguage{}
	err := json.NewDecoder(c.Request().Body).Decode(&jsonMap)
	if err != nil {
		response := common.ReplyUtil(false, "", "programing language json parse error"+err.Error())
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
	programingLanguageIdParam := c.Param("id")
	if strings.TrimSpace(programingLanguageIdParam) == "" {
		response := common.ReplyUtil(false, "", "no valid programing language id")
		return c.JSON(http.StatusBadRequest, response)
	}
	programingLanguageId, _ := strconv.Atoi(programingLanguageIdParam)
	err = h.ProgramingLanguageRepository.UpdateProgramingLanguage(uint(programingLanguageId), jsonMap)
	if err != nil {
		response := common.ReplyUtil(false, "", "update error: "+err.Error())
		return c.JSON(http.StatusBadRequest, response)
	}
	response := common.ReplyUtil(true, "", "success")
	return c.JSON(http.StatusOK, response)
}
