package admin

import (
	"gerenciador-condominio/internal/handler"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterAdminTentantRoutes(r *gin.Engine, tenantHandler *handler.TenantHandler) {
	r.GET("/list", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"list": "[]"})
	})
	r.POST("/admin/tenants/create", tenantHandler.CreateTenants)
	r.GET("/admin/tenants/list", tenantHandler.List)
}
