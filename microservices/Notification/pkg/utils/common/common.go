package common

import (
	"github.com/labstack/echo/v4"
	"log"
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

func FailOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func userFromToken(c echo.Context) string {
	email, ok := c.Get("userEmail").(string)
	if !ok {
		return ""
	}
	return email
}

func ReplyUtil(success bool, data interface{}, message string) *ResponseMessage {
	return &ResponseMessage{success, data, message}
}
