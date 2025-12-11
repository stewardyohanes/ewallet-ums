package models

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type (
	LoginRequest struct {
		Username    string `json:"username" validate:"required"`
		Password    string `json:"password" validate:"required"`
	}

	LoginResponse struct {
		Token         string `json:"token"`
		RefreshToken  string `json:"refresh_token"`
		TokenExpired  time.Time `json:"token_expired"`
		RefreshExpired time.Time `json:"refresh_expired"`
	}

	RegisterRequest struct {
		Username    string `json:"username" validate:"required"`
		Email       string `json:"email" validate:"required,email"`
		PhoneNumber string `json:"phone_number" validate:"required"`
		FullName    string `json:"full_name" validate:"required"`
		Address     string `json:"address" validate:"required"`
		Dob         string `json:"dob" validate:"required"`
		Password    string `json:"password" validate:"required"`
	}

	RegisterResponse struct {
		ID          int       `json:"id"`
		Username    string    `json:"username" gorm:"colum:username;type:varchar(20)"`
		Email       string    `json:"email" gorm:"colum:email;type:varchar(100)"`
		PhoneNumber string    `json:"phone_number" gorm:"colum:phone_number;type:varchar(15)"`
		FullName    string    `json:"full_name" gorm:"colum:full_name;type:varchar(100)"`
		Address     string    `json:"address" gorm:"colum:address;type:text"`
		Dob         string    `json:"dob" gorm:"colum:dob;type:date"`
		CreatedAt   time.Time `json:"-"`
		UpdatedAt   time.Time `json:"-"`
	}

	RefreshTokenResponse struct {
		Token         string `json:"token"`
	}
)

type (
	Users struct {
		ID          int       `json:"id" `
		Username    string    `json:"username" gorm:"colum:username;type:varchar(20)" validate:"required"`
		Email       string    `json:"email" gorm:"colum:email;type:varchar(100)" validate:"required"`
		PhoneNumber string    `json:"phone_number" gorm:"colum:phone_number;type:varchar(15)" validate:"required"`
		FullName    string    `json:"full_name" gorm:"colum:full_name;type:varchar(100)" validate:"required"`
		Address     string    `json:"address" gorm:"colum:address;type:text"`
		Dob         string    `json:"dob" gorm:"colum:dob;type:date"`
		Password    string    `json:"password,omitempty" gorm:"colum:password;type:varchar(255)" validate:"required"`
		CreatedAt   time.Time `json:"-" `
		UpdatedAt   time.Time `json:"-" `
	}

	UserSessions struct {
	ID                  int `gorm:"primarykey"`
	CreatedAt           time.Time
	UpdatedAt           time.Time
	UserID              int       `json:"user_id" gorm:"type:int" validate:"required"`
	Token               string    `json:"token" gorm:"type:text" validate:"required"`
	RefreshToken        string    `json:"refresh_token" gorm:"type:text" validate:"required"`
	TokenExpired        time.Time `json:"-" validate:"required"`
	RefreshTokenExpired time.Time `json:"-" validate:"required"`
}
)

func (*Users) TableName() string {
	return "users"
}
func (l RegisterRequest) ValidateRegister() error {
	v := validator.New()
	return v.Struct(l)
}

func (l LoginRequest) ValidateLogin() error {
	v := validator.New()
	return v.Struct(l)
}

func (*UserSessions) TableName() string {
	return "user_sessions"
}

func (l UserSessions) Validate() error {
	v := validator.New()
	return v.Struct(l)
}