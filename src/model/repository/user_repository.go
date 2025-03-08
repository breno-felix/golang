package repository

import (
	"github.com/breno-felix/golang/src/configuration/rest_err"
	"github.com/breno-felix/golang/src/model"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	MONGODB_USER_COLLECTION = "MONGODB_USER_COLLECTION"
)

func NewUserRepository(
	database *mongo.Database,
) UserRepository {
	return &userRepository{
		databaseConnection: database,
	}
}

type userRepository struct {
	databaseConnection *mongo.Database
}

type UserRepository interface {
	CreateUser(
		userDomain model.UserDomainInterface,
	) (model.UserDomainInterface, *rest_err.RestErr)

	UpdateUser(
		userId string,
		userDomain model.UserDomainInterface,
	) *rest_err.RestErr

	FindUserByEmail(
		email string,
	) (model.UserDomainInterface, *rest_err.RestErr)

	FindUserById(
		userId string,
	) (model.UserDomainInterface, *rest_err.RestErr)

	DeleteUser(
		userId string,
	) *rest_err.RestErr
}
