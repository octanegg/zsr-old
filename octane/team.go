package octane

import (
	"errors"
	"reflect"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Team .
type Team struct {
	ID   *primitive.ObjectID `json:"id" bson:"_id"`
	Name *string             `json:"name" bson:"name"`
}

func (c *client) FindTeams(filter bson.M, page, perPage int64) (*Data, error) {
	teams, err := c.Find("teams", filter, page, perPage, func(cursor *mongo.Cursor) (interface{}, error) {
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

	return &Data{
		Page:     page,
		PerPage:  perPage,
		PageSize: len(teams),
		Data:     teams,
	}, nil
}

func (c *client) FindTeam(oid *primitive.ObjectID) (*Team, error) {
	teams, err := c.FindTeams(bson.M{"_id": oid}, 0, 0)
	if err != nil {
		return nil, err
	}

	if len(teams.Data) == 0 {
		return nil, nil
	}

	team := teams.Data[0].(Team)
	return &team, nil
}

func (c *client) InsertTeam(team *Team) (*ObjectID, error) {
	id := primitive.NewObjectID()
	team.ID = &id

	oid, err := c.Insert("teams", team)
	if err != nil {
		return nil, err
	}

	return &ObjectID{oid.(primitive.ObjectID).Hex()}, nil
}

func (c *client) UpdateTeam(oid *primitive.ObjectID, fields *Team) (*ObjectID, error) {
	team, err := c.FindTeam(oid)
	if err != nil {
		return nil, err
	}

	if team == nil {
		return nil, errors.New("No team found for ID")
	}

	update := updateFields(reflect.ValueOf(team).Elem(), reflect.ValueOf(fields).Elem()).(Team)
	update.ID = oid

	id, err := c.Replace("teams", oid, update)
	if err != nil {
		return nil, err
	}

	if id != nil {
		return &ObjectID{id.(primitive.ObjectID).Hex()}, nil
	}

	return &ObjectID{oid.Hex()}, nil
}

func (c *client) DeleteTeam(oid *primitive.ObjectID) (int64, error) {
	return c.Delete("teams", oid)
}
