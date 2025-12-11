package handlers

import (
	"ewallet-ums/helpers"
	"ewallet-ums/internal/interfaces"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthCheck struct {
	HealthCheckService interfaces.IHealthCheckService
}

func (h *HealthCheck) HealthCheckHandlerHTTP(c *gin.Context) {
	msg, err := h.HealthCheckService.HealthCheckService()
	if err != nil {
		helpers.SendResponseHTTP(c, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	helpers.SendResponseHTTP(c, http.StatusOK, msg, nil)
}
