package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"net"
	"net/http"
	"net/url"
	sms_service "notificaiton/pkg/services/sms"
	"notificaiton/pkg/utils/common"
	"os"
	"strings"
)

func (h *Handler) SendSms(c echo.Context) error {

	smsServiceInstance := sms_service.SmsService{}
	smsService, activeSmsServiceError := smsServiceInstance.SetActiveSmsService()
	if activeSmsServiceError != nil {
		response := common.ReplyUtil(false, "", "no active sms service")
		return c.JSON(http.StatusBadRequest, response)
	}
	jsonMap := smsServiceInstance
	err := json.NewDecoder(c.Request().Body).Decode(&jsonMap)
	if err != nil {
		response := common.ReplyUtil(false, "", "sms request parse error"+err.Error())
		return c.JSON(http.StatusOK, response)
	}
	if strings.Trim(jsonMap.Text, " ") == "" {
		response := common.ReplyUtil(false, "", "text can not be empty")
		return c.JSON(http.StatusBadRequest, response)
	}
	if len(jsonMap.Phones) == 0 {
		response := common.ReplyUtil(false, "", "phones can not be empty")
		return c.JSON(http.StatusBadRequest, response)
	}

	u, err := url.Parse(os.Getenv("SMS_PACIFIC_URL"))
	if err != nil {
		response := common.ReplyUtil(false, "", "sms url parse error "+err.Error())
		return c.JSON(http.StatusOK, response)
	}
	scheme := u.Scheme
	username := u.User.Username()
	password, _ := u.User.Password()
	host, port, _ := net.SplitHostPort(u.Host)
	if port == "80" {
		port = ""
	} else {
		port = ":" + port
	}
	path := u.Path
	jsonMap.Url = fmt.Sprintf("%s://%s%s%s", scheme, host, port, path)
	jsonMap.Username = username
	jsonMap.Password = password
	query, queryError := url.ParseQuery(u.RawQuery)

	_, isExist := query["from"]
	if queryError != nil || !isExist || strings.TrimSpace(query["from"][0]) == "" {
		response := common.ReplyUtil(false, "", "from is not valid")
		return c.JSON(http.StatusBadRequest, response)
	}
	jsonMap.From = query["from"][0]

	response := smsService.ActiveSmsService.SendSms(jsonMap)
	return c.JSON(http.StatusOK, response)
}
