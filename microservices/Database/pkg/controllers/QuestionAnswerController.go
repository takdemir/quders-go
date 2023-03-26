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

func (h *Handler) GetQuestionAnswers(c echo.Context) error {
	questionAnswers, err := h.QuestionAnswerRepository.FetchQuestionAnswers()
	if err != nil {
		response := common.ReplyUtil(false, "", "error on fetch question answers: "+err.Error())
		return c.JSON(http.StatusBadRequest, response)
	}
	response := common.ReplyUtil(true, questionAnswers, "success")
	return c.JSON(http.StatusOK, response)
}

func (h *Handler) GetQuestionAnswerById(c echo.Context) error {
	questionAnswerIdParam := c.Param("id")
	if strings.TrimSpace(questionAnswerIdParam) == "" {
		response := common.ReplyUtil(false, "", "no valid question answer id")
		return c.JSON(http.StatusBadRequest, response)
	}
	questionAnswerId, _ := strconv.Atoi(questionAnswerIdParam)
	questionAnswerDetail, err := h.QuestionAnswerRepository.FetchQuestionAnswerById(uint(questionAnswerId))
	if err != nil {
		response := common.ReplyUtil(false, "", "error on fetch question answer: "+err.Error())
		return c.JSON(http.StatusBadRequest, response)
	}
	response := common.ReplyUtil(true, questionAnswerDetail, "success")
	return c.JSON(http.StatusOK, response)
}

func (h *Handler) CreateQuestionAnswer(c echo.Context) error {
	jsonMap := models.QuestionAnswer{}
	err := json.NewDecoder(c.Request().Body).Decode(&jsonMap)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	var user models.User
	var userCount int64
	h.DB.Table("user").Where("id=?", jsonMap.UserId).Find(&user).Count(&userCount)
	if userCount == 0 {
		response := common.ReplyUtil(false, "", "no user found")
		return c.JSON(http.StatusBadRequest, response)
	}
	jsonMap.User = user

	var registrationQuestion models.RegistrationQuestion
	var registrationQuestionCount int64
	h.DB.Table("registration_question").Where("id=?", jsonMap.QuestionId).Find(&registrationQuestion).Count(&registrationQuestionCount)
	if registrationQuestionCount == 0 {
		response := common.ReplyUtil(false, "", "no registration question found")
		return c.JSON(http.StatusBadRequest, response)
	}
	jsonMap.Question = registrationQuestion

	v := validator.New()
	err = v.Struct(jsonMap)
	if err != nil {
		for _, valError := range err.(validator.ValidationErrors) {
			response := common.ReplyUtil(false, "", "No valid field: "+fmt.Sprint(valError))
			return c.JSON(http.StatusBadRequest, response)
		}
	}
	fmt.Println(jsonMap)
	err = h.QuestionAnswerRepository.CreateQuestionAnswer(jsonMap)
	if err != nil {
		response := common.ReplyUtil(false, "", err.Error())
		return c.JSON(http.StatusBadRequest, response)
	}
	response := common.ReplyUtil(true, "", "success")
	return c.JSON(http.StatusOK, response)
}

func (h *Handler) UpdateQuestionAnswer(c echo.Context) error {
	jsonMap := models.QuestionAnswer{}
	err := json.NewDecoder(c.Request().Body).Decode(&jsonMap)
	if err != nil {
		response := common.ReplyUtil(false, "", "question answer json parse error"+err.Error())
		return c.JSON(http.StatusBadRequest, response)
	}

	var user models.User
	var userCount int64
	h.DB.Table("user").Where("id=?", jsonMap.UserId).Find(&user).Count(&userCount)
	if userCount == 0 {
		response := common.ReplyUtil(false, "", "no user found")
		return c.JSON(http.StatusBadRequest, response)
	}
	jsonMap.User = user

	var registrationQuestion models.RegistrationQuestion
	var registrationQuestionCount int64
	h.DB.Table("registration_question").Where("id=?", jsonMap.QuestionId).Find(&registrationQuestion).Count(&registrationQuestionCount)
	if registrationQuestionCount == 0 {
		response := common.ReplyUtil(false, "", "no registration question found")
		return c.JSON(http.StatusBadRequest, response)
	}
	jsonMap.Question = registrationQuestion

	v := validator.New()
	err = v.Struct(jsonMap)
	if err != nil {
		for _, valError := range err.(validator.ValidationErrors) {
			response := common.ReplyUtil(false, "", "no valid field on update: "+fmt.Sprint(valError))
			return c.JSON(http.StatusBadRequest, response)
		}
	}
	fmt.Println(jsonMap)
	QuestionAnswerIdParam := c.Param("id")
	if strings.TrimSpace(QuestionAnswerIdParam) == "" {
		response := common.ReplyUtil(false, "", "no valid question answer id")
		return c.JSON(http.StatusBadRequest, response)
	}
	QuestionAnswerId, _ := strconv.Atoi(QuestionAnswerIdParam)
	err = h.QuestionAnswerRepository.UpdateQuestionAnswer(uint(QuestionAnswerId), jsonMap)
	if err != nil {
		response := common.ReplyUtil(false, "", "update error: "+err.Error())
		return c.JSON(http.StatusBadRequest, response)
	}
	response := common.ReplyUtil(true, "", "success")
	return c.JSON(http.StatusOK, response)
}
