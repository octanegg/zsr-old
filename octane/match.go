package octane

import (
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Matches .
type Matches struct {
	Matches []*Match `json:"matches"`
	*Pagination
}

// Match .
type Match struct {
	ID       *primitive.ObjectID `json:"_id" bson:"_id"`
	OctaneID string              `json:"octane_id" bson:"octane_id"`
	Event    *Event              `json:"event" bson:"event"`
	Stage    *Stage              `json:"stage" bson:"stage"`
	Substage int                 `json:"substage,omitempty" bson:"substage,omitempty"`
	Date     *time.Time          `json:"date,omitempty" bson:"date,omitempty"`
	Format   string              `json:"format" bson:"format"`
	Blue     *MatchSide          `json:"blue,omitempty" bson:"blue,omitempty"`
	Orange   *MatchSide          `json:"orange,omitempty" bson:"orange,omitempty"`
	Number   int                 `json:"number" bson:"number"`
}

// MatchSide .
type MatchSide struct {
	Score  int   `json:"score" bson:"score"`
	Winner bool  `json:"winner" bson:"winner"`
	Team   *Team `json:"team" bson:"team"`
}

func (c *client) FindMatches(ctx *FindContext) (*Matches, error) {
	coll := c.DB.Database(Database).Collection(CollectionMatches)

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

	var matches []*Match
	for cursor.Next(context.TODO()) {
		var match Match
		if err := cursor.Decode(&match); err != nil {
			return nil, err
		}
		matches = append(matches, &match)
	}

	if err != nil {
		return nil, err
	}

	if ctx.Pagination != nil {
		ctx.Pagination.PageSize = len(matches)
	}

	return &Matches{
		matches,
		ctx.Pagination,
	}, nil
}

func (c *client) FindMatch(filter bson.M) (*Match, error) {
	matches, err := c.FindMatches(&FindContext{Filter: filter})
	if err != nil {
		return nil, err
	}

	if len(matches.Matches) == 0 {
		return nil, errors.New("no match found")
	}

	return matches.Matches[0], nil
}

func (c *client) InsertMatches(matches []interface{}) ([]interface{}, error) {
	ctx := context.TODO()
	coll := c.DB.Database(Database).Collection(CollectionMatches)

	res, err := coll.InsertMany(ctx, matches)
	if err != nil {
		return nil, err
	}

	return res.InsertedIDs, nil
}

func (c *client) InsertMatch(match interface{}) (*primitive.ObjectID, error) {
	ids, err := c.InsertMatches([]interface{}{match})
	if err != nil {
		return nil, err
	}

	id := ids[0].(primitive.ObjectID)
	return &id, nil
}

func (c *client) DeleteMatch(filter bson.M) (int64, error) {
	ctx := context.TODO()
	coll := c.DB.Database(Database).Collection(CollectionMatches)

	res, err := coll.DeleteMany(ctx, filter)
	if err != nil {
		return 0, err
	}

	return res.DeletedCount, nil
}
