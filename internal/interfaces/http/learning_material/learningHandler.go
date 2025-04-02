package learning_http

import (
	"net/http"
	Const "plms_be/const"
	learning_app "plms_be/internal/application/learning"
	ViewModel "plms_be/viewModel"

	"github.com/gin-gonic/gin"
)

type LearningHandler struct {
	LearningAppService *learning_app.LearningAppService
}

func NewLearningHandler(appService *learning_app.LearningAppService) *LearningHandler {
	return &LearningHandler{LearningAppService: appService}
}

func (h *LearningHandler) GetAllLearning(c *gin.Context) {
	var response ViewModel.CommonResponse
	learningMaterials, err := h.LearningAppService.GetAllLearning()
	if err != nil {
		response.Result = Const.FAIL
		response.Message = err.Error()
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response.Result = Const.SUCCESS
	response.Message = "Get all learning materials successfully"
	response.Data = learningMaterials
	c.JSON(http.StatusOK, response)
}