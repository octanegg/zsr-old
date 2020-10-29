package octane

import (
	"context"
	"reflect"

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
	InsertOne(string, interface{}) (*primitive.ObjectID, error)
	InsertMany(string, []interface{}) ([]interface{}, error)
	Update(string, bson.M, bson.M) (int64, error)
	Replace(string, *primitive.ObjectID, interface{}) error
	Delete(string, bson.M) (int64, error)

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

	InsertEvent(*Event) (*primitive.ObjectID, error)
	InsertMatch(*Match) (*primitive.ObjectID, error)
	InsertGame(*Game) (*primitive.ObjectID, error)
	InsertPlayer(*Player) (*primitive.ObjectID, error)
	InsertTeam(*Team) (*primitive.ObjectID, error)


	InsertEvents([]interface{}) ([]interface{}, error)
	InsertMatches([]interface{}) ([]interface{}, error)
	InsertGames([]interface{}) ([]interface{}, error)
	InsertPlayers([]interface{}) ([]interface{}, error)
	InsertTeams([]interface{}) ([]interface{}, error)

	UpdateEvents(bson.M, bson.M) (int64, error)
	UpdateMatches(bson.M, bson.M) (int64, error)
	UpdateGames(bson.M, bson.M) (int64, error)
	UpdatePlayers(bson.M, bson.M) (int64, error)
	UpdateTeams(bson.M, bson.M) (int64, error)

	ReplaceEvent(*primitive.ObjectID, *Event) (*primitive.ObjectID, error)
	ReplaceMatch(*primitive.ObjectID, *Match) (*primitive.ObjectID, error)
	ReplaceGame(*primitive.ObjectID, *Game) (*primitive.ObjectID, error)
	ReplacePlayer(*primitive.ObjectID, *Player) (*primitive.ObjectID, error)
	ReplaceTeam(*primitive.ObjectID, *Team) (*primitive.ObjectID, error)

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

func (c *client) Find(collection string, filter bson.M, pagination *Pagination, sort *Sort, convert func(*mongo.Cursor) (interface{}, error)) ([]interface{}, error) {
	ctx := context.TODO()
	coll := c.DB.Database(Database).Collection(collection)

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

func (c *client) InsertOne(collection string, document interface{}) (*primitive.ObjectID, error) {
	ctx := context.TODO()
	coll := c.DB.Database(Database).Collection(collection)

	res, err := coll.InsertOne(ctx, document)
	if err != nil {
		return nil, err
	}

	newID := res.InsertedID.(primitive.ObjectID)

	return &newID, nil
}

func (c *client) InsertMany(collection string, documents []interface{}) ([]interface{}, error) {
	ctx := context.TODO()
	coll := c.DB.Database(Database).Collection(collection)

	res, err := coll.InsertMany(ctx, documents)
	if err != nil {
		return nil, err
	}

	return res.InsertedIDs, nil
}

func (c *client) Replace(collection string, oid *primitive.ObjectID, update interface{}) error {
	ctx := context.TODO()
	coll := c.DB.Database(Database).Collection(collection)

	_, err := coll.ReplaceOne(ctx, bson.M{"_id": oid}, update)
	if err != nil {
		return err
	}

	return nil
}

func (c *client) Update(collection string, filter, update bson.M) (int64, error) {
	ctx := context.TODO()
	coll := c.DB.Database(Database).Collection(collection)

	res, err := coll.UpdateMany(ctx, filter, update)
	if err != nil {
		return 0, err
	}

	return res.ModifiedCount, nil
}

func (c *client) Delete(collection string, filter bson.M) (int64, error) {
	ctx := context.TODO()
	coll := c.DB.Database(Database).Collection(collection)

	res, err := coll.DeleteMany(ctx, filter)
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
