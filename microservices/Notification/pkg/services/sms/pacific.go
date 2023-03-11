package sms_service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"notificaiton/pkg/utils/common"
	"strings"
)

type Pacific struct {
}

func (pacific *Pacific) SendSms(service SmsService) *common.ResponseMessage {
	requestBody := []byte(`{
		"from": "` + service.From + `",
		"to": "` + strings.Join(service.Phones, ",") + `",
		"text":"` + service.Text + `",
		"universal":false,
		"alphabet":"Default"
	}`)

	client := &http.Client{}
	service.Url += "sms/submit"
	req, err := http.NewRequest("POST", service.Url, bytes.NewBuffer(requestBody))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", service.Password)
	responseFromApi, err := client.Do(req)
	if err != nil {
		return common.ReplyUtil(false, "", "send sms error-1: "+err.Error())
	}

	var res map[string]interface{}
	errFromJsonBody := json.NewDecoder(responseFromApi.Body).Decode(&res)
	fmt.Println(res)
	if responseFromApi.StatusCode > 300 {
		detail := "no detail"
		_, isDetailExist := res["detail"]
		if isDetailExist {
			detail = res["detail"].(string)
		}

		return common.ReplyUtil(false, "", "send sms error-2: "+detail)
	}

	if errFromJsonBody != nil {
		return common.ReplyUtil(false, "", "send sms error-3: "+errFromJsonBody.Error())
	}

	return common.ReplyUtil(true, "", "success")
}

func (pacific *Pacific) IsActiveProvider() bool {
	return true
}
