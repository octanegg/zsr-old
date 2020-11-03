package octane

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

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
