package controllers

import (
	"database/pkg/utils/common"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (h *Handler) FetchAgreementAndTerms(c echo.Context) error {
	agreementAndTerms, err := h.AgreementAndTermsRepository.FetchAgreementAndTerms()
	if err != nil {
		response := common.ReplyUtil(false, "", "error fetching agreement and terms: "+err.Error())
		return c.JSON(http.StatusBadRequest, response)
	}
	response := common.ReplyUtil(true, agreementAndTerms, "success")
	return c.JSON(http.StatusOK, response)
}
