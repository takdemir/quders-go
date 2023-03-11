package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	emailservice "notificaiton/pkg/services/email"
	"notificaiton/pkg/structs"
	"notificaiton/pkg/utils/common"
	"regexp"
	"strings"
)

func (h *Handler) SendEmail(c echo.Context) error {
	jsonMap := structs.EmailFeatures{}

	err := json.NewDecoder(c.Request().Body).Decode(&jsonMap)
	if err != nil {
		response := common.ReplyUtil(false, "", fmt.Sprintf("email request body content parse error %s", err.Error()))
		return c.JSON(http.StatusOK, response)
	}

	ok, response := verifyEmailRequestContent(jsonMap)
	if !ok {
		return c.JSON(http.StatusBadRequest, response)
	}

	var data structs.EmailFeatures
	data = response.Data.(structs.EmailFeatures)

	emailServiceInstance := emailservice.EmailService{}
	emailService, activeEmailServiceError := emailServiceInstance.SetActiveEmailService()
	if activeEmailServiceError != nil {
		response := common.ReplyUtil(false, "", "no active email service")
		return c.JSON(http.StatusBadRequest, response)
	}

	response = emailService.ActiveEmailService.SendEmail(data)
	return c.JSON(http.StatusOK, response)
}

func verifyEmailRequestContent(jsonMap structs.EmailFeatures) (bool, *common.ResponseMessage) {

	if len(jsonMap.From) == 0 {
		jsonMap.From = map[string]string{
			"email": "taneryzb@hotmail.com",
			"name":  "Quders",
		}
	}

	if len(jsonMap.HtmlTemplateData) == 0 {
		jsonMap.HtmlTemplateData = map[string]string{
			"Title":     "Info Mail",
			"Name":      "Quders",
			"QudersURL": "https://www.quders.com/",
		}
	}

	if strings.TrimSpace(jsonMap.Subject) == "" {
		response := common.ReplyUtil(false, "", "subject is not exist or empty")
		return false, response
	}

	if strings.TrimSpace(jsonMap.PlainText) == "" && strings.TrimSpace(jsonMap.HtmlContent) == "" {
		response := common.ReplyUtil(false, "", "plaintext and html content can not be empty at the same time")
		return false, response
	}

	var tos []map[string]string
	tos = jsonMap.Tos

	if len(tos) == 0 {
		response := common.ReplyUtil(false, "", "tos can not be empty")
		return false, response
	}

	for _, v := range tos {
		for _, email := range v {
			_, err := regexp.MatchString("^\\w+([\\.-]?\\w+)*@\\w+([\\.-]?\\w+)*(\\.\\w{2,3})+$", email)
			if err != nil {
				return true, common.ReplyUtil(false, "", err.Error())
			}
		}

	}

	return true, common.ReplyUtil(true, jsonMap, "success")
}
