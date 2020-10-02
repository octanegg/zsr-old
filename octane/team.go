package octane

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Team .
type Team struct {
	ID   *primitive.ObjectID `json:"_id" bson:"_id"`
	Name string              `json:"name" bson:"name"`
}

func (c *client) FindTeams(filter bson.M, pagination *Pagination, sort *Sort) (*Data, error) {
	teams, err := c.Find(CollectionTeams, filter, pagination, sort, func(cursor *mongo.Cursor) (interface{}, error) {
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
	teams, err := c.FindTeams(bson.M{"_id": oid}, nil, nil)
	if err != nil {
		return nil, err
	}

	if len(teams.Data) == 0 {
		return nil, nil
	}

	team := teams.Data[0].(Team)
	return &team, nil
}

func (c *client) InsertTeam(team *Team) (*primitive.ObjectID, error) {
	id := primitive.NewObjectID()
	team.ID = &id
	oid, err := c.Insert(CollectionTeams, team)
	if err != nil {
		return nil, err
	}

	return oid, nil
}

func (c *client) ReplaceTeam(oid *primitive.ObjectID, team *Team) (*primitive.ObjectID, error) {
	if err := c.Replace(CollectionTeams, oid, team); err != nil {
		return nil, err
	}

	return oid, nil
}

func (c *client) UpdateTeams(filter, update bson.M) (int64, error) {
	return c.Update(CollectionTeams, filter, update)
}

func (c *client) DeleteTeam(oid *primitive.ObjectID) (int64, error) {
	return c.Delete(CollectionTeams, oid)
}
