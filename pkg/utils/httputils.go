package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"golang.org/x/exp/slices"
	"net/http"
	"os"
	"strings"
)

type ResponseMessage struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

func (r ResponseMessage) getSuccess() bool {
	return r.Success
}

func (r ResponseMessage) getMessage() string {
	return r.Message
}

type BMCInitialParameters struct {
	jsonWebToken string
}

func (i *BMCInitialParameters) setInitials() (init *BMCInitialParameters, err error) {
	type accessRequest struct {
		AccessKey       string `json:"access_key"`
		AccessSecretKey string `json:"access_secret_key"`
	}
	requestBody := accessRequest{}
	requestBody.AccessKey = os.Getenv("ACCESS_KEY")
	requestBody.AccessSecretKey = os.Getenv("ACCESS_SECRET_KEY")
	responseFromAPI, err := HttpPost("/ims/api/v1/access_keys/login", requestBody, "")
	if err != nil {
		return nil, err
	}
	jsonWebToken, isExist := responseFromAPI["json_web_token"]
	if !isExist || strings.Trim(jsonWebToken, " ") == "" {
		jsonWebToken = ""
	}
	return &BMCInitialParameters{jsonWebToken: jsonWebToken}, nil
}

func HttpPost(URI string, requestBody interface{}, contentType string) (data map[string]string, err error) {
	var validContentTypes = []string{
		"application/json",
		"application/xml",
		"multipart/form-data",
		"application/x-www-form-urlencoded",
	}

	if strings.Trim(contentType, " ") == "" || !slices.Contains(validContentTypes, contentType) {
		contentType = "application/json"
	}
	url := os.Getenv("URL") + URI
	body, err := json.Marshal(requestBody)
	if err != nil {
		return nil, err
	}
	responseFromApi, err := http.Post(url, contentType, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	var res map[string]string
	errFromJsonBody := json.NewDecoder(responseFromApi.Body).Decode(&res)
	if errFromJsonBody != nil {
		return nil, err
	}
	return res, nil
}

func HttpGet(URI string, contentType string) (data map[string]interface{}, err error) {
	var validContentTypes = []string{
		"application/json",
		"application/xml",
		"multipart/form-data",
		"application/x-www-form-urlencoded",
	}

	if strings.Trim(contentType, " ") == "" || !slices.Contains(validContentTypes, contentType) {
		contentType = "application/json"
	}
	bmci := new(BMCInitialParameters)
	bmci, err = bmci.setInitials()
	if err != nil {
		return nil, err
	}

	url := os.Getenv("URL") + URI
	responseFromApi, err := http.Get(url)
	responseFromApi.Header.Set("Content-Type", contentType)
	responseFromApi.Header.Add("Authorization", "Bearer "+bmci.jsonWebToken)
	fmt.Println(responseFromApi)
	if err != nil {
		return nil, err
	}
	var res map[string]interface{}
	errFromJsonBody := json.NewDecoder(responseFromApi.Body).Decode(&res)

	if errFromJsonBody != nil {
		return nil, err
	}
	return res, nil
}

func ReplyUtil(success bool, data interface{}, message string) *ResponseMessage {
	return &ResponseMessage{success, data, message}
}
