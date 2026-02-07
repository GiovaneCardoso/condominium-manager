package server

import (
	"log"

	"gerenciador-condominio/internal/handler"
	"gerenciador-condominio/internal/infra/memory"
	"gerenciador-condominio/internal/infra/postgres"
	postgresconn "gerenciador-condominio/internal/infra/postgres/connection"
	"gerenciador-condominio/internal/repository"

	"gerenciador-condominio/internal/routes/admin"
	"gerenciador-condominio/internal/service"

	"github.com/gin-gonic/gin"
)

func StartServer() {
	r := gin.Default()

	db := postgresconn.GetConnection()
	if db == nil {
		log.Fatal("Failed to connect to database")
	}

	var adminUserRepo repository.AdminUserRepository
	var tenantRepo repository.TenantRepository
	adminUserRepo = postgres.NewAdminUserPostgres(db)
	tenantRepo = memory.NewTenantInMemory()

	adminTenantHandler := adminTenantHandler(tenantRepo)
	adminUserHandler := adminUserHandler(adminUserRepo)


	admin.RegisterAdminTentantRoutes(r, adminTenantHandler)
	admin.RegisterAdminUserRoutes(r, adminUserHandler)

	r.Run(":8080")
}

func adminTenantHandler(repo repository.TenantRepository) *handler.TenantHandler {
	service := service.NewTenantService(repo)
	tenantHandler := handler.NewTenantHandler(service)
	return tenantHandler
}


func adminUserHandler(repo repository.AdminUserRepository) *handler.UserHandler {
	service := service.NewUserAdminService(repo)
	userHandler := handler.NewUserAdminHandler(service)
	return userHandler
}