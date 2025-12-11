package helpers

import (
	"context"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type ClaimToken struct {
	UserID      int    `json:"user_id"`
	Username    string `json:"username"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	FullName    string `json:"full_name"`
	Address     string `json:"address"`
	Dob         string `json:"dob"`
	jwt.RegisteredClaims
}

type PayloadToken struct {
	UserID int `json:"user_id"`
	Username string `json:"username"`
	Email string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	FullName string `json:"full_name"`
	Address string `json:"address"`
	Dob string `json:"dob"`
}

var MapTypeToken = map[string]time.Duration{
	"access_token": time.Hour * 3,
	"refresh_token": time.Hour * 72,
}

var jwtSecretKey = []byte(GetEnv("JWT_SECRET_KEY", "secret"))

func GenerateToken(ctx context.Context, payload *PayloadToken, tokenType string) (string, error) {
	claims := &ClaimToken{
		UserID: payload.UserID,
		Username: payload.Username,
		Email: payload.Email,
		PhoneNumber: payload.PhoneNumber,
		FullName: payload.FullName,
		Address: payload.Address,
		Dob: payload.Dob,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer: GetEnv("APP_NAME", "ewallet-ums"),
			IssuedAt: jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(MapTypeToken[tokenType])),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	resultToken, err := token.SignedString(jwtSecretKey)
	if err != nil {
		return "", err
	}

	return resultToken, nil
}

func ValidateToken(ctx context.Context, token string) (*ClaimToken, error) {

	var (
		claimToken *ClaimToken
		ok         bool
	)

	jwtToken, err := jwt.ParseWithClaims(token, &ClaimToken{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("failed to validate method jwt: %v", t.Header["alg"])
		}
		return jwtSecretKey, nil
	})

	if err != nil {
		return nil, fmt.Errorf("failed to parse jwt: %v", err)
	}

	if claimToken, ok = jwtToken.Claims.(*ClaimToken); !ok || !jwtToken.Valid {
		return nil, fmt.Errorf("token invalid")
	}

	return claimToken, nil
}