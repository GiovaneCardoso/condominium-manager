package server

import (
	"gerenciador-condominio/internal/handler"
	"gerenciador-condominio/internal/infra/memory"
	"gerenciador-condominio/internal/routes/admin"
	"gerenciador-condominio/internal/service"

	"github.com/gin-gonic/gin"
)

func StartServer() {
	r := gin.Default()
	adminTenantHandler := adminTenantHandler()
	admin.RegisterAdminTentantRoutes(r, adminTenantHandler)
	r.Run(":8080")
}
func adminTenantHandler() *handler.TenantHandler {
	repository := memory.NewTenantInMemory()
	service := service.NewTenantService(repository)
	tenantHandler := handler.NewTenantHandler(service)
	return tenantHandler
}