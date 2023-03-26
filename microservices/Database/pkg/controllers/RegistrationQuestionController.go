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

func (h *Handler) GetRegistrationQuestions(c echo.Context) error {
	RegistrationQuestions, err := h.RegistrationQuestionRepository.FetchRegistrationQuestions()
	if err != nil {
		response := common.ReplyUtil(false, "", "error on fetch registration questions: "+err.Error())
		return c.JSON(http.StatusBadRequest, response)
	}
	response := common.ReplyUtil(true, RegistrationQuestions, "success")
	return c.JSON(http.StatusOK, response)
}

func (h *Handler) GetRegistrationQuestionById(c echo.Context) error {
	registrationQuestionIdParam := c.Param("id")
	if strings.TrimSpace(registrationQuestionIdParam) == "" {
		response := common.ReplyUtil(false, "", "no valid registration question id")
		return c.JSON(http.StatusBadRequest, response)
	}
	RegistrationQuestionId, _ := strconv.Atoi(registrationQuestionIdParam)
	RegistrationQuestionDetail, err := h.RegistrationQuestionRepository.FetchRegistrationQuestionById(uint(RegistrationQuestionId))
	if err != nil {
		response := common.ReplyUtil(false, "", "error on fetch registration question: "+err.Error())
		return c.JSON(http.StatusBadRequest, response)
	}
	response := common.ReplyUtil(true, RegistrationQuestionDetail, "success")
	return c.JSON(http.StatusOK, response)
}

func (h *Handler) CreateRegistrationQuestion(c echo.Context) error {
	jsonMap := models.RegistrationQuestion{}
	err := json.NewDecoder(c.Request().Body).Decode(&jsonMap)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	//fmt.Println(jsonMap)
	v := validator.New()
	err = v.Struct(jsonMap)
	if err != nil {
		for _, valError := range err.(validator.ValidationErrors) {
			response := common.ReplyUtil(false, "", "No valid field: "+fmt.Sprint(valError))
			return c.JSON(http.StatusBadRequest, response)
		}
	}

	err = h.RegistrationQuestionRepository.CreateRegistrationQuestion(jsonMap)
	if err != nil {
		response := common.ReplyUtil(false, "", err.Error())
		return c.JSON(http.StatusBadRequest, response)
	}
	response := common.ReplyUtil(true, "", "success")
	return c.JSON(http.StatusOK, response)
}

func (h *Handler) UpdateRegistrationQuestion(c echo.Context) error {
	jsonMap := models.RegistrationQuestion{}
	err := json.NewDecoder(c.Request().Body).Decode(&jsonMap)
	if err != nil {
		response := common.ReplyUtil(false, "", "registration question json parse error"+err.Error())
		return c.JSON(http.StatusBadRequest, response)
	}
	v := validator.New()
	err = v.Struct(jsonMap)
	if err != nil {
		for _, valError := range err.(validator.ValidationErrors) {
			response := common.ReplyUtil(false, "", "no valid field on update: "+fmt.Sprint(valError))
			return c.JSON(http.StatusBadRequest, response)
		}
	}
	registrationQuestionIdParam := c.Param("id")
	if strings.TrimSpace(registrationQuestionIdParam) == "" {
		response := common.ReplyUtil(false, "", "no valid registration question id")
		return c.JSON(http.StatusBadRequest, response)
	}
	registrationQuestionId, _ := strconv.Atoi(registrationQuestionIdParam)
	err = h.RegistrationQuestionRepository.UpdateRegistrationQuestion(uint(registrationQuestionId), jsonMap)
	if err != nil {
		response := common.ReplyUtil(false, "", "update error: "+err.Error())
		return c.JSON(http.StatusBadRequest, response)
	}
	response := common.ReplyUtil(true, "", "success")
	return c.JSON(http.StatusOK, response)
}
