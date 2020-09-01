package octane

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Teams .
type Teams struct {
	Teams []interface{} `json:"teams"`
}

// Team .
type Team struct {
	ID   primitive.ObjectID `json:"id" bson:"_id"`
	Name string             `json:"name" bson:"name"`
}

func (c *client) FindTeams(filter bson.M) (*Teams, error) {
	teams, err := c.Find("teams", filter, func(cursor *mongo.Cursor) (interface{}, error) {
		var team Team
		if err := cursor.Decode(&team); err != nil {
			return nil, err
		}
		return team, nil
	})

	if err != nil {
		return nil, err
	}

	return &Teams{teams}, nil
}

func (c *client) FindTeamByID(id string) (*Team, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	teams, err := c.FindTeams(bson.M{"_id": oid})
	if err != nil {
		return nil, err
	}

	if len(teams.Teams) == 0 {
		return nil, nil
	}

	team := teams.Teams[0].(Team)
	return &team, nil
}
