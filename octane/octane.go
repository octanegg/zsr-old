package octane

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type client struct {
	Octane            *mongo.Database
	EventsCollection  Collection
	MatchesCollection Collection
	GamesCollection   Collection
	PlayersCollection Collection
	TeamsCollection   Collection
	StatsCollection   Collection
	RecordsCollection Records
}

// Client .
type Client interface {
	Events() Collection
	Matches() Collection
	Games() Collection
	Players() Collection
	Teams() Collection
	Stats() Collection
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
		events = &collection{
			Collection: db.Collection("events"),
			Decode:     CursorToEvents,
		}
		matches = &collection{
			Collection: db.Collection("matches"),
			Decode:     CursorToMatches,
		}
		games = &collection{
			Collection: db.Collection("games"),
			Decode:     CursorToGames,
		}
		players = &collection{
			Collection: db.Collection("players"),
			Decode:     CursorToPlayers,
		}
		teams = &collection{
			Collection: db.Collection("teams"),
			Decode:     CursorToTeams,
		}
		stats = &collection{
			Collection: db.Collection("statlines"),
			Decode:     CursorToStats,
		}
	)

	return &client{
		Octane:            db,
		EventsCollection:  events,
		MatchesCollection: matches,
		GamesCollection:   games,
		PlayersCollection: players,
		TeamsCollection:   teams,
		StatsCollection:   stats,
		RecordsCollection: &records{games},
	}, nil
}

func (c *client) Events() Collection {
	return c.EventsCollection
}

func (c *client) Matches() Collection {
	return c.MatchesCollection
}

func (c *client) Games() Collection {
	return c.GamesCollection
}

func (c *client) Players() Collection {
	return c.PlayersCollection
}

func (c *client) Teams() Collection {
	return c.TeamsCollection
}

func (c *client) Stats() Collection {
	return c.StatsCollection
}

func (c *client) Records() Records {
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

// CursorToRecord .
func CursorToRecord(cursor *mongo.Cursor) (interface{}, error) {
	var record Record
	if err := cursor.Decode(&record); err != nil {
		return nil, err
	}
	return record, nil
}
