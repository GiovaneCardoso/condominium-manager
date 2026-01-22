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
	repo := memory.NewTenantInMemory()
	svc := service.NewTenantService(repo)
	tenantHandler := handler.NewTenantHandler(svc)
	admin.RegisterAdminTentantRoutes(r, tenantHandler)
	r.Run(":8080")
}
