package pipelines

import (
	"github.com/octanegg/zsr/octane"
)

// Player .
type Player struct {
	Player        *octane.Player  `json:"player" bson:"player,omitempty"`
	Team          *octane.Team    `json:"team,omitempty" bson:"team,omitempty"`
	Event         *octane.Event   `json:"event,omitempty" bson:"event,omitempty"`
	Games         int             `json:"games" bson:"games"`
	Wins          int             `json:"wins" bson:"wins"`
	WinPercentage float64         `json:"win_percentage" bson:"win_percentage"`
	Totals        *PlayerTotals   `json:"totals" bson:"totals"`
	Averages      *PlayerAverages `json:"averages" bson:"averages"`
}

// PlayerTotals .
type PlayerTotals struct {
	Score   int     `json:"score" bson:"score"`
	Goals   int     `json:"goals" bson:"goals"`
	Assists int     `json:"assists" bson:"assists"`
	Saves   int     `json:"saves" bson:"saves"`
	Shots   int     `json:"shots" bson:"shots"`
	Rating  float64 `json:"rating" bson:"rating"`
}

// PlayerAverages .
type PlayerAverages struct {
	Score   float64 `json:"score" bson:"score"`
	Goals   float64 `json:"goals" bson:"goals"`
	Assists float64 `json:"assists" bson:"assists"`
	Saves   float64 `json:"saves" bson:"saves"`
	Shots   float64 `json:"shots" bson:"shots"`
	Rating  float64 `json:"rating" bson:"rating"`
}
