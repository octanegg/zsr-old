package octane

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Players .
type Players struct {
	Players []interface{} `json:"players"`
}

// Player .
type Player struct {
	ID  *primitive.ObjectID `json:"id" bson:"_id"`
	Tag *string             `json:"tag" bson:"tag"`
}

func (c *client) FindPlayers(filter bson.M) (*Players, error) {
	players, err := c.Find("players", filter, func(cursor *mongo.Cursor) (interface{}, error) {
		var player Player
		if err := cursor.Decode(&player); err != nil {
			return nil, err
		}
		return player, nil
	})

	if err != nil {
		return nil, err
	}

	return &Players{players}, nil
}

func (c *client) FindPlayerByID(oid *primitive.ObjectID) (*Player, error) {
	players, err := c.FindPlayers(bson.M{"_id": oid})
	if err != nil {
		return nil, err
	}

	if len(players.Players) == 0 {
		return nil, nil
	}

	player := players.Players[0].(Player)
	return &player, nil
}
