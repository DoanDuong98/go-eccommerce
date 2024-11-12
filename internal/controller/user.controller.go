package controller

import (
	"github.com/gin-gonic/gin"
	"go-ecommerce-be/internal/service"
	"go-ecommerce-be/pkg/response"
)

type UserController struct {
	userService *service.UserService
}

func NewUserController() *UserController {
	return &UserController{
		userService: service.NewUserService(),
	}
}

func (uc *UserController) GetUserById(c *gin.Context) {
	response.SuccessResponse(c, 20001, []string{})
}
