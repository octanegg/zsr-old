package main

import (
	"context"
	"os"

	"github.com/octanegg/core/internal/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func connect() (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI(os.Getenv(config.EnvURI))
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, err
	}

	return client, nil
}
