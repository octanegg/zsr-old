package octane

import (
	"errors"
	"reflect"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Matches .
type Matches struct {
	Matches []interface{} `json:"matches"`
}

// Match .
type Match struct {
	ID       *primitive.ObjectID   `json:"id" bson:"_id"`
	EventID  *primitive.ObjectID   `json:"event" bson:"event"`
	Stage    *int                  `json:"stage" bson:"stage"`
	Substage *int                  `json:"substage,omitempty" bson:"substage,omitempty"`
	Date     *time.Time            `json:"date,omitempty" bson:"date,omitempty"`
	Format   *string               `json:"format" bson:"format"`
	Blue     *MatchSide            `json:"blue,omitempty" bson:"blue,omitempty"`
	Orange   *MatchSide            `json:"orange,omitempty" bson:"orange,omitempty"`
	Games    []*primitive.ObjectID `json:"games,omitempty" bson:"games,omitempty"`
	Mode     *int                  `json:"mode" bson:"mode"`
}

// MatchSide .
type MatchSide struct {
	Score   *int      `json:"score,omitempty" bson:"score,omitempty"`
	Winner  *bool     `json:"winner,omitempty" bson:"winner,omitempty"`
	Team    *Team     `json:"team,omitempty" bson:"team,omitempty"`
	Players []*Player `json:"players,omitempty" bson:"players,omitempty"`
}

func (c *client) FindMatches(filter bson.M) (*Matches, error) {
	matchs, err := c.Find("matches", filter, func(cursor *mongo.Cursor) (interface{}, error) {
		var match Match
		if err := cursor.Decode(&match); err != nil {
			return nil, err
		}
		return match, nil
	})

	if err != nil {
		return nil, err
	}

	return &Matches{matchs}, nil
}

func (c *client) FindMatch(oid *primitive.ObjectID) (*Match, error) {
	matches, err := c.FindMatches(bson.M{"_id": oid})
	if err != nil {
		return nil, err
	}

	if len(matches.Matches) == 0 {
		return nil, nil
	}

	match := matches.Matches[0].(Match)
	return &match, nil
}

func (c *client) InsertMatch(match *Match) (*ObjectID, error) {
	event, err := c.FindEvent(match.EventID)
	if err != nil {
		return nil, err
	}

	if event == nil {
		return nil, errors.New("No event found for ID")
	}

	id := primitive.NewObjectID()
	match.ID = &id
	match.Mode = event.Mode

	oid, err := c.Insert("matches", match)
	if err != nil {
		return nil, err
	}

	// TODO: Update event with list of matches

	return &ObjectID{oid.(primitive.ObjectID).Hex()}, nil
}

func (c *client) UpdateMatch(oid *primitive.ObjectID, fields *Match) (*ObjectID, error) {
	match, err := c.FindMatch(oid)
	if err != nil {
		return nil, err
	}

	if match == nil {
		return nil, errors.New("No match found for ID")
	}

	update := updateFields(reflect.ValueOf(match).Elem(), reflect.ValueOf(fields).Elem()).(Match)
	update.ID = oid

	id, err := c.Replace("matches", oid, update)
	if err != nil {
		return nil, err
	}

	if id != nil {
		return &ObjectID{id.(primitive.ObjectID).Hex()}, nil
	}

	// TODO: Update event with list of matches

	return &ObjectID{oid.Hex()}, nil
}

func (c *client) DeleteMatch(oid *primitive.ObjectID) (int64, error) {
	return c.Delete("events", oid)
}
