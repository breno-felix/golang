package service

import (
	"github.com/breno-felix/golang/src/configuration/logger"
	"github.com/breno-felix/golang/src/configuration/rest_err"
	"github.com/breno-felix/golang/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) UpdateUser(
	userId string,
	userDomain model.UserDomainInterface,
) *rest_err.RestErr {
	logger.Info("Init updateUser service",
		zap.String("journey", "updateUser"),
	)

	err := ud.userRepository.UpdateUser(userId, userDomain)
	if err != nil {
		logger.Error("Error trying to call update user repository",
			err,
			zap.String("journey", "updateUser"),
		)
		return err
	}

	logger.Info("UpdateUser service executed successfully",
		zap.String("userId", userId),
		zap.String("journey", "updateUser"),
	)

	return nil
}
