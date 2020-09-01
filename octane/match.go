package octane

import (
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
	ID       primitive.ObjectID   `json:"id" bson:"_id"`
	EventID  primitive.ObjectID   `json:"event" bson:"event"`
	Stage    int                  `json:"stage" bson:"stage"`
	Substage int                  `json:"substage" bson:"substage"`
	Date     time.Time            `json:"date" bson:"date"`
	Format   string               `json:"format" bson:"format"`
	Blue     MatchSide            `json:"blue" bson:"blue"`
	Orange   MatchSide            `json:"orange" bson:"orange"`
	Games    []primitive.ObjectID `json:"games" bson:"games"`
	Mode     int                  `json:"mode" bson:"mode"`
}

// MatchSide .
type MatchSide struct {
	Score   int      `json:"score" bson:"score"`
	Winner  bool     `json:"winner" bson:"winner"`
	Team    Team     `json:"team" bson:"team"`
	Players []Player `json:"players" bson:"players"`
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

func (c *client) FindMatchByID(id string) (*Match, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

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
