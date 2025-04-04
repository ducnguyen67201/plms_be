package learning_http

import (
	learning_app "plms_be/internal/application/learning"

	"github.com/gin-gonic/gin"
)

func RegisterLearningRoutes(router *gin.Engine, appService *learning_app.LearningAppService) {
	h := NewLearningHandler(appService)

	LearningGroup := router.Group("/learning")
	{
		LearningGroup.POST("/all", h.GetAllLearning)
		LearningGroup.POST("/:id", h.GetLearningByID)
		LearningGroup.POST("/save", h.SaveLearning)
	}
}