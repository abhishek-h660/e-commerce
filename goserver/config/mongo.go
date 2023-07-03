package config

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DefaultCtx() context.Context {
	return context.Background()
}

func Client() *mongo.Client {
	uri := os.Getenv("MONGO_URI")
	client, err := mongo.Connect(DefaultCtx(), options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}
	if err := client.Ping(DefaultCtx(), nil); err != nil {
		log.Fatal(err)
	}

	return client
}

func DefaultDB() *mongo.Database {
	return Client().Database("shri-krishnagan")
}
