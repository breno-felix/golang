package controller

import (
	"net/http"

	"github.com/breno-felix/golang/src/configuration/logger"
	"github.com/breno-felix/golang/src/configuration/validation"
	"github.com/breno-felix/golang/src/controller/model/request"
	"github.com/breno-felix/golang/src/model"
	"github.com/breno-felix/golang/src/view"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var (
	UserDomainInterface model.UserDomainInterface
)

func (uc *userControllerInterface) CreateUser(c *gin.Context) {
	logger.Info("Init CreateUser controller",
		zap.String("journey", "CreateUser"),
	)

	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error trying to validate user info",
			err,
			zap.String("journey", "CreateUser"),
		)

		errRest := validation.ValidateUserError(err)

		c.JSON(errRest.Code, errRest)
		return
	}

	userDomain := model.NewUserDomain(
		userRequest.Email,
		userRequest.Password,
		userRequest.Name,
		userRequest.Age,
	)

	domainResult, err := uc.service.CreateUserServices(userDomain)
	if err != nil {
		logger.Error("Error trying to CreateUser service",
			err,
			zap.String("journey", "CreateUser"),
		)

		c.JSON(err.Code, err)
		return
	}

	logger.Info("CreateUser controller executed successfully",
		zap.String("userId", domainResult.GetId()),
		zap.String("journey", "CreateUser"),
	)

	c.JSON(http.StatusCreated, view.ConvertDomainToResponse(domainResult))
}
