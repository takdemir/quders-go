package httputils

import (
	"bytes"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"golang.org/x/exp/slices"
	"net/http"
	"strings"
)

func HttpGet(c echo.Context, url string, contentType string) (data map[string]interface{}, err error) {
	var validContentTypes = []string{
		"application/json",
		"application/xml",
		"multipart/form-data",
		"application/x-www-form-urlencoded",
	}

	if strings.Trim(contentType, " ") == "" || !slices.Contains(validContentTypes, contentType) {
		contentType = "application/json"
	}
	auth := c.Request().Header.Get("x-api-token")
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("Content-Type", contentType)
	req.Header.Add("x-api-token", auth)
	//fmt.Println(reflect.ValueOf(req.Header).MapKeys())
	responseFromApi, err := client.Do(req)
	//fmt.Println(responseFromApi)
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

func HttpPost(c echo.Context, url string, requestBody []byte, contentType string) (data map[string]interface{}, err error) {
	var validContentTypes = []string{
		"application/json",
		"application/xml",
		"multipart/form-data",
		"application/x-www-form-urlencoded",
	}

	if strings.Trim(contentType, " ") == "" || !slices.Contains(validContentTypes, contentType) {
		contentType = "application/json"
	}
	auth := c.Request().Header.Get("x-api-token")
	//body, err := json.Marshal(requestBody)
	if err != nil {
		return nil, err
	}
	client := &http.Client{}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
	req.Header.Set("Content-Type", contentType)
	req.Header.Add("x-api-token", auth)

	responseFromApi, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if responseFromApi.StatusCode > 300 {

	}
	var res map[string]interface{}
	errFromJsonBody := json.NewDecoder(responseFromApi.Body).Decode(&res)

	if errFromJsonBody != nil {
		return nil, errFromJsonBody
	}

	return res, nil
}
