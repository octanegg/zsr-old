package octane

import (
	"context"

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

func (c *client) FindPlayers(filter bson.M, pagination *Pagination, sort *Sort) (*Players, error) {
	ctx := context.TODO()
	coll := c.DB.Database(Database).Collection(CollectionPlayers)

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

	var players []*Player
	for cursor.Next(ctx) {
		var player Player
		if err := cursor.Decode(&player); err != nil {
			return nil, err
		}
		players = append(players, &player)
	}

	if err != nil {
		return nil, err
	}

	if pagination != nil {
		pagination.PageSize = len(players)
	}

	return &Players{
		players,
		pagination,
	}, nil
}

func (c *client) FindPlayer(oid *primitive.ObjectID) (*Player, error) {
	players, err := c.FindPlayers(bson.M{"_id": oid}, nil, nil)
	if err != nil {
		return nil, err
	}

	if len(players.Players) == 0 {
		return nil, nil
	}

	return players.Players[0], nil
}

func (c *client) InsertPlayer(player interface{}) (*primitive.ObjectID, error) {
	ids, err := c.InsertPlayers([]interface{}{player})
	if err != nil {
		return nil, err
	}

	id := ids[0].(primitive.ObjectID)
	return &id, nil
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

func (c *client) DeletePlayer(filter bson.M) (int64, error) {
	ctx := context.TODO()
	coll := c.DB.Database(Database).Collection(CollectionPlayers)

	res, err := coll.DeleteMany(ctx, filter)
	if err != nil {
		return 0, err
	}

	return res.DeletedCount, nil
}
