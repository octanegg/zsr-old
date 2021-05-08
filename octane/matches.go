package octane

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Match .
type Match struct {
	ID                  *primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Slug                string              `json:"slug,omitempty" bson:"slug,omitempty"`
	OctaneID            string              `json:"octane_id,omitempty" bson:"octane_id,omitempty"`
	Event               *Event              `json:"event,omitempty" bson:"event,omitempty"`
	Stage               *Stage              `json:"stage,omitempty" bson:"stage,omitempty"`
	Substage            int                 `json:"substage,omitempty" bson:"substage,omitempty"`
	Date                *time.Time          `json:"date,omitempty" bson:"date,omitempty"`
	Format              *Format             `json:"format,omitempty" bson:"format,omitempty"`
	Blue                *MatchSide          `json:"blue,omitempty" bson:"blue,omitempty"`
	Orange              *MatchSide          `json:"orange,omitempty" bson:"orange,omitempty"`
	Number              int                 `json:"number,omitempty" bson:"number,omitempty"`
	ReverseSweep        bool                `json:"reverseSweep,omitempty" bson:"reverse_sweep,omitempty"`
	ReverseSweepAttempt bool                `json:"reverseSweepAttempt,omitempty" bson:"reverse_sweep_attempt,omitempty"`
	Games               []*GameOverview     `json:"games,omitempty" bson:"games,omitempty"`
}

// GameOverview .
type GameOverview struct {
	ID            *primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Blue          float64             `json:"blue" bson:"blue"`
	Orange        float64             `json:"orange" bson:"orange"`
	Duration      int                 `json:"duration,omitempty" bson:"duration,omitempty"`
	BallchasingID string              `json:"ballchasing,omitempty" bson:"ballchasing,omitempty"`
}

// Format .
type Format struct {
	Type   string `json:"type,omitempty" bson:"type,omitempty"`
	Length int    `json:"length,omitempty" bson:"length,omitempty"`
}

// MatchSide .
type MatchSide struct {
	Score   int           `json:"score,omitempty" bson:"score,omitempty"`
	Winner  bool          `json:"winner,omitempty" bson:"winner,omitempty"`
	Team    *TeamInfo     `json:"team,omitempty" bson:"team,omitempty"`
	Players []*PlayerInfo `json:"players,omitempty" bson:"players,omitempty"`
}
