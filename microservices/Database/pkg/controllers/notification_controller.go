package controllers

import (
	"database/pkg/utils/common"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
	"strings"
)

func (h *Handler) GetNotifications(c echo.Context) error {
	notification, err := h.NotificationRepository.FetchNotifications()
	if err != nil {
		response := common.ReplyUtil(false, "", "error on fetching notification: "+err.Error())
		return c.JSON(http.StatusBadRequest, response)
	}
	response := common.ReplyUtil(true, notification, "success")
	return c.JSON(http.StatusOK, response)
}

func (h *Handler) GetNotificationById(c echo.Context) error {
	idParam := c.Param("id")
	if strings.TrimSpace(idParam) == "" {
		response := common.ReplyUtil(false, "", "id is not valid")
		return c.JSON(http.StatusBadRequest, response)
	}
	id, _ := strconv.Atoi(idParam)
	notification, err := h.NotificationRepository.FetchNotificationById(id)
	if err != nil {
		response := common.ReplyUtil(false, "", "error on fetching notification by id: "+err.Error())
		return c.JSON(http.StatusBadRequest, response)
	}
	if notification.ID == 0 {
		response := common.ReplyUtil(false, "", "not found")
		return c.JSON(http.StatusNotFound, response)
	}
	response := common.ReplyUtil(true, notification, "success")
	return c.JSON(http.StatusOK, response)
}

func (h *Handler) GetActiveNotifications(c echo.Context) error {
	notification, err := h.NotificationRepository.FetchActiveNotifications()
	if err != nil {
		response := common.ReplyUtil(false, "", "error on fetching active notification: "+err.Error())
		return c.JSON(http.StatusBadRequest, response)
	}
	response := common.ReplyUtil(true, notification, "success")
	return c.JSON(http.StatusOK, response)
}
