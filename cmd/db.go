package main

import (
	"context"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	envURI = "DB_URI"
)

func connect() (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI(os.Getenv(envURI))
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, err
	}

	return client, nil
}
