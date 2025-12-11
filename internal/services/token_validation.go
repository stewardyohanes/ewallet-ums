package services

import (
	"context"
	"ewallet-ums/helpers"
	"ewallet-ums/internal/interfaces"
	"time"
)

type TokenValidationService struct {
	UserRepo interfaces.IUserRepository
}

func (s *TokenValidationService) TokenValidation(ctx context.Context, token string) (*helpers.ClaimToken, error) {
	var (
		claimToken *helpers.ClaimToken
		err error
	)

	dbCtx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	claimToken, err = helpers.ValidateToken(dbCtx, token)
	if err != nil {
		return claimToken, err
	}

	_, err = s.UserRepo.GetUserSessionByToken(dbCtx, token)
	if err != nil {
		return claimToken, err
	}

	return claimToken, nil
}