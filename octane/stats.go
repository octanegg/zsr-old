package octane

import (
	"time"

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
	Score   float64    `json:"score,omitempty" bson:"score,omitempty"`
	Winner  bool       `json:"winner,omitempty" bson:"winner,omitempty"`
	Team    *Team      `json:"team,omitempty" bson:"team,omitempty"`
	Stats   *TeamStats `json:"stats,omitempty" bson:"stats,omitempty"`
	Players []*Player  `json:"players,omitempty" bson:"players,omitempty"`
}

// PlayerAggregateStats .
type PlayerAggregateStats struct {
	ID        *primitive.ObjectID   `json:"_id,omitempty" bson:"_id,omitempty"`
	Player    *Player               `json:"player,omitempty" bson:"player,omitempty"`
	Events    []*Event              `json:"events,omitempty" bson:"events,omitempty"`
	Teams     []*Team               `json:"teams,omitempty" bson:"teams,omitempty"`
	Opponents []*Team               `json:"teams,omitempty" bson:"opponents,omitempty"`
	StartDate *time.Time            `json:"startDate,omitempty" bson:"start_date,omitempty"`
	EndDate   *time.Time            `json:"endDate,omitempty" bson:"end_date,omitempty"`
	Amounts   *AggregateAmounts     `json:"amounts,omitempty" bson:"amounts,omitempty"`
	Stats     *AggregatePlayerStats `json:"stats,omitempty" bson:"stats,omitempty"`
}

// AggregatePlayerStats .
type AggregatePlayerStats struct {
	Core        *PlayerCore        `json:"core" bson:"core"`
	Boost       *PlayerBoost       `json:"boost" bson:"boost"`
	Movement    *PlayerMovement    `json:"movement" bson:"movement"`
	Positioning *PlayerPositioning `json:"positioning" bson:"positioning"`
	Demolitions *PlayerDemolitions `json:"demo" bson:"demo"`
	Advanced    *AdvancedStats     `json:"advanced" bson:"advanced"`
}

// TeamAggregateStats .
type TeamAggregateStats struct {
	ID        *primitive.ObjectID   `json:"_id,omitempty" bson:"_id,omitempty"`
	Team      *Team                 `json:"team,omitempty" bson:"team,omitempty"`
	Players   []*Player             `json:"players,omitempty" bson:"players,omitempty"`
	Events    []*Event              `json:"events,omitempty" bson:"events,omitempty"`
	Opponents []*Team               `json:"teams,omitempty" bson:"opponents,omitempty"`
	StartDate *time.Time            `json:"startDate,omitempty" bson:"start_date,omitempty"`
	EndDate   *time.Time            `json:"endDate,omitempty" bson:"end_date,omitempty"`
	Amounts   *AggregateAmounts     `json:"amounts,omitempty" bson:"amounts,omitempty"`
	Stats     *AggregatePlayerStats `json:"stats,omitempty" bson:"stats,omitempty"`
}

// AggregateTeamStats .
type AggregateTeamStats struct {
	Core        *TeamCore        `json:"core" bson:"core"`
	Against     *TeamCore        `json:"against" bson:"against"`
	Boost       *TeamBoost       `json:"boost" bson:"boost"`
	Movement    *TeamMovement    `json:"movement" bson:"movement"`
	Positioning *TeamPositioning `json:"positioning" bson:"positioning"`
	Demolitions *TeamDemolitions `json:"demo" bson:"demo"`
	Ball        *TeamBall        `json:"ball" bson:"ball"`
}

// AggregateAmounts .
type AggregateAmounts struct {
	Games        int `json:"games,omitempty" bson:"games,omitempty"`
	Matches      int `json:"matches,omitempty" bson:"matches,omitempty"`
	GameReplays  int `json:"gameReplays,omitempty" bson:"game_replays,omitempty"`
	MatchReplays int `json:"matchReplays,omitempty" bson:"match_replays,omitempty"`
}
