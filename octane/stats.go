package octane

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Statline .
type Statline struct {
	ID       *primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Game     *Game               `json:"game,omitempty" bson:"game,omitempty"`
	Team     *StatlineSide       `json:"team,omitempty" bson:"team,omitempty"`
	Opponent *StatlineSide       `json:"opponent,omitempty" bson:"opponent,omitempty"`
	Player   *PlayerInfo         `json:"player,omitempty" bson:"player,omitempty"`
}

// StatlineSide .
type StatlineSide struct {
	Score       float64    `json:"score,omitempty" bson:"score,omitempty"`
	MatchWinner bool       `json:"matchWinner,omitempty" bson:"match_winner,omitempty"`
	Winner      bool       `json:"winner,omitempty" bson:"winner,omitempty"`
	Team        *Team      `json:"team,omitempty" bson:"team,omitempty"`
	Stats       *TeamStats `json:"stats,omitempty" bson:"stats,omitempty"`
	Players     []*Player  `json:"players,omitempty" bson:"players,omitempty"`
}
