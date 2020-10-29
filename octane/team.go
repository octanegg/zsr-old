package octane

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Teams .
type Teams struct {
	Teams []*Team `json:"teams"`
	*Pagination
}

// Team .
type Team struct {
	ID   *primitive.ObjectID `json:"_id" bson:"_id"`
	Name string              `json:"name" bson:"name"`
}

func (c *client) FindTeams(filter bson.M, pagination *Pagination, sort *Sort) (*Teams, error) {
	ctx := context.TODO()
	coll := c.DB.Database(Database).Collection(CollectionTeams)

	opts := options.Find()
	if pagination != nil {
		opts.SetSkip((pagination.Page - 1) * pagination.PerPage)
		opts.SetLimit(pagination.PerPage)
	}

	if sort != nil {
		opts.SetSort(bson.M{sort.Field: sort.Order})
	}

	cursor, err := coll.Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var teams []*Team
	for cursor.Next(ctx) {
		var team Team
		if err := cursor.Decode(&team); err != nil {
			return nil, err
		}
		teams = append(teams, &team)
	}

	if err != nil {
		return nil, err
	}

	if pagination != nil {
		pagination.PageSize = len(teams)
	}

	return &Teams{
		teams,
		pagination,
	}, nil
}

func (c *client) FindTeam(oid *primitive.ObjectID) (*Team, error) {
	teams, err := c.FindTeams(bson.M{"_id": oid}, nil, nil)
	if err != nil {
		return nil, err
	}

	if len(teams.Teams) == 0 {
		return nil, nil
	}

	return teams.Teams[0], nil
}

func (c *client) InsertTeam(team interface{}) (*primitive.ObjectID, error) {
	ids, err := c.InsertTeams([]interface{}{team})
	if err != nil {
		return nil, err
	}

	id := ids[0].(primitive.ObjectID)
	return &id, nil
}

func (c *client) InsertTeams(teams []interface{}) ([]interface{}, error) {
	ctx := context.TODO()
	coll := c.DB.Database(Database).Collection(CollectionTeams)

	res, err := coll.InsertMany(ctx, teams)
	if err != nil {
		return nil, err
	}

	return res.InsertedIDs, nil
}

func (c *client) DeleteTeam(filter bson.M) (int64, error) {
	ctx := context.TODO()
	coll := c.DB.Database(Database).Collection(CollectionTeams)

	res, err := coll.DeleteMany(ctx, filter)
	if err != nil {
		return 0, err
	}

	return res.DeletedCount, nil
}
