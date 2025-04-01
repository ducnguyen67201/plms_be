package problem_http

import (
	"net/http"
	Const "plms_be/const"
	problem_app "plms_be/internal/application/problem"
	ViewModel "plms_be/viewModel"

	"github.com/gin-gonic/gin"
)

type ProblemHandler struct {
	ProblemService *problem_app.ProblemAppService
}

func NewProblemHandler(problemService *problem_app.ProblemAppService) *ProblemHandler {
	return &ProblemHandler{ProblemService: problemService}
}

func (h *ProblemHandler) GetAllProblem(c *gin.Context) {
	var response ViewModel.CommonResponse
	problems, err := h.ProblemService.GetAllProblem()
	if err != nil {
		response.Result = Const.FAIL
		response.Message = err.Error()
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response.Result = Const.SUCCESS
	response.Message = "Get all problems successfully"
	response.Data = problems
	c.JSON(http.StatusOK, response)
}