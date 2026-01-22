package test

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterTestRoutes(r *gin.Engine) {
	r.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Test route is working!"})
	})
}
