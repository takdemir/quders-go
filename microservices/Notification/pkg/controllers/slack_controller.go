package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/slack-go/slack"
	"golang.org/x/exp/slices"
	"net/http"
	"notificaiton/pkg/structs"
	"notificaiton/pkg/utils/common"
	"os"
	"strconv"
	"strings"
	"time"
)

func (h *Handler) SendSlackMessage(c echo.Context) error {

	jsonMap := structs.SlackFeatures{}

	err := json.NewDecoder(c.Request().Body).Decode(&jsonMap)
	if err != nil {
		response := common.ReplyUtil(false, "", fmt.Sprintf("slack request body content parse error %s", err.Error()))
		return c.JSON(http.StatusOK, response)
	}

	ok, response := verifySlackRequestContent(jsonMap)
	if !ok {
		return c.JSON(http.StatusBadRequest, response)
	}

	attachment := slack.Attachment{
		Color:      jsonMap.Color,
		AuthorName: jsonMap.AuthorName,
		AuthorLink: jsonMap.AuthorLink,
		AuthorIcon: jsonMap.AuthorIcon,
		Text:       jsonMap.Text,
		Ts:         jsonMap.Ts,
	}

	msg := slack.WebhookMessage{
		Attachments: []slack.Attachment{attachment},
	}

	slackError := slack.PostWebhook(os.Getenv("SLACK_WEBHOOK"), &msg)
	if slackError != nil {
		res := common.ReplyUtil(false, "", "slack post message error: "+slackError.Error())
		return c.JSON(http.StatusBadRequest, res)
	}

	res := common.ReplyUtil(true, "", "success")
	return c.JSON(http.StatusOK, res)
}

func verifySlackRequestContent(jsonMap structs.SlackFeatures) (bool, *common.ResponseMessage) {
	validColorNames := []string{"Good", "Bad"}
	if strings.Trim(jsonMap.Color, " ") == "" || !slices.Contains(validColorNames, jsonMap.Color) {
		jsonMap.Color = "good"
	}
	if strings.Trim(jsonMap.AuthorName, " ") == "" {
		jsonMap.AuthorName = "Basistek"
	}
	if strings.Trim(jsonMap.AuthorLink, " ") == "" {
		jsonMap.AuthorName = "https://basistek.com"
	}
	if strings.Trim(jsonMap.AuthorIcon, " ") == "" {
		jsonMap.AuthorIcon = "https://static.wixstatic.com/media/72ac26_93a5ba6baaa741ffaa0d06db680205e6~mv2.png/v1/fill/w_206,h_45,al_c,q_85,usm_0.66_1.00_0.01,enc_auto/basistek%20logo-TR.png"
	}
	if strings.Trim(jsonMap.Text, " ") == "" {
		return false, common.ReplyUtil(false, jsonMap, "text message can not be empty")
	}
	jsonMap.Ts = json.Number(strconv.FormatInt(time.Now().Unix(), 10))
	return true, common.ReplyUtil(true, jsonMap, "success")
}
