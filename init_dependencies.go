package main

import (
	"github.com/breno-felix/golang/src/controller"
	"github.com/breno-felix/golang/src/model/repository"
	"github.com/breno-felix/golang/src/model/service"
	"go.mongodb.org/mongo-driver/mongo"
)

func initDependencies(
	database *mongo.Database,
) controller.UserControllerInterface {
	repo := repository.NewUserRepository(database)
	service := service.NewUserDomainService(repo)
	return controller.NewUserControllerInterface(service)
}
