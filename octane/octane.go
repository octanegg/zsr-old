package octane

import (
	"context"
	"reflect"

	"github.com/octanegg/core/internal/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	database = "octane"
)

type client struct {
	DB *mongo.Client
}

// ObjectID .
type ObjectID struct {
	ID string `json:"id"`
}

// Data .
type Data struct {
	Data []interface{} `json:"data"`
	*Pagination
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
	Find(string, bson.M, *Pagination, *Sort, func(*mongo.Cursor) (interface{}, error)) ([]interface{}, error)
	Insert(string, interface{}) (interface{}, error)
	Update(string, bson.M, bson.M) (interface{}, error)
	Replace(string, *primitive.ObjectID, interface{}) (interface{}, error)
	Delete(string, *primitive.ObjectID) (int64, error)

	FindEvents(bson.M, *Pagination, *Sort) (*Data, error)
	FindMatches(bson.M, *Pagination, *Sort) (*Data, error)
	FindGames(bson.M, *Pagination, *Sort) (*Data, error)
	FindPlayers(bson.M, *Pagination, *Sort) (*Data, error)
	FindTeams(bson.M, *Pagination, *Sort) (*Data, error)

	FindEvent(*primitive.ObjectID) (*Event, error)
	FindMatch(*primitive.ObjectID) (*Match, error)
	FindGame(*primitive.ObjectID) (*Game, error)
	FindPlayer(*primitive.ObjectID) (*Player, error)
	FindTeam(*primitive.ObjectID) (*Team, error)

	InsertEvent(*Event) (*ObjectID, error)
	InsertMatch(*Match) (*ObjectID, error)
	InsertGame(*Game) (*ObjectID, error)
	InsertPlayer(*Player) (*ObjectID, error)
	InsertTeam(*Team) (*ObjectID, error)

	UpdateEvent(*primitive.ObjectID, *Event) (*ObjectID, error)
	UpdateMatch(*primitive.ObjectID, *Match) (*ObjectID, error)
	UpdateGame(*primitive.ObjectID, *Game) (*ObjectID, error)
	UpdatePlayer(*primitive.ObjectID, *Player) (*ObjectID, error)
	UpdateTeam(*primitive.ObjectID, *Team) (*ObjectID, error)

	DeleteEvent(*primitive.ObjectID) (int64, error)
	DeleteMatch(*primitive.ObjectID) (int64, error)
	DeleteGame(*primitive.ObjectID) (int64, error)
	DeletePlayer(*primitive.ObjectID) (int64, error)
	DeleteTeam(*primitive.ObjectID) (int64, error)
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

func (c *client) Find(collection string, filter bson.M, pagination *Pagination, sort *Sort, convert func(*mongo.Cursor) (interface{}, error)) ([]interface{}, error) {
	ctx := context.TODO()
	coll := c.DB.Database(database).Collection(collection)

	opts := options.Find()
	if pagination != nil {
		opts.SetSkip((pagination.Page - 1) * pagination.PerPage)
		opts.SetLimit(pagination.PerPage)
	}

	if sort != nil {
		opts.SetSort(bson.M{sort.Field: sort.Order})
	}

	cursor, err := coll.Find(ctx, filter, opts)
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
		return nil, err
	}

	return res.InsertedID, nil
}

func (c *client) Replace(collection string, oid *primitive.ObjectID, update interface{}) (interface{}, error) {
	ctx := context.TODO()
	coll := c.DB.Database(database).Collection(collection)

	res, err := coll.ReplaceOne(ctx, bson.M{config.KeyID: oid}, update)
	if err != nil {
		return nil, err
	}

	return res.UpsertedID, nil
}

func (c *client) Update(collection string, filter, update bson.M) (interface{}, error) {
	ctx := context.TODO()
	coll := c.DB.Database(database).Collection(collection)

	res, err := coll.UpdateMany(ctx, filter, update)
	if err != nil {
		return nil, err
	}

	return res.UpsertedID, nil
}

func (c *client) Delete(collection string, oid *primitive.ObjectID) (int64, error) {
	ctx := context.TODO()
	coll := c.DB.Database(database).Collection(collection)

	res, err := coll.DeleteOne(ctx, bson.M{config.KeyID: oid})
	if err != nil {
		return 0, err
	}

	return res.DeletedCount, nil
}

func updateFields(x, y reflect.Value) interface{} {
	for i := 0; i < x.NumField(); i++ {
		if val := reflect.Value(y.Field(i)); val.IsValid() {
			x.Field(i).Set(val)
		}
	}
	return x.Interface()
}
