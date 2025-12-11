package repositories

import (
	"context"
	"errors"
	"ewallet-ums/internal/models"

	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func (r *UserRepository) InsertNewUser(ctx context.Context, user *models.Users) error {
	return r.DB.WithContext(ctx).Create(user).Error
}


func (r *UserRepository) GetUserByUsername(ctx context.Context, username string) (*models.Users, error) {
	var user models.Users
	
	if err := r.DB.WithContext(ctx).Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	
	return &user, nil
}

func (r *UserRepository) InsertNewUserSession(ctx context.Context, userSession *models.UserSessions) error {
	return r.DB.WithContext(ctx).Create(userSession).Error
}

func (r *UserRepository) DeleteUserSession(ctx context.Context, userSession *models.UserSessions) error {
	return r.DB.WithContext(ctx).Exec("DELETE FROM user_sessions WHERE token = ?", userSession.UserID).Error
}


func (r *UserRepository) UpdateTokenWByRefreshToken(ctx context.Context, token string, refreshToken string) error {
	return r.DB.WithContext(ctx).Exec("UPDATE user_sessions SET token = ? WHERE refresh_token = ?", token, refreshToken).Error
}

func (r *UserRepository) GetUserSessionByToken(ctx context.Context, token string) (models.UserSessions, error) {
	var (
		session models.UserSessions
		err     error
	)
	err = r.DB.Where("token = ?", token).First(&session).Error
	if err != nil {
		return session, err
	}
	if session.ID == 0 {
		return session, errors.New("session not found")
	}
	return session, nil
}
func (r *UserRepository) GetUserSessionByRefreshToken(ctx context.Context, refreshToken string) (models.UserSessions, error) {
	var (
		session models.UserSessions
		err     error
	)
	err = r.DB.Where("refresh_token = ?", refreshToken).First(&session).Error
	if err != nil {
		return session, err
	}
	if session.ID == 0 {
		return session, errors.New("session not found")
	}
	return session, nil
}