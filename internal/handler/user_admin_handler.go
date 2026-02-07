package handler

import (
	"fmt"
	"gerenciador-condominio/internal/domain"
	"gerenciador-condominio/internal/repository"
	"gerenciador-condominio/internal/service"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service *service.UserAdminService
}
type CreateAdminUserRequest struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func NewUserAdminHandler(service *service.UserAdminService) *UserHandler {
	return &UserHandler{service: service}
}
func (h *UserHandler) CreateAdminUser(c *gin.Context) {
	fmt.Println("teste")
	var req CreateAdminUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	user := domain.AdminUser{
		Email:    req.Email,
		Password: req.Password,
		Name:     req.Name,
	}
	userCreated, err := h.service.Create(user)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(201, userCreated)
}
func (h *UserHandler) List(c *gin.Context) {
	users, err := h.service.List()
	if err != nil {
		c.JSON(503, err)
		return
	}
	c.JSON(200, users)
}
func (h *UserHandler) Update(c *gin.Context, id string, updateUser repository.AdminUserUpdate) {
	userUpdated, err := h.service.Update(id, updateUser)
	if err != nil {
		c.JSON(503, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, userUpdated)
}
func (h *UserHandler) Inactivate(c *gin.Context, id string) {
	err := h.service.Inactivate(id)
	if err != nil {
		c.JSON(503, err)
		return
	}
	c.Status(200)
}
