package handler

import (
	"gerenciador-condominio/internal/domain"
	"gerenciador-condominio/internal/repository"
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
	tenantCreated, err := h.service.Create(tenant)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(201, tenantCreated)
}
func (h *TenantHandler) List(c *gin.Context) {
	tenantList, err := h.service.List()

	if err != nil {
		c.JSON(403, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, tenantList)
}
func (h *TenantHandler) Update(c *gin.Context, id string, t repository.TenantUpdate) {
	tenantUpdated, err := h.service.Update(id, t)
	if err != nil {
		c.JSON(503, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, tenantUpdated)
}
func (h *TenantHandler) Inactivate(c *gin.Context, id string) {
	err := h.service.Inactivate(id)
	if err != nil {
		c.JSON(503, err)
		return
	}
	c.Status(200)
}