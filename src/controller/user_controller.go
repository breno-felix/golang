package controller

import (
	"github.com/breno-felix/golang/src/model/service"
	"github.com/gin-gonic/gin"
)

func NewUserControllerInterface(
	service service.UserDomainService,
) UserControllerInterface {
	return &userControllerInterface{
		service,
	}
}

type userControllerInterface struct {
	service service.UserDomainService
}

type UserControllerInterface interface {
	FindUserById(c *gin.Context)
	FindUserByEmail(c *gin.Context)

	CreateUser(c *gin.Context)
	UpdateUser(c *gin.Context)
	DeleteUser(c *gin.Context)
}
