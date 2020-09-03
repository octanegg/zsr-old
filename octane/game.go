package octane

import (
	"errors"
	"reflect"
	"time"

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
	Goals   *int                `json:"goals" bson:"goals"`
	Winner  bool                `json:"winner" bson:"winner"`
	Team    *primitive.ObjectID `json:"team" bson:"team"`
	Players []*PlayerStats      `json:"players" bson:"players"`
}

// PlayerStats .
type PlayerStats struct {
	Player *Player `json:"player" bson:"player"`
	Stats  *Stats  `json:"stats" bson:"stats"`
}

func (c *client) FindGames(filter bson.M, pagination *Pagination, sort *Sort) (*Data, error) {
	games, err := c.Find("games", filter, pagination, sort, func(cursor *mongo.Cursor) (interface{}, error) {
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
	games, err := c.FindGames(bson.M{"_id": oid}, nil, nil)
	if err != nil {
		return nil, err
	}

	if len(games.Data) == 0 {
		return nil, nil
	}

	game := games.Data[0].(Game)
	return &game, nil
}

func (c *client) InsertGame(game *Game) (*ObjectID, error) {
	event, err := c.FindEvent(game.EventID)
	if err != nil {
		return nil, err
	}
	if event == nil {
		return nil, errors.New("No event found for ID")
	}

	match, err := c.FindMatch(game.MatchID)
	if err != nil {
		return nil, err
	}
	if match == nil {
		return nil, errors.New("No match found for ID")
	}

	id := primitive.NewObjectID()
	game.ID = &id
	game.Mode = event.Mode
	game.Date = match.Date

	oid, err := c.Insert("games", game)
	if err != nil {
		return nil, err
	}

	return &ObjectID{oid.(primitive.ObjectID).Hex()}, nil
}

func (c *client) UpdateGame(oid *primitive.ObjectID, fields *Game) (*ObjectID, error) {
	game, err := c.FindGame(oid)
	if err != nil {
		return nil, err
	}

	if game == nil {
		return nil, errors.New("No game found for ID")
	}

	update := updateFields(reflect.ValueOf(game).Elem(), reflect.ValueOf(fields).Elem()).(Game)
	update.ID = oid

	id, err := c.Replace("games", oid, update)
	if err != nil {
		return nil, err
	}

	if id != nil {
		return &ObjectID{id.(primitive.ObjectID).Hex()}, nil
	}
	return &ObjectID{oid.Hex()}, err
}

func (c *client) DeleteGame(oid *primitive.ObjectID) (int64, error) {
	return c.Delete("games", oid)
}
