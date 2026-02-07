package admin

import (
	authhandler "gerenciador-condominio/internal/auth/handler"

	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(r *gin.Engine, authHandler *authhandler.AuthHandler) {
	auth := r.Group("/auth")
	{
		auth.POST("/login", authHandler.Login)
	}
}
