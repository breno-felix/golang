package service

import (
	"github.com/breno-felix/golang/src/configuration/logger"
	"github.com/breno-felix/golang/src/configuration/rest_err"
	"go.uber.org/zap"
)

func (ud *userDomainService) DeleteUser(
	userId string,
) *rest_err.RestErr {
	logger.Info("Init deleteUser service",
		zap.String("journey", "deleteUser"),
	)

	err := ud.userRepository.DeleteUser(userId)
	if err != nil {
		logger.Error("Error trying to call delete user repository",
			err,
			zap.String("journey", "deleteUser"),
		)
		return err
	}

	logger.Info("DeleteUser service executed successfully",
		zap.String("userId", userId),
		zap.String("journey", "deleteUser"),
	)

	return nil
}
