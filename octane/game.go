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
	ID            *primitive.ObjectID `json:"id" bson:"_id"`
	OctaneID      string              `json:"octane_id" bson:"octane_id"`
	Number        int                 `json:"number" bson:"number"`
	MatchID       *primitive.ObjectID `json:"match" bson:"match"`
	EventID       *primitive.ObjectID `json:"event" bson:"event"`
	Map           string              `json:"map" bson:"map"`
	Duration      int                 `json:"duration" bson:"duration"`
	Mode          int                 `json:"mode" bson:"mode"`
	Date          *time.Time          `json:"date,omitempty" bson:"date,omitempty"`
	Blue          *GameSide           `json:"blue" bson:"blue"`
	Orange        *GameSide           `json:"orange" bson:"orange"`
	BallchasingID string              `json:"ballchasing,omitempty" bson:"ballchasing,omitempty"`
}

// GameSide .
type GameSide struct {
	Goals   int            `json:"goals" bson:"goals"`
	Winner  bool           `json:"winner" bson:"winner"`
	Team    *TeamStats     `json:"team" bson:"team"`
	Players []*PlayerStats `json:"players" bson:"players"`
}

// PlayerStats .
type PlayerStats struct {
	Player *primitive.ObjectID `json:"player" bson:"player"`
	Stats  interface{}         `json:"stats" bson:"stats"`
	Rating float64             `json:"rating" bson:"rating"`
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

func (c *client) FindGame(oid *primitive.ObjectID) (*Game, error) {
	games, err := c.FindGames(bson.M{config.KeyID: oid}, nil, nil)
	if err != nil {
		return nil, err
	}

	if len(games.Data) == 0 {
		return nil, nil
	}

	game := games.Data[0].(Game)
	return &game, nil
}

func (c *client) InsertGameWithReader(body io.ReadCloser) (*primitive.ObjectID, error) {
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

	return oid, nil
}

func (c *client) UpdateGameWithReader(oid *primitive.ObjectID, body io.ReadCloser) (*primitive.ObjectID, error) {
	var game Game
	if err := json.NewDecoder(body).Decode(&game); err != nil {
		return nil, err
	}

	if err := c.Replace(config.CollectionGames, oid, game); err != nil {
		return nil, err
	}

	return oid, nil
}

func (c *client) InsertGame(game *Game) (*primitive.ObjectID, error) {
	id := primitive.NewObjectID()
	game.ID = &id
	oid, err := c.Insert(config.CollectionGames, game)
	if err != nil {
		return nil, err
	}

	return oid, nil
}

func (c *client) UpdateGame(oid *primitive.ObjectID, fields *Game) (*primitive.ObjectID, error) {
	game, err := c.FindGame(oid)
	if err != nil {
		return nil, err
	}

	if game == nil {
		return nil, errors.New(config.ErrNoObjectFoundForID)
	}

	update := updateFields(reflect.ValueOf(game).Elem(), reflect.ValueOf(fields).Elem()).(Game)
	update.ID = oid

	if err := c.Replace(config.CollectionGames, oid, update); err != nil {
		return nil, err
	}

	return oid, err
}

func (c *client) DeleteGame(oid *primitive.ObjectID) (int64, error) {
	return c.Delete(config.CollectionGames, oid)
}
