package controller

import (
	"net/http"

	"github.com/breno-felix/golang/src/configuration/logger"
	"github.com/breno-felix/golang/src/configuration/rest_err"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

func (uc *userControllerInterface) DeleteUser(c *gin.Context) {
	logger.Info("Init DeleteUser controller",
		zap.String("journey", "deleteUser"),
	)

	userId := c.Param("userId")
	if _, err := primitive.ObjectIDFromHex(userId); err != nil {
		logger.Error("Error trying to convert userId to ObjectID",
			err,
			zap.String("journey", "deleteUser"),
		)
		errRest := rest_err.NewBadRequestError("Invalid userId, must be a valid ObjectID")
		c.JSON(errRest.Code, errRest)
		return
	}

	err := uc.service.DeleteUser(userId)
	if err != nil {
		logger.Error("Error trying to call delete user services",
			err,
			zap.String("journey", "deleteUser"),
		)
		c.JSON(err.Code, err)
		return
	}

	logger.Info("DeleteUser controller executed successfully",
		zap.String("userId", userId),
		zap.String("journey", "deleteUser"),
	)

	c.Status(http.StatusNoContent)
}
