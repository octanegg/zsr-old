package octane

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	database = "octane"
)

type client struct {
	DB *mongo.Client
}

// Client .
type Client interface {
	Ping() error
	Find(string, bson.M, func(*mongo.Cursor) (interface{}, error)) ([]interface{}, error)

	FindEvents(bson.M) (*Events, error)
	FindEventByID(string) (*Event, error)

	FindMatches(bson.M) (*Matches, error)
	FindMatchByID(string) (*Match, error)

	FindGames(bson.M) (*Games, error)
	FindGameByID(string) (*Game, error)

	FindPlayers(bson.M) (*Players, error)
	FindPlayerByID(string) (*Player, error)

	FindTeams(bson.M) (*Teams, error)
	FindTeamByID(string) (*Team, error)
}

// NewClient .
func NewClient(db *mongo.Client) Client {
	return &client{
		DB: db,
	}
}

func (c *client) Ping() error {
	return c.DB.Ping(context.TODO(), nil)
}

func (c *client) Find(collection string, filter bson.M, convert func(*mongo.Cursor) (interface{}, error)) ([]interface{}, error) {
	ctx := context.TODO()
	coll := c.DB.Database(database).Collection(collection)

	cursor, err := coll.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var res []interface{}
	for cursor.Next(context.TODO()) {
		i, err := convert(cursor)
		if err != nil {
			return nil, err
		}
		res = append(res, i)
	}

	return res, nil
}
