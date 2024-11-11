package controller

import (
	"github.com/gin-gonic/gin"
	"go-ecommerce-be/internal/service"
	"net/http"
)

type UserController struct {
	userService *service.UserService
}

func NewUserController() *UserController  {
	return &UserController{
		userService: service.NewUserService(),
	}
}

func (uc *UserController) GetUserById(c *gin.Context)  {
	c.JSON(http.StatusOK, gin.H{
		"message": "Suceess!",
		"name": uc.userService.GetInfoUser(),
	})
}
