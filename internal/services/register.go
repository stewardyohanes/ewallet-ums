package services

import (
	"context"
	"ewallet-ums/internal/interfaces"
	"ewallet-ums/internal/models"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type RegisterService struct {
	UserRepo interfaces.IUserRepository
}


func (s *RegisterService) Register(ctx context.Context, req *models.RegisterRequest) (models.RegisterResponse, error) {
	dbCtx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return models.RegisterResponse{}, err
	}

	req.Password = string(hashPassword)
	
	user := &models.Users{
		Username: req.Username,
		Email: req.Email,
		PhoneNumber: req.PhoneNumber,
		FullName: req.FullName,
		Address: req.Address,
		Dob: req.Dob,
		Password: req.Password,
	}

	err = s.UserRepo.InsertNewUser(dbCtx, user)
	if err != nil {
		return models.RegisterResponse{}, err
	}

	resp := models.RegisterResponse{
		ID: user.ID,
		Username: user.Username,
		Email: user.Email,
		PhoneNumber: user.PhoneNumber,
		FullName: user.FullName,
		Address: user.Address,
		Dob: user.Dob,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	return resp, nil
}

