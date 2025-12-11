package handlers

import (
	"ewallet-ums/constants"
	"ewallet-ums/helpers"
	"ewallet-ums/internal/interfaces"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LogoutHandler struct {
	LogoutService interfaces.ILogoutService
}

func (h *LogoutHandler) LogoutHandler(c *gin.Context) {
	var (
		log = helpers.Logger
	)

	token := c.GetHeader("Authorization")

	if token == "" {
		log.Error("Token is required")
		helpers.SendResponseHTTP(c, http.StatusBadRequest, constants.ErrBadRequest, nil)
		return
	}

	err := h.LogoutService.Logout(c.Request.Context(), token)
	if err != nil {
		log.Error("Error logout", err)
		helpers.SendResponseHTTP(c, http.StatusInternalServerError, constants.ErrInternalServer, nil)
		return
	}

	helpers.SendResponseHTTP(c, http.StatusOK, constants.SuccessMessage, nil)
}