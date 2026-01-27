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
	adminUserHandler := adminUserHandler()
	admin.RegisterAdminTentantRoutes(r, adminTenantHandler)
	admin.RegisterAdminUserRoutes(r, adminUserHandler)
	r.Run(":8080")
}
func adminTenantHandler() *handler.TenantHandler {
	repository := memory.NewTenantInMemory()
	service := service.NewTenantService(repository)
	tenantHandler := handler.NewTenantHandler(service)
	return tenantHandler
}
func adminUserHandler() *handler.UserHandler {
	repository := memory.NewAdminUserInMemory()
	service := service.NewUserAdminService(repository)
	userHandler := handler.NewUserAdminHandler(service)
	return userHandler

}