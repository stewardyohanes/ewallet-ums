package handlers

import (
	"context"
	tokenvalidation "ewallet-ums/cmd/proto"
	"ewallet-ums/helpers"
	"ewallet-ums/internal/interfaces"
	"fmt"
)

type TokenValidationHandler struct {
	tokenvalidation.UnimplementedTokenValidationServer
	TokenValidationService interfaces.ITokenValidationService
}

func (h *TokenValidationHandler) ValidateToken(ctx context.Context, req *tokenvalidation.TokenRequest) (*tokenvalidation.TokenResponse, error) {
	var (
		token = req.Token
		log = helpers.Logger
	)

	if token == "" {
		err := fmt.Errorf("token is required")
		log.Error("token is required", err)
		return &tokenvalidation.TokenResponse{
			Message: "token is required",
		}, nil
	}

	claimToken, err := h.TokenValidationService.TokenValidation(ctx, token)
	if err != nil {
		log.Error("error validate token", err)
		return &tokenvalidation.TokenResponse{
			Message: "error validate token",
		}, nil
	}	

	return &tokenvalidation.TokenResponse{
		Message: "token is valid",
		Data: &tokenvalidation.UserData{
			UserId: int64(claimToken.UserID),
			Username: claimToken.Username,
			FullName: claimToken.FullName,
			Email: claimToken.Email,
		},
	}, nil
}