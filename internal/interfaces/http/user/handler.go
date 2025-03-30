package user_http

import (
	"net/http"
	user_app "plms_be/internal/application/user"
	user_domain "plms_be/internal/domain/user"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	AppService *user_app.UserAppService
}

func NewHandler(appService *user_app.UserAppService) *Handler { 
	return &Handler{AppService: appService}
}

func (h *Handler) RegisterUser(c *gin.Context) { 
	var input user_domain.User

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	user, err := h.AppService.Register(input)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to register user"})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *Handler) LoginUser(c *gin.Context) {
	var input struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	user, err := h.AppService.Login(input.Username, input.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to login user"})
		return
	}

	c.JSON(http.StatusOK, user)
}