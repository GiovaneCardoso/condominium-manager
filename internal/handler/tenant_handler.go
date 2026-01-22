package handler

import (
	"gerenciador-condominio/internal/domain"
	"gerenciador-condominio/internal/service"

	"github.com/gin-gonic/gin"
)

type TenantHandler struct {
	service *service.TenantService
}
type CreateTenantRequest struct {
	Name   string `json:"name" binding:"required"`
	Domain string `json:"domain" binding:"required"`
}

func NewTenantHandler(service *service.TenantService) *TenantHandler {
	return &TenantHandler{service: service}
}
func (h *TenantHandler) CreateTenants(c *gin.Context) {
	var req CreateTenantRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	tenant := domain.Tenant{
		Name:   req.Name,
		Domain: req.Domain,
	}
	if err := h.service.Create(tenant); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(201, tenant)
}
func (h *TenantHandler) List(c *gin.Context) {
	tenantList, err := h.service.List()

	if err != nil {
		c.JSON(403, err)
	}
	c.JSON(200, tenantList)
}
