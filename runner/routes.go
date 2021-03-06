package runner

import (
	"github.com/gin-gonic/gin"
)

// InitializeRoutes initializes the router with handlers
func InitializeRoutes(router *gin.Engine) {
	v1 := router.Group("/runner/v1")
	{
		v1.GET("/status", statusResponseHandler)
		v1.POST("/start", startResponseHander)
	}
}
