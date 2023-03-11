package controllers

import (
	"bytes"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"net/http"
	"notificaiton/pkg/structs"
	"notificaiton/pkg/utils/common"
	"os"
	"strings"
)

func (h *Handler) SendTeamsMessage(c echo.Context) error {

	jsonMap := structs.TeamsFeatures{}
	err := json.NewDecoder(c.Request().Body).Decode(&jsonMap)
	if err != nil {
		response := common.ReplyUtil(false, "", "teams payload parse error"+err.Error())
		return c.JSON(http.StatusOK, response)
	}

	webhookUrl := os.Getenv("TEAMS_WEBHOOK")

	if strings.TrimSpace(jsonMap.Title) == "" {
		response := common.ReplyUtil(false, "", "title can not be empty")
		return c.JSON(http.StatusBadRequest, response)
	}
	if strings.TrimSpace(jsonMap.Text) == "" {
		response := common.ReplyUtil(false, "", "text content can not be empty")
		return c.JSON(http.StatusBadRequest, response)
	}
	if strings.TrimSpace(jsonMap.ThemeColor) == "" {
		jsonMap.ThemeColor = "#DF813D"
	}

	requestBody := []byte(`{
		"title": "` + jsonMap.Title + `",
		"text": "` + jsonMap.Text + `",
		"themeColor":"` + jsonMap.ThemeColor + `"
	}`)

	client := &http.Client{}
	req, err := http.NewRequest("POST", webhookUrl, bytes.NewBuffer(requestBody))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	responseFromApi, err := client.Do(req)
	if err != nil {
		response := common.ReplyUtil(false, "", "send teams error-1: "+err.Error())
		return c.JSON(http.StatusBadRequest, response)
	}
	var res interface{}
	errFromJsonBody := json.NewDecoder(responseFromApi.Body).Decode(&res)
	if responseFromApi.StatusCode > 300 {
		detail := "no detail"
		resp := common.ReplyUtil(false, "", "send teams error-2: "+detail)
		return c.JSON(http.StatusBadRequest, resp)
	}

	if errFromJsonBody != nil {
		resp := common.ReplyUtil(false, "", "send teams error-3: "+errFromJsonBody.Error())
		return c.JSON(http.StatusBadRequest, resp)
	}

	response := common.ReplyUtil(true, "", "success")
	return c.JSON(http.StatusOK, response)
}
