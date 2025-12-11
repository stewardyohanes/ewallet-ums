package interfaces

import (
	"context"
	"ewallet-ums/internal/models"
)

type IUserRepository interface {
	InsertNewUser(ctx context.Context, user *models.Users) error
	GetUserByUsername(ctx context.Context, username string) (*models.Users, error)
	InsertNewUserSession(ctx context.Context, userSession *models.UserSessions) error
	DeleteUserSession(ctx context.Context, userSession *models.UserSessions) error
	UpdateTokenWByRefreshToken(ctx context.Context, token string, refreshToken string) error
	GetUserSessionByToken(ctx context.Context, token string) (models.UserSessions, error)
	GetUserSessionByRefreshToken(ctx context.Context, refreshToken string) (models.UserSessions, error)
}