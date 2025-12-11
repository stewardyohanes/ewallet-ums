package services

import (
	"context"
	"ewallet-ums/helpers"
	"ewallet-ums/internal/interfaces"
	"ewallet-ums/internal/models"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type LoginService struct {
	UserRepo interfaces.IUserRepository
}

func (s *LoginService) Login(ctx context.Context, req *models.LoginRequest) (models.LoginResponse, error) {
	var (
		loginResponse models.LoginResponse
	)

	dbCtx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	user, err := s.UserRepo.GetUserByUsername(dbCtx, req.Username)
	if err != nil {
		return loginResponse, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return loginResponse, err
	}

	payloadToken := &helpers.PayloadToken{
		UserID: user.ID,
		Username: user.Username,
		Email: user.Email,
		PhoneNumber: user.PhoneNumber,
		FullName: user.FullName,
		Address: user.Address,
		Dob: user.Dob,
	}

	accessToken, err := helpers.GenerateToken(ctx, payloadToken, "access_token")
	if err != nil {
		return loginResponse, err
	}

	refreshToken, err := helpers.GenerateToken(ctx, payloadToken, "refresh_token")
	if err != nil {
		return loginResponse, err
	}

	loginResponse = models.LoginResponse{
		Token: accessToken,
		RefreshToken: refreshToken,
		TokenExpired: time.Now().Add(helpers.MapTypeToken["access_token"]),
		RefreshExpired: time.Now().Add(helpers.MapTypeToken["refresh_token"]),
	}

	userSession := &models.UserSessions{
		UserID: user.ID,
		Token: loginResponse.Token,
		RefreshToken: loginResponse.RefreshToken,
		TokenExpired: loginResponse.TokenExpired,
		RefreshTokenExpired: loginResponse.RefreshExpired,
	}

	sessionCtx, sessionCancel := context.WithTimeout(ctx, 5*time.Second)
	defer sessionCancel()

	err = s.UserRepo.InsertNewUserSession(sessionCtx, userSession)
	if err != nil {
		return loginResponse, err
	}

	return loginResponse, nil
}