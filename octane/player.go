package octane

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Player .
type Player struct {
	ID  *primitive.ObjectID `json:"id" bson:"_id"`
	Tag *string             `json:"tag" bson:"tag"`
}

func (c *client) FindPlayers(filter bson.M, page, perPage int64) (*Data, error) {
	players, err := c.Find("players", filter, page, perPage, func(cursor *mongo.Cursor) (interface{}, error) {
		var player Player
		if err := cursor.Decode(&player); err != nil {
			return nil, err
		}
		return player, nil
	})

	if err != nil {
		return nil, err
	}

	if players == nil {
		players = make([]interface{}, 0)
	}

	return &Data{
		Page:     page,
		PerPage:  perPage,
		PageSize: len(players),
		Data:     players,
	}, nil
}

func (c *client) FindPlayer(oid *primitive.ObjectID) (*Player, error) {
	players, err := c.FindPlayers(bson.M{"_id": oid}, 0, 0)
	if err != nil {
		return nil, err
	}

	if len(players.Data) == 0 {
		return nil, nil
	}

	player := players.Data[0].(Player)
	return &player, nil
}

// TODO: Update/Insert Players
