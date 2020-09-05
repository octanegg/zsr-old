package octane

import (
	"encoding/json"
	"errors"
	"io"
	"reflect"
	"time"

	"github.com/octanegg/core/internal/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Match .
type Match struct {
	ID       *primitive.ObjectID `json:"id" bson:"_id"`
	OctaneID string              `json:"octane_id" bson:"octane_id"`
	EventID  *primitive.ObjectID `json:"event" bson:"event"`
	Stage    int                 `json:"stage" bson:"stage"`
	Substage int                 `json:"substage,omitempty" bson:"substage,omitempty"`
	Date     *time.Time          `json:"date,omitempty" bson:"date,omitempty"`
	Format   string              `json:"format" bson:"format"`
	Blue     *MatchSide          `json:"blue,omitempty" bson:"blue,omitempty"`
	Orange   *MatchSide          `json:"orange,omitempty" bson:"orange,omitempty"`
	Mode     int                 `json:"mode" bson:"mode"`
	Number   int                 `json:"number" bson:"number"`
}

// MatchSide .
type MatchSide struct {
	Score   int                   `json:"score" bson:"score"`
	Winner  bool                  `json:"winner" bson:"winner"`
	Team    *primitive.ObjectID   `json:"team" bson:"team"`
	Players []*primitive.ObjectID `json:"players,omitempty" bson:"players,omitempty"`
}

func (c *client) FindMatches(filter bson.M, pagination *Pagination, sort *Sort) (*Data, error) {
	matches, err := c.Find(config.CollectionMatches, filter, pagination, sort, func(cursor *mongo.Cursor) (interface{}, error) {
		var match Match
		if err := cursor.Decode(&match); err != nil {
			return nil, err
		}
		return match, nil
	})

	if err != nil {
		return nil, err
	}

	if matches == nil {
		matches = make([]interface{}, 0)
	}

	if pagination != nil {
		pagination.PageSize = len(matches)
	}

	return &Data{
		matches,
		pagination,
	}, nil
}

func (c *client) FindMatch(oid *primitive.ObjectID) (*Match, error) {
	matches, err := c.FindMatches(bson.M{config.KeyID: oid}, nil, nil)
	if err != nil {
		return nil, err
	}

	if len(matches.Data) == 0 {
		return nil, nil
	}

	match := matches.Data[0].(Match)
	return &match, nil
}

func (c *client) InsertMatchWithReader(body io.ReadCloser) (*primitive.ObjectID, error) {
	var match Match
	if err := json.NewDecoder(body).Decode(&match); err != nil {
		return nil, err
	}

	id := primitive.NewObjectID()
	match.ID = &id
	oid, err := c.Insert(config.CollectionMatches, match)
	if err != nil {
		return nil, err
	}

	return oid, nil
}

func (c *client) UpdateMatchWithReader(oid *primitive.ObjectID, body io.ReadCloser) (*primitive.ObjectID, error) {
	match, err := c.FindMatch(oid)
	if err != nil {
		return nil, err
	}

	if match == nil {
		return nil, errors.New(config.ErrNoObjectFoundForID)
	}

	var fields Match
	if err := json.NewDecoder(body).Decode(&fields); err != nil {
		return nil, err
	}

	update := updateFields(reflect.ValueOf(&match).Elem(), reflect.ValueOf(&fields).Elem()).(Match)
	update.ID = oid

	if err := c.Replace(config.CollectionMatches, oid, update); err != nil {
		return nil, err
	}

	return oid, nil
}

func (c *client) InsertMatch(match *Match) (*primitive.ObjectID, error) {
	id := primitive.NewObjectID()
	match.ID = &id
	oid, err := c.Insert(config.CollectionMatches, match)
	if err != nil {
		return nil, err
	}

	return oid, nil
}

func (c *client) UpdateMatch(oid *primitive.ObjectID, fields *Match) (*primitive.ObjectID, error) {
	match, err := c.FindMatch(oid)
	if err != nil {
		return nil, err
	}

	if match == nil {
		return nil, errors.New(config.ErrNoObjectFoundForID)
	}

	update := updateFields(reflect.ValueOf(match).Elem(), reflect.ValueOf(fields).Elem()).(Match)
	update.ID = oid

	if err := c.Replace(config.CollectionMatches, oid, update); err != nil {
		return nil, err
	}

	return oid, nil
}

func (c *client) DeleteMatch(oid *primitive.ObjectID) (int64, error) {
	return c.Delete(config.CollectionMatches, oid)
}
