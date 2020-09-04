package octane

import (
	"encoding/json"
	"errors"
	"io"
	"reflect"

	"github.com/octanegg/core/internal/config"
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

func (c *client) FindPlayers(filter bson.M, pagination *Pagination, sort *Sort) (*Data, error) {
	players, err := c.Find(config.CollectionPlayers, filter, pagination, sort, func(cursor *mongo.Cursor) (interface{}, error) {
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

func (c *client) FindPlayer(oid *primitive.ObjectID) (interface{}, error) {
	players, err := c.FindPlayers(bson.M{config.KeyID: oid}, nil, nil)
	if err != nil {
		return nil, err
	}

	if len(players.Data) == 0 {
		return nil, nil
	}

	return players.Data[0].(Player), nil
}

func (c *client) InsertPlayer(body io.ReadCloser) (*ObjectID, error) {
	var player Player
	if err := json.NewDecoder(body).Decode(&player); err != nil {
		return nil, err
	}

	id := primitive.NewObjectID()
	player.ID = &id
	oid, err := c.Insert(config.CollectionPlayers, player)
	if err != nil {
		return nil, err
	}

	return &ObjectID{oid.(primitive.ObjectID).Hex()}, nil
}

func (c *client) UpdatePlayer(oid *primitive.ObjectID, body io.ReadCloser) (*ObjectID, error) {
	data, err := c.FindPlayer(oid)
	if err != nil {
		return nil, err
	}

	if data == nil {
		return nil, errors.New(config.ErrNoObjectFoundForID)
	}

	var fields Player
	if err := json.NewDecoder(body).Decode(&fields); err != nil {
		return nil, err
	}

	player := data.(Player)
	update := updateFields(reflect.ValueOf(&player).Elem(), reflect.ValueOf(&fields).Elem()).(Player)
	update.ID = oid

	id, err := c.Replace(config.CollectionPlayers, oid, update)
	if err != nil {
		return nil, err
	}

	if id != nil {
		return &ObjectID{id.(primitive.ObjectID).Hex()}, nil
	}

	return &ObjectID{oid.Hex()}, nil
}

func (c *client) DeletePlayer(oid *primitive.ObjectID) (int64, error) {
	return c.Delete(config.CollectionPlayers, oid)
}
