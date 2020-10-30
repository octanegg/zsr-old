package octane

import (
	"context"
	"errors"
	"time"

	"github.com/octanegg/zsr/ballchasing"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Games .
type Games struct {
	Games []*Game `json:"games"`
	*Pagination
}

// Game .
type Game struct {
	ID            *primitive.ObjectID `json:"_id" bson:"_id"`
	OctaneID      string              `json:"octane_id" bson:"octane_id"`
	Number        int                 `json:"number" bson:"number"`
	Match         *Match              `json:"match" bson:"match"`
	Map           string              `json:"map" bson:"map"`
	Duration      int                 `json:"duration" bson:"duration"`
	Date          *time.Time          `json:"date,omitempty" bson:"date,omitempty"`
	Blue          *GameSide           `json:"blue" bson:"blue"`
	Orange        *GameSide           `json:"orange" bson:"orange"`
	BallchasingID string              `json:"ballchasing,omitempty" bson:"ballchasing,omitempty"`
}

// GameSide .
type GameSide struct {
	Goals   int            `json:"goals" bson:"goals"`
	Winner  bool           `json:"winner" bson:"winner"`
	Team    *Team          `json:"team" bson:"team"`
	Players []*PlayerStats `json:"players" bson:"players"`
}

// PlayerStats .
type PlayerStats struct {
	Player *Player                  `json:"player" bson:"player"`
	Stats  *ballchasing.PlayerStats `json:"stats" bson:"stats"`
	Rating float64                  `json:"rating" bson:"rating"`
}

func (c *client) FindGames(ctx *FindContext) (*Games, error) {
	coll := c.DB.Database(Database).Collection(CollectionGames)

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

	var games []*Game
	for cursor.Next(context.TODO()) {
		var game Game
		if err := cursor.Decode(&game); err != nil {
			return nil, err
		}
		games = append(games, &game)
	}

	if err != nil {
		return nil, err
	}

	if ctx.Pagination != nil {
		ctx.Pagination.PageSize = len(games)
	}

	return &Games{
		games,
		ctx.Pagination,
	}, nil
}

func (c *client) FindGame(filter bson.M) (*Game, error) {
	games, err := c.FindGames(&FindContext{Filter: filter})
	if err != nil {
		return nil, err
	}

	if len(games.Games) == 0 {
		return nil, errors.New("no game found")
	}

	return games.Games[0], nil
}

func (c *client) InsertGames(games []interface{}) ([]interface{}, error) {
	ctx := context.TODO()
	coll := c.DB.Database(Database).Collection(CollectionGames)

	res, err := coll.InsertMany(ctx, games)
	if err != nil {
		return nil, err
	}

	return res.InsertedIDs, nil
}

func (c *client) InsertGame(game interface{}) (*primitive.ObjectID, error) {
	ids, err := c.InsertGames([]interface{}{game})
	if err != nil {
		return nil, err
	}

	id := ids[0].(primitive.ObjectID)
	return &id, nil
}

func (c *client) DeleteGame(filter bson.M) (int64, error) {
	ctx := context.TODO()
	coll := c.DB.Database(Database).Collection(CollectionGames)

	res, err := coll.DeleteMany(ctx, filter)
	if err != nil {
		return 0, err
	}

	return res.DeletedCount, nil
}
