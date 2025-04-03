package problem_http

import (
	"net/http"
	Const "plms_be/const"
	problem_app "plms_be/internal/application/problem"
	problem_domain "plms_be/internal/domain/problem"
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

func (h *ProblemHandler) GetProblemById(c *gin.Context) {
	var response ViewModel.CommonResponse
	id := c.Param("id")
	problem, err := h.ProblemService.GetProblemById(id)
	if err != nil {
		response.Result = Const.FAIL
		response.Message = err.Error()
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response.Result = Const.SUCCESS
	response.Message = "Get problem successfully"
	response.Data = problem
	c.JSON(http.StatusOK, response)
}

func (h *ProblemHandler) SaveProblem(c *gin.Context) {
	var req problem_domain.PartialProblemUpdate
	var response ViewModel.CommonResponse

	if err := c.ShouldBindJSON(&req); err != nil {
		response.Result = Const.FAIL
		response.Message = err.Error()
		c.JSON(http.StatusBadRequest, response)
		return
	}

	err := h.ProblemService.SaveProblem(&req)
	if err != nil {
		response.Result = Const.FAIL
		response.Message = err.Error()
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response.Result = Const.SUCCESS
	response.Message = "Save problem successfully"
	c.JSON(http.StatusOK, response)
}

func (h *ProblemHandler) SaveTestCase(c *gin.Context) {
	var req problem_domain.PartialTestCaseUpdate
	var response ViewModel.CommonResponse

	if err := c.ShouldBindJSON(&req); err != nil {
		response.Result = Const.FAIL
		response.Message = err.Error()
		c.JSON(http.StatusBadRequest, response)
		return
	}

	err := h.ProblemService.SaveTestCase(&req)
	if err != nil {
		response.Result = Const.FAIL
		response.Message = err.Error()
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response.Result = Const.SUCCESS
	response.Message = "Save test case successfully"
	c.JSON(http.StatusOK, response)
}