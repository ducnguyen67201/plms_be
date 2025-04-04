package learning_http

import (
	"net/http"
	Const "plms_be/const"
	learning_app "plms_be/internal/application/learning"
	learning_domain "plms_be/internal/domain/learning"
	ViewModel "plms_be/viewModel"
	"strconv"

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

func (h *LearningHandler) GetLearningByID(c *gin.Context) {
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
	learningMaterial, err := h.LearningAppService.GetLearningByID(idInt64)
	if err != nil {
		response.Result = Const.FAIL
		response.Message = err.Error()
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response.Result = Const.SUCCESS
	response.Message = "Get learning material successfully"
	response.Data = learningMaterial
	c.JSON(http.StatusOK, response)
}

func (h *LearningHandler) SaveLearning(c *gin.Context) {
	var response ViewModel.CommonResponse
	var partialLearningMaterialInput learning_domain.PartialLearningMaterial

	if err := c.ShouldBindJSON(&partialLearningMaterialInput); err != nil {
		response.Result = Const.FAIL
		response.Message = err.Error()
		c.JSON(http.StatusBadRequest, response)
		return
	}

	err := h.LearningAppService.SaveLearning(&partialLearningMaterialInput)
	if err != nil {
		response.Result = Const.FAIL
		response.Message = err.Error()
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response.Result = Const.SUCCESS
	response.Message = "Save learning material successfully"
	c.JSON(http.StatusOK, response)
}