package octane

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Players .
type Players struct {
	Players []*Player `json:"players"`
	*Pagination
}

// Player .
type Player struct {
	ID      *primitive.ObjectID `json:"_id" bson:"_id"`
	Tag     string              `json:"tag" bson:"tag"`
	Name    string              `json:"name,omitempty" bson:"name,omitempty"`
	Country string              `json:"country,omitempty" bson:"country,omitempty"`
	Team    string              `json:"team,omitempty" bson:"team,omitempty"`
	Account *Account            `json:"account,omitempty" bson:"account,omitempty"`
}

// Account .
type Account struct {
	Platform string `json:"platform" bson:"platform"`
	ID       string `json:"id" bson:"id"`
}

func (c *client) FindPlayers(ctx *FindContext) (*Players, error) {
	coll := c.DB.Database(Database).Collection(CollectionPlayers)

	opts := options.Find()
	if ctx.Pagination != nil {
		opts.SetSkip((ctx.Pagination.Page - 1) * ctx.Pagination.PerPage)
		opts.SetLimit(ctx.Pagination.PerPage)
	}

	opts.SetSort(ctx.Sort)
	cursor, err := coll.Find(context.TODO(), ctx.Filter, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	var players []*Player
	for cursor.Next(context.TODO()) {
		var player Player
		if err := cursor.Decode(&player); err != nil {
			return nil, err
		}
		players = append(players, &player)
	}

	if err != nil {
		return nil, err
	}

	if ctx.Pagination != nil {
		ctx.Pagination.PageSize = len(players)
	}

	return &Players{
		players,
		ctx.Pagination,
	}, nil
}

func (c *client) FindPlayer(filter bson.M) (*Player, error) {
	players, err := c.FindPlayers(&FindContext{Filter: filter})
	if err != nil {
		return nil, err
	}

	if len(players.Players) == 0 {
		return nil, errors.New("no player found")
	}

	return players.Players[0], nil
}

func (c *client) InsertPlayers(players []interface{}) ([]interface{}, error) {
	ctx := context.TODO()
	coll := c.DB.Database(Database).Collection(CollectionPlayers)

	res, err := coll.InsertMany(ctx, players)
	if err != nil {
		return nil, err
	}

	return res.InsertedIDs, nil
}

func (c *client) InsertPlayer(player interface{}) (*primitive.ObjectID, error) {
	ids, err := c.InsertPlayers([]interface{}{player})
	if err != nil {
		return nil, err
	}

	id := ids[0].(primitive.ObjectID)
	return &id, nil
}

func (c *client) DeletePlayer(filter bson.M) (int64, error) {
	ctx := context.TODO()
	coll := c.DB.Database(Database).Collection(CollectionPlayers)

	res, err := coll.DeleteMany(ctx, filter)
	if err != nil {
		return 0, err
	}

	return res.DeletedCount, nil
}
