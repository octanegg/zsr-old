package octane

import (
	"context"

	"github.com/octanegg/zsr/octane/collection"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type client struct {
	Octane              *mongo.Database
	EventsCollection    collection.Collection
	MatchesCollection   collection.Collection
	GamesCollection     collection.Collection
	PlayersCollection   collection.Collection
	TeamsCollection     collection.Collection
	StatlinesCollection collection.Collection
	RecordsCollection   collection.Collection
}

// Client .
type Client interface {
	Events() collection.Collection
	Matches() collection.Collection
	Games() collection.Collection
	Players() collection.Collection
	Teams() collection.Collection
	Statlines() collection.Collection
	Records() collection.Collection
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

	var (
		db     = c.Database("octane")
		events = collection.New(
			db.Collection("events"),
			toEvents,
		)
		matches = collection.New(
			db.Collection("matches"),
			toMatches,
		)
		games = collection.New(
			db.Collection("games"),
			toGames,
		)
		players = collection.New(
			db.Collection("players"),
			toPlayers,
		)
		teams = collection.New(
			db.Collection("teams"),
			toTeams,
		)
		statlines = collection.New(
			db.Collection("statlines"),
			toStatlines,
		)
		records = collection.New(
			db.Collection("records"),
			toStatlines,
		)
	)

	return &client{
		Octane:              db,
		EventsCollection:    events,
		MatchesCollection:   matches,
		GamesCollection:     games,
		PlayersCollection:   players,
		TeamsCollection:     teams,
		StatlinesCollection: statlines,
		RecordsCollection: records,
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

func (c *client) Statlines() collection.Collection {
	return c.StatlinesCollection
}

func (c *client) Records() collection.Collection {
	return c.RecordsCollection
}

func toEvents(cursor *mongo.Cursor) (interface{}, error) {
	var event Event
	if err := cursor.Decode(&event); err != nil {
		return nil, err
	}
	return event, nil
}

func toMatches(cursor *mongo.Cursor) (interface{}, error) {
	var match Match
	if err := cursor.Decode(&match); err != nil {
		return nil, err
	}
	return match, nil
}

func toGames(cursor *mongo.Cursor) (interface{}, error) {
	var game Game
	if err := cursor.Decode(&game); err != nil {
		return nil, err
	}
	return game, nil
}

func toPlayers(cursor *mongo.Cursor) (interface{}, error) {
	var player Player
	if err := cursor.Decode(&player); err != nil {
		return nil, err
	}
	return player, nil
}

func toTeams(cursor *mongo.Cursor) (interface{}, error) {
	var team Team
	if err := cursor.Decode(&team); err != nil {
		return nil, err
	}
	return team, nil
}

func toStatlines(cursor *mongo.Cursor) (interface{}, error) {
	var statline Statline
	if err := cursor.Decode(&statline); err != nil {
		return nil, err
	}
	return statline, nil
}


func toRecords(cursor *mongo.Cursor) (interface{}, error) {
	var record Record
	if err := cursor.Decode(&record); err != nil {
		return nil, err
	}
	return record, nil
}
