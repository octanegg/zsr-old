package octane

import (
	"context"

	"github.com/octanegg/zsr/octane/collection"
	"github.com/octanegg/zsr/octane/records"
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
	StatsCollection   collection.Collection
	RecordsCollection records.Records
}

// Client .
type Client interface {
	Events() collection.Collection
	Matches() collection.Collection
	Games() collection.Collection
	Players() collection.Collection
	Teams() collection.Collection
	Stats() collection.Collection
	Records() records.Records
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
			CursorToEvents,
		)
		matches = collection.New(
			db.Collection("matches"),
			CursorToMatches,
		)
		games = collection.New(
			db.Collection("games"),
			CursorToGames,
		)
		players = collection.New(
			db.Collection("players"),
			CursorToPlayers,
		)
		teams = collection.New(
			db.Collection("teams"),
			CursorToTeams,
		)
		stats = collection.New(
			db.Collection("statlines"),
			CursorToStats,
		)
	)

	return &client{
		Octane:            db,
		EventsCollection:  events,
		MatchesCollection: matches,
		GamesCollection:   games,
		PlayersCollection: players,
		TeamsCollection:   teams,
		StatsCollection:   stats,
		RecordsCollection: records.New(stats),
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

func (c *client) Stats() collection.Collection {
	return c.StatsCollection
}

func (c *client) Records() records.Records {
	return c.RecordsCollection
}

// CursorToEvents .
func CursorToEvents(cursor *mongo.Cursor) (interface{}, error) {
	var event Event
	if err := cursor.Decode(&event); err != nil {
		return nil, err
	}
	return event, nil
}

// CursorToMatches .
func CursorToMatches(cursor *mongo.Cursor) (interface{}, error) {
	var match Match
	if err := cursor.Decode(&match); err != nil {
		return nil, err
	}
	return match, nil
}

// CursorToGames .
func CursorToGames(cursor *mongo.Cursor) (interface{}, error) {
	var game Game
	if err := cursor.Decode(&game); err != nil {
		return nil, err
	}
	return game, nil
}

// CursorToPlayers .
func CursorToPlayers(cursor *mongo.Cursor) (interface{}, error) {
	var player Player
	if err := cursor.Decode(&player); err != nil {
		return nil, err
	}
	return player, nil
}

// CursorToTeams .
func CursorToTeams(cursor *mongo.Cursor) (interface{}, error) {
	var team Team
	if err := cursor.Decode(&team); err != nil {
		return nil, err
	}
	return team, nil
}

// CursorToStats .
func CursorToStats(cursor *mongo.Cursor) (interface{}, error) {
	var stats Stats
	if err := cursor.Decode(&stats); err != nil {
		return nil, err
	}
	return stats, nil
}
