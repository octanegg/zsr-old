package octane

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	// Database .
	Database = "octane"

	// CollectionEvents .
	CollectionEvents = "events"

	// CollectionMatches .
	CollectionMatches = "matches"

	// CollectionGames .
	CollectionGames = "games"

	// CollectionPlayers .
	CollectionPlayers = "players"

	// CollectionTeams .
	CollectionTeams = "teams"

	// CollectionUsers .
	CollectionUsers = "users"
)

type client struct {
	DB *mongo.Client
}

// Pagination .
type Pagination struct {
	Page     int64 `json:"page"`
	PerPage  int64 `json:"perPage"`
	PageSize int   `json:"pageSize"`
}

// Sort .
type Sort struct {
	Field string
	Order int
}

// Client .
type Client interface {
	Ping() error

	FindEvents(bson.M, *Pagination, *Sort) (*Events, error)
	FindMatches(bson.M, *Pagination, *Sort) (*Matches, error)
	FindGames(bson.M, *Pagination, *Sort) (*Games, error)
	FindPlayers(bson.M, *Pagination, *Sort) (*Players, error)
	FindTeams(bson.M, *Pagination, *Sort) (*Teams, error)

	FindEvent(*primitive.ObjectID) (*Event, error)
	FindMatch(*primitive.ObjectID) (*Match, error)
	FindGame(*primitive.ObjectID) (*Game, error)
	FindPlayer(*primitive.ObjectID) (*Player, error)
	FindTeam(*primitive.ObjectID) (*Team, error)

	InsertEvent(interface{}) (*primitive.ObjectID, error)
	InsertMatch(interface{}) (*primitive.ObjectID, error)
	InsertGame(interface{}) (*primitive.ObjectID, error)
	InsertPlayer(interface{}) (*primitive.ObjectID, error)
	InsertTeam(interface{}) (*primitive.ObjectID, error)

	InsertEvents([]interface{}) ([]interface{}, error)
	InsertMatches([]interface{}) ([]interface{}, error)
	InsertGames([]interface{}) ([]interface{}, error)
	InsertPlayers([]interface{}) ([]interface{}, error)
	InsertTeams([]interface{}) ([]interface{}, error)

	DeleteEvent(bson.M) (int64, error)
	DeleteMatch(bson.M) (int64, error)
	DeleteGame(bson.M) (int64, error)
	DeletePlayer(bson.M) (int64, error)
	DeleteTeam(bson.M) (int64, error)
}

// New .
func New(uri string) (Client, error) {
	db, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	client := &client{db}
	if err := client.Ping(); err != nil {
		return nil, err
	}

	return client, nil
}

func (c *client) Ping() error {
	return c.DB.Ping(context.TODO(), nil)
}
