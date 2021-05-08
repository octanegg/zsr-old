package octane

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Game .
type Game struct {
	ID              *primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	OctaneID        string              `json:"octane_id,omitempty" bson:"octane_id,omitempty"`
	Number          int                 `json:"number,omitempty" bson:"number,omitempty"`
	Match           *Match              `json:"match,omitempty" bson:"match,omitempty"`
	Map             *Map                `json:"map,omitempty" bson:"map,omitempty"`
	Duration        int                 `json:"duration,omitempty" bson:"duration,omitempty"`
	Date            *time.Time          `json:"date,omitempty" bson:"date,omitempty"`
	Blue            *GameSide           `json:"blue,omitempty" bson:"blue,omitempty"`
	Orange          *GameSide           `json:"orange,omitempty" bson:"orange,omitempty"`
	BallchasingID   string              `json:"ballchasing,omitempty" bson:"ballchasing,omitempty"`
	FlipBallchasing bool                `json:"flipBallchasing,omitempty" bson:"flip_ballchasing,omitempty"`
}

// GameSide .
type GameSide struct {
	Winner  bool          `json:"winner,omitempty" bson:"winner,omitempty"`
	Team    *TeamInfo     `json:"team,omitempty" bson:"team,omitempty"`
	Players []*PlayerInfo `json:"players,omitempty" bson:"players,omitempty"`
}

// Map .
type Map struct {
	ID   string `json:"id,omitempty" bson:"id,omitempty"`
	Name string `json:"name,omitempty" bson:"name,omitempty"`
}

// PlayerInfo .
type PlayerInfo struct {
	Player   *Player        `json:"player,omitempty" bson:"player,omitempty"`
	Car      *Car           `json:"car,omitempty" bson:"car,omitempty"`
	Camera   *Camera        `json:"camera,omitempty" bson:"camera,omitempty"`
	Stats    *PlayerStats   `json:"stats,omitempty" bson:"stats,omitempty"`
	Advanced *AdvancedStats `json:"advanced,omitempty" bson:"advanced,omitempty"`
}

// Car .
type Car struct {
	ID   int    `json:"id,omitempty" bson:"id,omitempty"`
	Name string `json:"name,omitempty" bson:"name,omitempty"`
}

// Camera .
type Camera struct {
	Fov             int     `json:"fov" bson:"fov"`
	Height          int     `json:"height" bson:"height"`
	Pitch           int     `json:"pitch" bson:"pitch"`
	Distance        int     `json:"distance" bson:"distance"`
	Stiffness       float64 `json:"stiffness" bson:"stiffness"`
	SwivelSpeed     float64 `json:"swivelSpeed" bson:"swivel_speed"`
	TransitionSpeed float64 `json:"transitionSpeed" bson:"transition_speed"`
}

// TeamInfo .
type TeamInfo struct {
	Team  *Team      `json:"team,omitempty" bson:"team,omitempty"`
	Stats *TeamStats `json:"stats,omitempty" bson:"stats,omitempty"`
}
