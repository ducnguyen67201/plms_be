package user_http

import (
	user_app "plms_be/internal/application/user"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine, appService *user_app.UserAppService) { 
	h := NewHandler(appService)
	
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})
	UserGroup := router.Group("/user")
	{
		UserGroup.POST("/register", h.RegisterUser)
		UserGroup.POST("/login", h.LoginUser)
	}
}