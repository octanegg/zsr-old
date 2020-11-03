package octane

import (
	"context"

	"github.com/octanegg/zsr/octane/collection"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type client struct {
	Octane            *mongo.Database
	EventsCollection  collection.Collection
	MatchesCollection collection.Collection
	GamesCollection   collection.Collection
	PlayersCollection collection.Collection
	TeamsCollection   collection.Collection
}

// Client .
type Client interface {
	Events() collection.Collection
	Matches() collection.Collection
	Games() collection.Collection
	Players() collection.Collection
	Teams() collection.Collection
}

// New .
func New(uri string) (Client, error) {
	c, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	if err := c.Ping(context.TODO(), nil); err != nil {
		return nil, err
	}

	db := c.Database("octane")

	return &client{
		Octane: db,
		EventsCollection: collection.New(
			db.Collection("events"),
			func(cursor *mongo.Cursor) (interface{}, error) {
				var event Event
				if err := cursor.Decode(&event); err != nil {
					return nil, err
				}
				return event, nil
			},
		),
		MatchesCollection: collection.New(
			db.Collection("matches"),
			func(cursor *mongo.Cursor) (interface{}, error) {
				var match Match
				if err := cursor.Decode(&match); err != nil {
					return nil, err
				}
				return match, nil
			},
		),
		GamesCollection: collection.New(
			db.Collection("games"),
			func(cursor *mongo.Cursor) (interface{}, error) {
				var game Game
				if err := cursor.Decode(&game); err != nil {
					return nil, err
				}
				return game, nil
			},
		),
		PlayersCollection: collection.New(
			db.Collection("players"),
			func(cursor *mongo.Cursor) (interface{}, error) {
				var player Player
				if err := cursor.Decode(&player); err != nil {
					return nil, err
				}
				return player, nil
			},
		),
		TeamsCollection: collection.New(
			db.Collection("teams"),
			func(cursor *mongo.Cursor) (interface{}, error) {
				var team Team
				if err := cursor.Decode(&team); err != nil {
					return nil, err
				}
				return team, nil
			},
		),
	}, nil
}

func (c *client) Events() collection.Collection {
	return c.EventsCollection
}

func (c *client) Matches() collection.Collection {
	return c.MatchesCollection
}

func (c *client) Games() collection.Collection {
	return c.GamesCollection
}

func (c *client) Players() collection.Collection {
	return c.PlayersCollection
}

func (c *client) Teams() collection.Collection {
	return c.TeamsCollection
}
