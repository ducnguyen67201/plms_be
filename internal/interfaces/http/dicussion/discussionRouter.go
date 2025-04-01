package discussion_http

import (
	discussion_app "plms_be/internal/application/discussion"

	"github.com/gin-gonic/gin"
)

func RegisterDiscussionRoutes(router *gin.Engine, appService *discussion_app.DiscussionAppService) {
	h := NewDiscussionHandler(appService)

	DiscussionGroup := router.Group("/discussion")
	{
		DiscussionGroup.POST("/all", h.GetAllDiscussion)
	}
}