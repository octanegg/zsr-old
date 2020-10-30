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
)

type client struct {
	DB *mongo.Client
}

// Client .
type Client interface {
	Ping() error

	FindEvents(*FindContext) (*Events, error)
	FindMatches(*FindContext) (*Matches, error)
	FindGames(*FindContext) (*Games, error)
	FindPlayers(*FindContext) (*Players, error)
	FindTeams(*FindContext) (*Teams, error)

	FindEvent(bson.M) (*Event, error)
	FindMatch(bson.M) (*Match, error)
	FindGame(bson.M) (*Game, error)
	FindPlayer(bson.M) (*Player, error)
	FindTeam(bson.M) (*Team, error)

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

// FindContext .
type FindContext struct {
	Filter     bson.M
	Sort       bson.M
	Pagination *Pagination
}

// Pagination .
type Pagination struct {
	Page     int64 `json:"page"`
	PerPage  int64 `json:"perPage"`
	PageSize int   `json:"pageSize"`
}

// NewFindContext .
func NewFindContext(filter bson.M, sort bson.M, pagination *Pagination) *FindContext {
	return &FindContext{
		Filter:     filter,
		Sort:       sort,
		Pagination: pagination,
	}
}
