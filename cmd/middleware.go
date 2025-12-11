package cmd

import (
	"ewallet-ums/helpers"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func (d *Dependency) MiddlewareValidateAuth(ctx *gin.Context)  {
	auth := ctx.Request.Header.Get("Authorization")

	if auth == "" {
		log.Println("Unauthorized: No token provided")
		helpers.SendResponseHTTP(ctx, http.StatusUnauthorized, "Unauthorized", nil)
		ctx.Abort()
		return
	}

	_, err := d.UserRepository.GetUserSessionByToken(ctx.Request.Context(), auth)
	if err != nil {
		log.Println("Unauthorized: User session not found")
		helpers.SendResponseHTTP(ctx, http.StatusUnauthorized, "Unauthorized", nil)
		ctx.Abort()
		return
	}

	claimToken, err := helpers.ValidateToken(ctx.Request.Context(), auth)
	if err != nil {
		log.Println("Unauthorized: Invalid token", err)
		helpers.SendResponseHTTP(ctx, http.StatusUnauthorized, "Unauthorized", nil)
		ctx.Abort()
		return
	}

	if time.Now().Unix() > claimToken.ExpiresAt.Unix() {
		log.Println("Unauthorized: Token expired")
		helpers.SendResponseHTTP(ctx, http.StatusUnauthorized, "Unauthorized", nil)
		ctx.Abort()
		return
	}

	ctx.Set("token", claimToken)

	ctx.Next()
}

func (d *Dependency) MiddlewareRefreshToken(ctx *gin.Context)  {
	auth := ctx.Request.Header.Get("Authorization")

	if auth == "" {
		log.Println("Unauthorized: No token provided")
		helpers.SendResponseHTTP(ctx, http.StatusUnauthorized, "Unauthorized", nil)
		ctx.Abort()
		return
	}

	claimToken, err := helpers.ValidateToken(ctx.Request.Context(), auth)
	if err != nil {
		log.Println("Unauthorized: Invalid token", err)
		helpers.SendResponseHTTP(ctx, http.StatusUnauthorized, "Unauthorized", nil)
		ctx.Abort()
		return
	}

	if time.Now().Unix() > claimToken.ExpiresAt.Unix() {
		log.Println("Unauthorized: Token expired")
		helpers.SendResponseHTTP(ctx, http.StatusUnauthorized, "Unauthorized", nil)
		ctx.Abort()
		return
	}

	ctx.Set("token", claimToken)

	ctx.Next()
}