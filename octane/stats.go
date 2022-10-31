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

// Record .
type Record struct {
	ID                 *primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Game               *Game               `json:"game,omitempty" bson:"game,omitempty"`
	Team               *RecordSide         `json:"team,omitempty" bson:"team,omitempty"`
	Opponent           *RecordSide         `json:"opponent,omitempty" bson:"opponent,omitempty"`
	Player             *Player             `json:"player,omitempty" bson:"player,omitempty"`
	Duration           float64             `json:"duration,omitempty" bson:"duration,omitempty"`
	Stat               string              `json:"stat,omitempty" bson:"stat,omitempty"`
	PlayerValue        float64             `json:"playerValue" bson:"player_value"`
	PlayerMatchValue   float64             `json:"playerMatchValue" bson:"player_match_value"`
	PlayerMatchAverage float64             `json:"playerMatchAverage" bson:"player_match_average"`
	TeamValue          float64             `json:"teamValue" bson:"team_value"`
	TeamMatchValue     float64             `json:"teamMatchValue" bson:"team_match_value"`
	TeamMatchAverage   float64             `json:"teamMatchAverage" bson:"team_match_average"`
	GameValue          float64             `json:"gameValue" bson:"game_value"`
	GameAverage        float64             `json:"gameAverage" bson:"game_average"`
	GameDifferential   float64             `json:"gameDifferential" bson:"game_differential"`
	MatchValue         float64             `json:"matchValue" bson:"match_value"`
	MatchDifferential  float64             `json:"matchDifferential" bson:"match_differential"`
}

// RecordSide .
type RecordSide struct {
	Score       float64   `json:"score,omitempty" bson:"score,omitempty"`
	MatchWinner bool      `json:"matchWinner,omitempty" bson:"match_winner,omitempty"`
	Winner      bool      `json:"winner,omitempty" bson:"winner,omitempty"`
	Team        *Team     `json:"team,omitempty" bson:"team,omitempty"`
	Players     []*Player `json:"players,omitempty" bson:"players,omitempty"`
}
