package services

import (
	"context"
	"ewallet-ums/helpers"
	"ewallet-ums/internal/interfaces"
	"ewallet-ums/internal/models"
	"time"
)

type RefreshTokenService struct {
	UserRepo interfaces.IUserRepository
}

func (s *RefreshTokenService) RefreshToken(ctx context.Context, refreshToken string, tokenClaim helpers.ClaimToken) (models.RefreshTokenResponse, error){
	dbCtx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	resp := models.RefreshTokenResponse{}

	token, err := helpers.GenerateToken(ctx, &helpers.PayloadToken{
		UserID: tokenClaim.UserID,
		Username: tokenClaim.Username,
		Email: tokenClaim.Email,
		PhoneNumber: tokenClaim.PhoneNumber,
		FullName: tokenClaim.FullName,
		Address: tokenClaim.Address,
		Dob: tokenClaim.Dob,
	}, "access")
	if err != nil {
		return resp, err
	}

	err = s.UserRepo.UpdateTokenWByRefreshToken(dbCtx, token, refreshToken)
	if err != nil {
		return resp, err
	}

	resp.Token = token

	return resp, nil
}