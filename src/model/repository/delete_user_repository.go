package repository

import (
	"context"
	"os"

	"github.com/breno-felix/golang/src/configuration/logger"
	"github.com/breno-felix/golang/src/configuration/rest_err"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

func (ur *userRepository) DeleteUser(
	userId string,
) *rest_err.RestErr {
	logger.Info("Init deleteUser repository",
		zap.String("journey", "deleteUser"),
	)

	collection_name := os.Getenv(MONGODB_USER_COLLECTION)
	collection := ur.databaseConnection.Collection(collection_name)

	userIdHex, _ := primitive.ObjectIDFromHex(userId)
	filter := bson.M{"_id": userIdHex}
	_, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		logger.Error("Error trying to delete user",
			err,
			zap.String("journey", "deleteUser"),
		)
		return rest_err.NewInternalServerError(err.Error())
	}

	logger.Info("DeleteUser repository executed successfully",
		zap.String("userId", userId),
		zap.String("journey", "deleteUser"),
	)

	return nil
}
