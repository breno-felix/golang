package controller

import (
	"net/http"

	"github.com/breno-felix/golang/src/configuration/logger"
	"github.com/breno-felix/golang/src/configuration/rest_err"
	"github.com/breno-felix/golang/src/configuration/validation"
	"github.com/breno-felix/golang/src/controller/model/request"
	"github.com/breno-felix/golang/src/model"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

func (uc *userControllerInterface) UpdateUser(c *gin.Context) {
	logger.Info("Init UpdateUser controller",
		zap.String("journey", "updateUser"),
	)

	var userUpdateRequest request.UserUpdateRequest
	if err := c.ShouldBindJSON(&userUpdateRequest); err != nil {
		logger.Error("Error trying to validate user info",
			err,
			zap.String("journey", "updateUser"),
		)
		errRest := validation.ValidateUserError(err)

		c.JSON(errRest.Code, errRest)
		return
	}

	userId := c.Param("userId")
	if _, err := primitive.ObjectIDFromHex(userId); err != nil {
		logger.Error("Error trying to convert userId to ObjectID",
			err,
			zap.String("journey", "updateUser"),
		)
		errRest := rest_err.NewBadRequestError("Invalid userId, must be a valid ObjectID")
		c.JSON(errRest.Code, errRest)
		return
	}

	domain := model.NewUserUpdateDomain(
		userUpdateRequest.Name,
		userUpdateRequest.Age,
	)

	err := uc.service.UpdateUser(userId, domain)
	if err != nil {
		logger.Error("Error trying to call update user services",
			err,
			zap.String("journey", "updateUser"),
		)
		c.JSON(err.Code, err)
		return
	}

	logger.Info("UpdateUser controller executed successfully",
		zap.String("userId", userId),
		zap.String("journey", "updateUser"),
	)

	c.Status(http.StatusNoContent)
}
