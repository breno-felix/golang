package service

import (
	"github.com/breno-felix/golang/src/configuration/logger"
	"github.com/breno-felix/golang/src/configuration/rest_err"
	"github.com/breno-felix/golang/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) CreateUserServices(
	userDomain model.UserDomainInterface,
) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init CreateUser service",
		zap.String("journey", "createUser"),
	)

	user, _ := ud.FindUserByEmailServices(userDomain.GetEmail())
	if user != nil {
		logger.Error("Error trying to create user, email already exists",
			nil,
			zap.String("email", userDomain.GetEmail()),
			zap.String("journey", "createUser"),
		)
		return nil, rest_err.NewBadRequestError("Email is already registered in another account")
	}

	userDomain.EncryptPassword()

	userDomainRepository, err := ud.userRepository.CreateUser(userDomain)
	if err != nil {
		logger.Error("Error trying call CreateUser repository",
			err,
			zap.String("journey", "createUser"),
		)
		return nil, err
	}

	logger.Info("CreateUser service executed successfully",
		zap.String("userId", userDomainRepository.GetId()),
		zap.String("journey", "createUser"),
	)
	return userDomainRepository, nil
}
