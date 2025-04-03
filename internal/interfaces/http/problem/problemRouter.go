package problem_http

import (
	problem_app "plms_be/internal/application/problem"

	"github.com/gin-gonic/gin"
)

func RegisterProblemRoutes(router *gin.Engine, appService *problem_app.ProblemAppService) {
	h := NewProblemHandler(appService)

	ProblemGroup := router.Group("/problem")
	{
		ProblemGroup.POST("/all", h.GetAllProblem)
		ProblemGroup.POST("/:id", h.GetProblemById)
		ProblemGroup.POST("/save", h.SaveProblem)
	}

	TestCaseGroup := router.Group("/testcase")
	{
		TestCaseGroup.POST("/save", h.SaveTestCase)
	}
}