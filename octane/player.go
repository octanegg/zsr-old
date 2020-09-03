package octane

import (
	"errors"
	"reflect"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Player .
type Player struct {
	ID      *primitive.ObjectID `json:"id" bson:"_id"`
	Tag     *string             `json:"tag" bson:"tag"`
	Name    *string             `json:"name" bson:"name"`
	Country *string             `json:"country" bson:"country"`
	Team    *string             `json:"team" bson:"team"`
}

func (c *client) FindPlayers(filter bson.M, pagination *Pagination) (*Data, error) {
	players, err := c.Find("players", filter, pagination, func(cursor *mongo.Cursor) (interface{}, error) {
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

	if pagination != nil {
		pagination.PageSize = len(players)
	}

	return &Data{
		players,
		pagination,
	}, nil
}

func (c *client) FindPlayer(oid *primitive.ObjectID) (*Player, error) {
	players, err := c.FindPlayers(bson.M{"_id": oid}, nil)
	if err != nil {
		return nil, err
	}

	if len(players.Data) == 0 {
		return nil, nil
	}

	player := players.Data[0].(Player)
	return &player, nil
}

func (c *client) InsertPlayer(player *Player) (*ObjectID, error) {
	id := primitive.NewObjectID()
	player.ID = &id

	oid, err := c.Insert("players", player)
	if err != nil {
		return nil, err
	}

	return &ObjectID{oid.(primitive.ObjectID).Hex()}, nil
}

func (c *client) UpdatePlayer(oid *primitive.ObjectID, fields *Player) (*ObjectID, error) {
	player, err := c.FindPlayer(oid)
	if err != nil {
		return nil, err
	}

	if player == nil {
		return nil, errors.New("No player found for ID")
	}

	update := updateFields(reflect.ValueOf(player).Elem(), reflect.ValueOf(fields).Elem()).(Player)
	update.ID = oid

	id, err := c.Replace("players", oid, update)
	if err != nil {
		return nil, err
	}

	if id != nil {
		return &ObjectID{id.(primitive.ObjectID).Hex()}, nil
	}

	return &ObjectID{oid.Hex()}, nil
}

func (c *client) DeletePlayer(oid *primitive.ObjectID) (int64, error) {
	return c.Delete("players", oid)
}
