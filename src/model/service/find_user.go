package service

import (
	"github.com/breno-felix/golang/src/configuration/logger"
	"github.com/breno-felix/golang/src/configuration/rest_err"
	"github.com/breno-felix/golang/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) FindUserByIdServices(
	userId string,
) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init FindUserByIdServices services",
		zap.String("journey", "findUser"),
	)
	return ud.userRepository.FindUserById(userId)
}

func (ud *userDomainService) FindUserByEmailServices(
	email string,
) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init FindUserByEmailServices services",
		zap.String("journey", "findUser"),
	)
	return ud.userRepository.FindUserByEmail(email)
}
