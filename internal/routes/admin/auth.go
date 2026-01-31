package admin

import (
	authHandler "gerenciador-condominio/internal/auth/handler"

	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(r *gin.Engine, authHandler *authHandler.AuthHandler) {
	auth := r.Group("/auth")
	{
		auth.POST("/login", authHandler.Login)
	}
}
