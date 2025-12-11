package handlers

import (
	"ewallet-ums/constants"
	"ewallet-ums/helpers"
	"ewallet-ums/internal/interfaces"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RefreshTokenHandler struct {
	RefreshTokenService interfaces.IRefreshTokenService
}

func (h *RefreshTokenHandler) RefreshToken(ctx *gin.Context) {
	var (
		log = helpers.Logger
	)

	refreshToken := ctx.GetHeader("Authorization")

	if refreshToken == "" {
		log.Error("Refresh-Token is required")
		helpers.SendResponseHTTP(ctx, http.StatusBadRequest, "Refresh-Token is required", nil)
		return
	}

	claim, ok := ctx.Get("token")
	if !ok {
		log.Error("Failed to get Claim in context")
		helpers.SendResponseHTTP(ctx, http.StatusBadRequest, "Failed to get Claim in context", nil)
		return
	}

	claimToken, ok := claim.(*helpers.ClaimToken)
	if !ok {
		log.Error("failed to parse claim to claimToken")
		helpers.SendResponseHTTP(ctx, http.StatusInternalServerError, constants.ErrInternalServer, nil)
		return
	}

	resp, err := h.RefreshTokenService.RefreshToken(ctx.Request.Context(), refreshToken, *claimToken)
	if err != nil {
		log.Error("Error refresh token", err)
		helpers.SendResponseHTTP(ctx, http.StatusInternalServerError, constants.ErrInternalServer, nil)
		return
	}

	helpers.SendResponseHTTP(ctx, http.StatusOK, constants.SuccessMessage, resp)
}