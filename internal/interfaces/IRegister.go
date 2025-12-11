package interfaces

import (
	"context"
	"ewallet-ums/internal/models"

	"github.com/gin-gonic/gin"
)

type (
	IRegisterService interface {
		Register(ctx context.Context, req *models.RegisterRequest) (models.RegisterResponse, error)
	}

	IRegisterHandler interface {
		RegisterHandler(c *gin.Context)
	}
)