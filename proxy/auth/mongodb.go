package auth

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const uri = "mongodb://localhost:27017"

var (
	Db *mongo.Database
)

func init() {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)
	client, err := mongo.Connect(context.Background(), opts)

	if err != nil {
		panic(err)
	}

	if err = client.Ping(context.Background(), nil); err != nil {
		panic(err)
	}

	Db = client.Database("yqy_sys")
}
