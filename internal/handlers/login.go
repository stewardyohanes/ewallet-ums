package handlers

import (
	"ewallet-ums/constants"
	"ewallet-ums/helpers"
	"ewallet-ums/internal/interfaces"
	"ewallet-ums/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginHandler struct {
	LoginService interfaces.ILoginService
}

func (h *LoginHandler) LoginHandler(c *gin.Context) {
	var (
		log = helpers.Logger
	)

	req := &models.LoginRequest{}

	if err := c.ShouldBindJSON(req); err != nil {
		log.Error("Error binding JSON", err)
		helpers.SendResponseHTTP(c, http.StatusBadRequest, constants.ErrBadRequest, nil)
		return
	}

	if err := req.ValidateLogin(); err != nil {
		log.Error("Error validating request", err)
		helpers.SendResponseHTTP(c, http.StatusBadRequest, constants.ErrBadRequest, nil)
		return
	}
	
	response, err := h.LoginService.Login(c.Request.Context(), req)
	if err != nil {
		log.Error("Error login", err)	
		helpers.SendResponseHTTP(c, http.StatusInternalServerError, constants.ErrInternalServer, nil)
		return
	}

	helpers.SendResponseHTTP(c, http.StatusOK, constants.SuccessMessage, response)
}