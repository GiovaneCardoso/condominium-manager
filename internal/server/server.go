package server

import (
	"os"

	"gerenciador-condominio/internal/handler"
	"gerenciador-condominio/internal/infra/memory"
	"gerenciador-condominio/internal/middlewares"

	authHandler "gerenciador-condominio/internal/auth/handler"
	authService "gerenciador-condominio/internal/auth/service"
	"gerenciador-condominio/internal/routes/admin"
	"gerenciador-condominio/internal/service"

	"github.com/gin-gonic/gin"
)

func StartServer() {
	r := gin.Default()
	
	adminUserRepo := memory.NewAdminUserInMemory()
	tenantRepo := memory.NewTenantInMemory()

	adminTenantHandler := adminTenantHandler(tenantRepo)
	adminUserHandler := adminUserHandler(adminUserRepo)
	authHandlerInstance := authHandlerInstance(adminUserRepo)

	admin.RegisterAuthRoutes(r, authHandlerInstance)

	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		jwtSecret = "dev-secret-key"
	}
	tokenService := authService.NewTokenService(jwtSecret)
	_ = middlewares.AuthMiddleware(tokenService, adminUserRepo)

	admin.RegisterAdminTentantRoutes(r, adminTenantHandler)
	admin.RegisterAdminUserRoutes(r, adminUserHandler)

	r.Run(":8080")
}

func adminTenantHandler(repo *memory.TenantInMemory) *handler.TenantHandler {
	service := service.NewTenantService(repo)
	tenantHandler := handler.NewTenantHandler(service)
	return tenantHandler
}

func authHandlerInstance(adminUserRepo *memory.AdminUserInMemory) *authHandler.AuthHandler {
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		jwtSecret = "dev-secret-key"
	}
	tokenService := authService.NewTokenService(jwtSecret)
	authServiceInstance := authService.NewAuthService(adminUserRepo, tokenService)
	return authHandler.NewAuthHandler(authServiceInstance)
}

func adminUserHandler(repo *memory.AdminUserInMemory) *handler.UserHandler {
	service := service.NewUserAdminService(repo)
	userHandler := handler.NewUserAdminHandler(service)
	return userHandler
}