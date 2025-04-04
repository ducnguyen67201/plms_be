package discussion_http

import (
	"net/http"
	Const "plms_be/const"
	discussion_app "plms_be/internal/application/discussion"
	discussion_domain "plms_be/internal/domain/discussion"
	ViewModel "plms_be/viewModel"
	"strconv"

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

func (h *DiscussionHandler) GetDiscussionById(c *gin.Context) {
	var response ViewModel.CommonResponse
	id := c.Param("id")
	idInt64, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		var response ViewModel.CommonResponse
		response.Result = Const.FAIL
		response.Message = "Invalid ID format"
		c.JSON(http.StatusBadRequest, response)
		return
	}
	discussion, err := h.DiscussionAppService.GetDiscussionById(idInt64)
	if err != nil {
		response.Result = Const.FAIL
		response.Message = err.Error()
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response.Result = Const.SUCCESS
	response.Message = "Get discussion by id successfully"
	response.Data = discussion
	c.JSON(http.StatusOK, response)
}

func (h *DiscussionHandler) SaveDiscussion(c *gin.Context) {
	var response ViewModel.CommonResponse
	var partialUpdateInput discussion_domain.PartialDiscussionUpdate

	if err := c.ShouldBindJSON(&partialUpdateInput); err != nil {
		response.Result = Const.FAIL
		response.Message = err.Error()
		c.JSON(http.StatusBadRequest, response)
		return
	}

	err := h.DiscussionAppService.SaveDiscussion(&partialUpdateInput)
	if err != nil {
		response.Result = Const.FAIL
		response.Message = err.Error()
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response.Result = Const.SUCCESS
	response.Message = "Save discussion successfully"

	c.JSON(http.StatusOK, response)
}