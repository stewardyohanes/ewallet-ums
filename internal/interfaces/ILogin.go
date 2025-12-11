package interfaces

import (
	"context"
	"ewallet-ums/internal/models"

	"github.com/gin-gonic/gin"
)

type (
	ILoginService interface {
		Login(ctx context.Context, req *models.LoginRequest) (models.LoginResponse, error)
	}

	ILoginHandler interface {
		LoginHandler(c *gin.Context)
	}
)