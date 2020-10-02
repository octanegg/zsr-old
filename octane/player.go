package octane

import (
	"github.com/octanegg/core/internal/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Player .
type Player struct {
	ID      *primitive.ObjectID `json:"_id" bson:"_id"`
	Tag     string              `json:"tag" bson:"tag"`
	Name    string              `json:"name,omitempty" bson:"name,omitempty"`
	Country string              `json:"country,omitempty" bson:"country,omitempty"`
	Team    string              `json:"team,omitempty" bson:"team,omitempty"`
	Account *Account            `json:"account" bson:"account"`
}

// Account .
type Account struct {
	Platform string `json:"platform" bson:"platform"`
	ID       string `json:"id" bson:"id"`
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

func (c *client) FindPlayer(oid *primitive.ObjectID) (*Player, error) {
	players, err := c.FindPlayers(bson.M{config.KeyID: oid}, nil, nil)
	if err != nil {
		return nil, err
	}

	if len(players.Data) == 0 {
		return nil, nil
	}

	player := players.Data[0].(Player)
	return &player, nil
}

func (c *client) InsertPlayer(player *Player) (*primitive.ObjectID, error) {
	id := primitive.NewObjectID()
	player.ID = &id
	oid, err := c.Insert(config.CollectionPlayers, player)
	if err != nil {
		return nil, err
	}

	return oid, nil
}

func (c *client) ReplacePlayer(oid *primitive.ObjectID, player *Player) (*primitive.ObjectID, error) {
	if err := c.Replace(config.CollectionPlayers, oid, player); err != nil {
		return nil, err
	}

	return oid, nil
}

func (c *client) UpdatePlayers(filter, update bson.M) (int64, error) {
	return c.Update(config.CollectionPlayers, filter, update)
}

func (c *client) DeletePlayer(oid *primitive.ObjectID) (int64, error) {
	return c.Delete(config.CollectionPlayers, oid)
}
