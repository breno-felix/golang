package mongodb

import (
	"context"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	MONGODB_URL             = "MONGODB_URL"
	MONGODB_GOLANG_DATABASE = "MONGODB_GOLANG_DATABASE"
)

func NewMongoDBConnection(
	ctx context.Context,
) (*mongo.Database, error) {
	mongodb_url := os.Getenv(MONGODB_URL)
	mongodb_database := os.Getenv(MONGODB_GOLANG_DATABASE)

	client, err := mongo.Connect(
		ctx,
		options.Client().ApplyURI(mongodb_url))
	if err != nil {
		return nil, err
	}
	if err := client.Ping(ctx, nil); err != nil {
		return nil, err
	}

	return client.Database(mongodb_database), nil
}
