package octane

import (
	"context"
	"errors"

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

func (c *client) FindTeams(ctx *FindContext) (*Teams, error) {
	coll := c.DB.Database(Database).Collection(CollectionTeams)

	opts := options.Find()
	if ctx.Pagination != nil {
		opts.SetSkip((ctx.Pagination.Page - 1) * ctx.Pagination.PerPage)
		opts.SetLimit(ctx.Pagination.PerPage)
	}

	opts.SetSort(ctx.Sort)
	cursor, err := coll.Find(context.TODO(), ctx.Filter, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	var teams []*Team
	for cursor.Next(context.TODO()) {
		var team Team
		if err := cursor.Decode(&team); err != nil {
			return nil, err
		}
		teams = append(teams, &team)
	}

	if err != nil {
		return nil, err
	}

	if ctx.Pagination != nil {
		ctx.Pagination.PageSize = len(teams)
	}

	return &Teams{
		teams,
		ctx.Pagination,
	}, nil
}

func (c *client) FindTeam(filter bson.M) (*Team, error) {
	teams, err := c.FindTeams(&FindContext{Filter: filter})
	if err != nil {
		return nil, err
	}

	if len(teams.Teams) == 0 {
		return nil, errors.New("no team found")
	}

	return teams.Teams[0], nil
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

func (c *client) InsertTeam(team interface{}) (*primitive.ObjectID, error) {
	ids, err := c.InsertTeams([]interface{}{team})
	if err != nil {
		return nil, err
	}

	id := ids[0].(primitive.ObjectID)
	return &id, nil
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
