package discussion_http

import (
	"net/http"
	Const "plms_be/const"
	discussion_app "plms_be/internal/application/discussion"
	ViewModel "plms_be/viewModel"

	"github.com/gin-gonic/gin"
)

type DiscussionHandler struct {
	DiscussionAppService *discussion_app.DiscussionAppService
}

func NewDiscussionHandler(appService *discussion_app.DiscussionAppService) *DiscussionHandler {
	return &DiscussionHandler{DiscussionAppService: appService}
}

func (h *DiscussionHandler) GetAllDiscussion(c *gin.Context) {
	var response ViewModel.CommonResponse
	discussions, err := h.DiscussionAppService.GetAllDiscussion()
	if err != nil {
		response.Result = Const.FAIL
		response.Message = err.Error()
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response.Result = Const.SUCCESS
	response.Message = "Get all discussions successfully"
	response.Data = discussions
	c.JSON(http.StatusOK, response)
}