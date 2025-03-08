package repository

import (
	"context"
	"os"

	"github.com/breno-felix/golang/src/configuration/logger"
	"github.com/breno-felix/golang/src/configuration/rest_err"
	"github.com/breno-felix/golang/src/model"
	"github.com/breno-felix/golang/src/model/repository/entity/converter"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

func (ur *userRepository) CreateUser(
	userDomain model.UserDomainInterface,
) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init CreateUser repository",
		zap.String("journey", "createUser"),
	)

	collection_name := os.Getenv(MONGODB_USER_COLLECTION)

	collection := ur.databaseConnection.Collection(collection_name)

	value := converter.ConvertDomainToEntity(userDomain)

	result, err := collection.InsertOne(context.Background(), value)
	if err != nil {
		logger.Error("Error trying to create user",
			err,
			zap.String("journey", "createUser"),
		)
		return nil, rest_err.NewInternalServerError(err.Error())
	}

	value.Id = result.InsertedID.(primitive.ObjectID)

	logger.Info("CreateUser repository executed successfully",
		zap.String("userId", value.Id.Hex()),
		zap.String("journey", "createUser"),
	)

	return converter.ConvertEntityToDomain(*value), nil
}
