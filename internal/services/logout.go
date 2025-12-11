package services

import (
	"context"
	"ewallet-ums/internal/interfaces"
	"ewallet-ums/internal/models"
	"time"
)

type LogoutService struct {
	UserRepo interfaces.IUserRepository
}

func (s *LogoutService) Logout(ctx context.Context, token string) error {
	dbCtx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	
	err := s.UserRepo.DeleteUserSession(dbCtx, &models.UserSessions{
		Token: token,
	})
	if err != nil {
		return err
	}

	return nil
}