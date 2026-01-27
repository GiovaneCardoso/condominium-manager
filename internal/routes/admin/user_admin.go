package admin

import (
	"gerenciador-condominio/internal/handler"
	"gerenciador-condominio/internal/repository"

	"github.com/gin-gonic/gin"
)

func RegisterAdminUserRoutes(r *gin.Engine, userHandler *handler.UserHandler) {
	admin := r.Group("/admin/users")
	admin.POST("/create", userHandler.CreateAdminUser)
	admin.GET("/list", userHandler.List)
	admin.PATCH("/update/:userId", func(ctx *gin.Context) {
		id := ctx.Param("userId")
		var body repository.AdminUserUpdate
		if err := ctx.ShouldBindJSON(&body); err != nil {
			ctx.JSON(400, gin.H{"error": err.Error()})
			return
		}

		userHandler.Update(ctx, id, body)

	})
	admin.PATCH("/inactivate/:userId", func(ctx *gin.Context) {
		userId := ctx.Param("userId")
		userHandler.Inactivate(ctx, userId)
	})
}
