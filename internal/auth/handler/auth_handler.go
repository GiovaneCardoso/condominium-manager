package authhandler

import (
	authservice "gerenciador-condominio/internal/auth/service"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	service *authservice.AuthService
}
type LoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func NewAuthHandler(service *authservice.AuthService) *AuthHandler {
	return &AuthHandler{service: service}
}
func (h *AuthHandler) Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	token, err := h.service.Authenticate(req.Email, req.Password)
	if err != nil {
		c.JSON(403, gin.H{"error": "invalid credentials"})
		return
	}
	c.JSON(200, gin.H{"access_token": token})
}
