package octane

import (
	"time"

	"github.com/octanegg/zsr/ballchasing"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

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

func (c *client) FindGames(filter bson.M, pagination *Pagination, sort *Sort) (*Data, error) {
	games, err := c.Find(CollectionGames, filter, pagination, sort, func(cursor *mongo.Cursor) (interface{}, error) {
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

func (c *client) InsertGame(game *Game) (*primitive.ObjectID, error) {
	oid, err := c.InsertOne(CollectionGames, game)
	if err != nil {
		return nil, err
	}

	return oid, nil
}

func (c *client) InsertGames(games []interface{}) ([]interface{}, error) {
	ids, err := c.InsertMany(CollectionGames, games)
	if err != nil {
		return nil, err
	}

	return ids, nil
}

func (c *client) ReplaceGame(oid *primitive.ObjectID, game *Game) (*primitive.ObjectID, error) {
	if err := c.Replace(CollectionGames, oid, game); err != nil {
		return nil, err
	}

	return oid, nil
}

func (c *client) UpdateGames(filter, update bson.M) (int64, error) {
	return c.Update(CollectionGames, filter, update)
}

func (c *client) DeleteGame(filter bson.M) (int64, error) {
	return c.Delete(CollectionGames, filter)
}
