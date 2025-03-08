package repository

import (
	"context"
	"fmt"
	"os"

	"github.com/breno-felix/golang/src/configuration/logger"
	"github.com/breno-felix/golang/src/configuration/rest_err"
	"github.com/breno-felix/golang/src/model"
	"github.com/breno-felix/golang/src/model/repository/entity"
	"github.com/breno-felix/golang/src/model/repository/entity/converter"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

func (ur *userRepository) FindUserByEmail(
	email string,
) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init FindUserByEmail repository",
		zap.String("journey", "findUserByEmail"),
	)

	collection_name := os.Getenv(MONGODB_USER_COLLECTION)
	collection := ur.databaseConnection.Collection(collection_name)

	filter := bson.M{"email": email}

	userEntity := &entity.UserEntity{}

	err := collection.FindOne(
		context.Background(),
		filter,
	).Decode(userEntity)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			errorMessage := fmt.Sprintf(
				"User not found with email: %s", email,
			)
			logger.Error(errorMessage,
				err,
				zap.String("journey", "findUserByEmail"),
			)
			return nil, rest_err.NewNotFoundError(errorMessage)
		}
		errorMessage := "Error trying to find user by email"
		logger.Error(errorMessage,
			err,
			zap.String("journey", "findUserByEmail"),
		)
		return nil, rest_err.NewInternalServerError(errorMessage)
	}

	logger.Info("FindUserByEmail repository executed successfully",
		zap.String("userId", userEntity.Id.Hex()),
		zap.String("email", email),
		zap.String("journey", "findUserByEmail"),
	)

	return converter.ConvertEntityToDomain(*userEntity), nil
}

func (ur *userRepository) FindUserById(
	userId string,
) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init FindUserById repository",
		zap.String("journey", "findUserById"),
	)

	collection_name := os.Getenv(MONGODB_USER_COLLECTION)
	collection := ur.databaseConnection.Collection(collection_name)

	objectId, _ := primitive.ObjectIDFromHex(userId)
	filter := bson.M{"_id": objectId}

	userEntity := &entity.UserEntity{}

	err := collection.FindOne(
		context.Background(),
		filter,
	).Decode(userEntity)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			errorMessage := fmt.Sprintf(
				"User not found with id: %s", userId,
			)
			logger.Error(errorMessage,
				err,
				zap.String("journey", "findUserById"),
			)
			return nil, rest_err.NewNotFoundError(errorMessage)
		}
		errorMessage := "Error trying to find user by id"
		logger.Error(errorMessage,
			err,
			zap.String("journey", "findUserById"),
		)
		return nil, rest_err.NewInternalServerError(errorMessage)
	}

	logger.Info("FindUserById repository executed successfully",
		zap.String("userId", userId),
		zap.String("journey", "findUserById"),
	)

	return converter.ConvertEntityToDomain(*userEntity), nil
}
