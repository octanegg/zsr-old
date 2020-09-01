package octane

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Games .
type Games struct {
	Games []interface{} `json:"games"`
}

// Game .
type Game struct {
	ID       primitive.ObjectID `json:"id" bson:"_id"`
	MatchID  primitive.ObjectID `json:"match" bson:"match"`
	EventID  primitive.ObjectID `json:"event" bson:"event"`
	Map      string             `json:"map" bson:"map"`
	Duration int                `json:"duration" bson:"duration"`
	Mode     int                `json:"mode" bson:"mode"`
	Blue     GameTeam           `json:"blue" bson:"blue"`
	Orange   GameTeam           `json:"orange" bson:"orange"`
}

// GameTeam .
type GameTeam struct {
	Goals   int          `json:"goals" bson:"goals"`
	Winner  bool         `json:"winner" bson:"winner"`
	Team    Team         `json:"team" bson:"team"`
	Players []GamePlayer `json:"players" bson:"players"`
}

// GamePlayer .
type GamePlayer struct {
	Player Player `json:"player" bson:"player"`
	Stats  Stats  `json:"stats" bson:"stats"`
}

func (c *client) FindGames(filter bson.M) (*Games, error) {
	games, err := c.Find("games", filter, func(cursor *mongo.Cursor) (interface{}, error) {
		var game Game
		if err := cursor.Decode(&game); err != nil {
			return nil, err
		}
		return game, nil
	})

	if err != nil {
		return nil, err
	}

	return &Games{games}, nil
}

func (c *client) FindGameByID(id string) (*Game, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	games, err := c.FindGames(bson.M{"_id": oid})
	if err != nil {
		return nil, err
	}

	if len(games.Games) == 0 {
		return nil, nil
	}

	game := games.Games[0].(Game)
	return &game, nil
}
