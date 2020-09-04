package octane

import (
	"encoding/json"
	"errors"
	"io"
	"reflect"
	"time"

	"github.com/octanegg/core/internal/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Game .
type Game struct {
	ID       *primitive.ObjectID `json:"id" bson:"_id"`
	Number   *int                `json:"number" bson:"number"`
	MatchID  *primitive.ObjectID `json:"match" bson:"match"`
	EventID  *primitive.ObjectID `json:"event" bson:"event"`
	Map      *string             `json:"map" bson:"map"`
	Duration *int                `json:"duration" bson:"duration"`
	Mode     *int                `json:"mode" bson:"mode"`
	Date     *time.Time          `json:"date,omitempty" bson:"date,omitempty"`
	Blue     *GameSide           `json:"blue" bson:"blue"`
	Orange   *GameSide           `json:"orange" bson:"orange"`
}

// GameSide .
type GameSide struct {
	Goals   *int           `json:"goals" bson:"goals"`
	Winner  bool           `json:"winner" bson:"winner"`
	Team    *TeamStats     `json:"team" bson:"team"`
	Players []*PlayerStats `json:"players" bson:"players"`
}

// PlayerStats .
type PlayerStats struct {
	Player *primitive.ObjectID `json:"player" bson:"player"`
	Stats  *Stats              `json:"stats" bson:"stats"`
}

// TeamStats .
type TeamStats struct {
	ID *primitive.ObjectID `json:"id" bson:"id"`
}

func (c *client) FindGames(filter bson.M, pagination *Pagination, sort *Sort) (*Data, error) {
	games, err := c.Find(config.CollectionGames, filter, pagination, sort, func(cursor *mongo.Cursor) (interface{}, error) {
		var game Game
		if err := cursor.Decode(&game); err != nil {
			return nil, err
		}
		return game, nil
	})

	if err != nil {
		return nil, err
	}

	if games == nil {
		games = make([]interface{}, 0)
	}

	if pagination != nil {
		pagination.PageSize = len(games)
	}

	return &Data{
		games,
		pagination,
	}, nil
}

func (c *client) FindGame(oid *primitive.ObjectID) (interface{}, error) {
	games, err := c.FindGames(bson.M{config.KeyID: oid}, nil, nil)
	if err != nil {
		return nil, err
	}

	if len(games.Data) == 0 {
		return nil, nil
	}

	return games.Data[0].(Game), nil
}

func (c *client) InsertGame(body io.ReadCloser) (*ObjectID, error) {
	var game Game
	if err := json.NewDecoder(body).Decode(&game); err != nil {
		return nil, err
	}

	id := primitive.NewObjectID()
	game.ID = &id
	oid, err := c.Insert(config.CollectionGames, game)
	if err != nil {
		return nil, err
	}

	return &ObjectID{oid.(primitive.ObjectID).Hex()}, nil
}

func (c *client) UpdateGame(oid *primitive.ObjectID, body io.ReadCloser) (*ObjectID, error) {
	data, err := c.FindGame(oid)
	if err != nil {
		return nil, err
	}

	if data == nil {
		return nil, errors.New(config.ErrNoObjectFoundForID)
	}

	var fields Game
	if err := json.NewDecoder(body).Decode(&fields); err != nil {
		return nil, err
	}

	game := data.(Game)
	update := updateFields(reflect.ValueOf(&game).Elem(), reflect.ValueOf(&fields).Elem()).(Game)
	update.ID = oid

	id, err := c.Replace(config.CollectionGames, oid, update)
	if err != nil {
		return nil, err
	}

	if id != nil {
		return &ObjectID{id.(primitive.ObjectID).Hex()}, nil
	}
	return &ObjectID{oid.Hex()}, err
}

func (c *client) DeleteGame(oid *primitive.ObjectID) (int64, error) {
	return c.Delete(config.CollectionGames, oid)
}
