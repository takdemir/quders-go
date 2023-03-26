package controllers

import (
	"database/pkg/models"
	"database/pkg/utils/common"
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"gopkg.in/go-playground/validator.v9"
	"net/http"
)

func (h *Handler) CheckOneTimeCode(c echo.Context) error {
	jsonMap := models.RemindPasswordLog{}
	err := json.NewDecoder(c.Request().Body).Decode(&jsonMap)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	oneTimeCode := h.RemindPasswordLogRepository.CheckOneTimeCode(jsonMap.Email, jsonMap.OneTimeCode)
	if !oneTimeCode {
		response := common.ReplyUtil(false, "", "error on fetch check one time code: "+err.Error())
		return c.JSON(http.StatusNotFound, response)
	}
	response := common.ReplyUtil(true, "", "success")
	return c.JSON(http.StatusOK, response)
}

func (h *Handler) CreateOneTimeCode(c echo.Context) error {
	jsonMap := models.RemindPasswordLog{}
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

	v := validator.New()
	err = v.Struct(jsonMap)
	if err != nil {
		for _, valError := range err.(validator.ValidationErrors) {
			response := common.ReplyUtil(false, "", "No valid field: "+fmt.Sprint(valError))
			return c.JSON(http.StatusBadRequest, response)
		}
	}
	fmt.Println(jsonMap)
	err = h.RemindPasswordLogRepository.CreateOneTimeCode(jsonMap)
	if err != nil {
		response := common.ReplyUtil(false, "", err.Error())
		return c.JSON(http.StatusBadRequest, response)
	}
	response := common.ReplyUtil(true, "", "success")
	return c.JSON(http.StatusOK, response)
}

func (h *Handler) UpdateOneTimeCode(c echo.Context) error {
	jsonMap := models.RemindPasswordLog{}
	err := json.NewDecoder(c.Request().Body).Decode(&jsonMap)
	if err != nil {
		response := common.ReplyUtil(false, "", "remind password json parse error"+err.Error())
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

	v := validator.New()
	err = v.Struct(jsonMap)
	if err != nil {
		for _, valError := range err.(validator.ValidationErrors) {
			response := common.ReplyUtil(false, "", "no valid field on update: "+fmt.Sprint(valError))
			return c.JSON(http.StatusBadRequest, response)
		}
	}
	fmt.Println(jsonMap)
	/*	QuestionAnswerIdParam := c.Param("id")
		if strings.TrimSpace(QuestionAnswerIdParam) == "" {
			response := common.ReplyUtil(false, "", "no valid question answer id")
			return c.JSON(http.StatusBadRequest, response)
		}
		QuestionAnswerId, _ := strconv.Atoi(QuestionAnswerIdParam)*/
	err = h.RemindPasswordLogRepository.UpdateOneTimeCode(jsonMap)
	if err != nil {
		response := common.ReplyUtil(false, "", "update error: "+err.Error())
		return c.JSON(http.StatusBadRequest, response)
	}
	response := common.ReplyUtil(true, "", "success")
	return c.JSON(http.StatusOK, response)
}
