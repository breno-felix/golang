package main

import (
	"context"
	"log"

	"github.com/breno-felix/golang/src/configuration/database/mongodb"
	"github.com/breno-felix/golang/src/configuration/logger"
	"github.com/breno-felix/golang/src/controller/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	logger.Info("Starting server...")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	database, err := mongodb.NewMongoDBConnection(context.Background())
	if err != nil {
		log.Fatalf(
			"Error try to connect to database, error=%s \n",
			err.Error())
		return
	}

	userController := initDependencies(database)

	router := gin.Default()

	routes.InitRoutes(router, userController)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
