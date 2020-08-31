package octane

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type client struct {
	DB *mongo.Client
}

type Client interface {
	Ping() error
}

func NewClient(db *mongo.Client) Client {
	return &client{
		DB: db,
	}
}

func (c *client) Ping() error {
	return c.DB.Ping(context.TODO(), nil)
}
