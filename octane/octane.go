package octane

import (
	"context"
	"reflect"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	database = "octane"
)

type client struct {
	DB *mongo.Client
}

// ObjectID .
type ObjectID struct {
	ID string `json:"_id"`
}

// Client .
type Client interface {
	Ping() error
	Find(string, bson.M, func(*mongo.Cursor) (interface{}, error)) ([]interface{}, error)
	Insert(string, interface{}) (interface{}, error)
	Replace(string, bson.M, interface{}) (interface{}, error)

	FindEvents(bson.M) (*Events, error)
	FindEventByID(*primitive.ObjectID) (*Event, error)
	InsertEvent(*Event) (*ObjectID, error)
	UpdateEvent(*primitive.ObjectID, *Event) (*ObjectID, error)

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

func (c *client) Insert(collection string, document interface{}) (interface{}, error) {
	ctx := context.TODO()
	coll := c.DB.Database(database).Collection(collection)

	res, err := coll.InsertOne(ctx, document)
	if err != nil {
		return "", err
	}

	return res.InsertedID, nil
}

func (c *client) Replace(collection string, filter bson.M, update interface{}) (interface{}, error) {
	ctx := context.TODO()
	coll := c.DB.Database(database).Collection(collection)

	res, err := coll.ReplaceOne(ctx, filter, update)
	if err != nil {
		return "", err
	}

	return res.UpsertedID, nil
}

func updateFields(x, y reflect.Value) interface{} {
	for i := 0; i < x.NumField(); i++ {
		if val := reflect.Value(y.Field(i)); val.IsValid() {
			x.Field(i).Set(val)
		}
	}
	return x.Interface()
}
