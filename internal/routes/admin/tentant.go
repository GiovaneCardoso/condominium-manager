package admin

import (
	"gerenciador-condominio/internal/handler"
	"gerenciador-condominio/internal/repository"

	"github.com/gin-gonic/gin"
)

func RegisterAdminTentantRoutes(r *gin.Engine, tenantHandler *handler.TenantHandler) {
	admin := r.Group("/admin/tenants")
	admin.POST("/create", tenantHandler.CreateTenants)
	admin.GET("/list", tenantHandler.List)
	admin.PATCH("/update/:tenantId", func(ctx *gin.Context) {
		id := ctx.Param("tenantId")
		var body repository.TenantUpdate
		if err := ctx.ShouldBindJSON(&body); err != nil {
			ctx.JSON(400, gin.H{"error": err.Error()})
			return
		}

		tenantHandler.Update(ctx, id, body)

	})
	admin.PATCH("/inactivate/:tenantId", func(ctx *gin.Context) {
		tenantId := ctx.Param("tenantId")
		tenantHandler.Inactivate(ctx, tenantId)
	})
}
