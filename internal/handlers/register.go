package handlers

import (
	"ewallet-ums/constants"
	"ewallet-ums/helpers"
	"ewallet-ums/internal/interfaces"
	"ewallet-ums/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RegisterHandler struct {
	RegisterService interfaces.IRegisterService
}

func (h *RegisterHandler) RegisterHandler(c *gin.Context) {
	var (
		log = helpers.Logger
	)

	req := &models.RegisterRequest{}

	if err := c.ShouldBindJSON(req); err != nil {
		log.Error("Error binding JSON", err)
		helpers.SendResponseHTTP(c, http.StatusBadRequest, constants.ErrBadRequest, nil)
		return
	}

	if err := req.ValidateRegister(); err != nil {
		log.Error("Error validating request", err)
		helpers.SendResponseHTTP(c, http.StatusBadRequest, constants.ErrBadRequest, nil)
		return
	}

	resp, err := h.RegisterService.Register(c.Request.Context(), req)
	if err != nil {
		log.Error("Error registering user", err)
		helpers.SendResponseHTTP(c, http.StatusInternalServerError, constants.ErrInternalServer, nil)
		return
	}

	helpers.SendResponseHTTP(c, http.StatusOK, constants.SuccessMessage, resp)
}

