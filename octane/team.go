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

// Team .
type Team struct {
	ID   *primitive.ObjectID `json:"id" bson:"_id"`
	Name string              `json:"name" bson:"name"`
}

func (c *client) FindTeams(filter bson.M, pagination *Pagination, sort *Sort) (*Data, error) {
	teams, err := c.Find(config.CollectionTeams, filter, pagination, sort, func(cursor *mongo.Cursor) (interface{}, error) {
		var team Team
		if err := cursor.Decode(&team); err != nil {
			return nil, err
		}
		return team, nil
	})

	if err != nil {
		return nil, err
	}

	if teams == nil {
		teams = make([]interface{}, 0)
	}

	if pagination != nil {
		pagination.PageSize = len(teams)
	}

	return &Data{
		teams,
		pagination,
	}, nil
}

func (c *client) FindTeam(oid *primitive.ObjectID) (*Team, error) {
	teams, err := c.FindTeams(bson.M{config.KeyID: oid}, nil, nil)
	if err != nil {
		return nil, err
	}

	if len(teams.Data) == 0 {
		return nil, nil
	}

	team := teams.Data[0].(Team)
	return &team, nil
}

func (c *client) InsertTeamWithReader(body io.ReadCloser) (*primitive.ObjectID, error) {
	var team Team
	if err := json.NewDecoder(body).Decode(&team); err != nil {
		return nil, err
	}

	id := primitive.NewObjectID()
	team.ID = &id
	oid, err := c.Insert(config.CollectionTeams, team)
	if err != nil {
		return nil, err
	}

	return oid, nil
}

func (c *client) UpdateTeamWithReader(oid *primitive.ObjectID, body io.ReadCloser) (*primitive.ObjectID, error) {
	team, err := c.FindTeam(oid)
	if err != nil {
		return nil, err
	}

	if team == nil {
		return nil, errors.New(config.ErrNoObjectFoundForID)
	}

	var fields Team
	if err := json.NewDecoder(body).Decode(&fields); err != nil {
		return nil, err
	}

	update := updateFields(reflect.ValueOf(&team).Elem(), reflect.ValueOf(&fields).Elem()).(Team)
	update.ID = oid

	if err := c.Replace(config.CollectionTeams, oid, update); err != nil {
		return nil, err
	}

	return oid, nil
}

func (c *client) InsertTeam(team *Team) (*primitive.ObjectID, error) {
	id := primitive.NewObjectID()
	team.ID = &id
	oid, err := c.Insert(config.CollectionTeams, team)
	if err != nil {
		return nil, err
	}

	return oid, nil
}

func (c *client) UpdateTeam(oid *primitive.ObjectID, fields *Team) (*primitive.ObjectID, error) {
	team, err := c.FindTeam(oid)
	if err != nil {
		return nil, err
	}

	if team == nil {
		return nil, errors.New(config.ErrNoObjectFoundForID)
	}

	update := updateFields(reflect.ValueOf(team).Elem(), reflect.ValueOf(fields).Elem()).(Team)
	update.ID = oid

	if err := c.Replace(config.CollectionTeams, oid, update); err != nil {
		return nil, err
	}

	return oid, nil
}

func (c *client) DeleteTeam(oid *primitive.ObjectID) (int64, error) {
	return c.Delete(config.CollectionTeams, oid)
}
