package octane

import (
	"time"

	"github.com/octanegg/zsr/ballchasing"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Game .
type Game struct {
	ID            *primitive.ObjectID `json:"_id" bson:"_id"`
	OctaneID      string              `json:"octane_id" bson:"octane_id"`
	Number        int                 `json:"number" bson:"number"`
	Match         *Match              `json:"match" bson:"match"`
	Map           string              `json:"map" bson:"map"`
	Duration      int                 `json:"duration" bson:"duration"`
	Date          *time.Time          `json:"date,omitempty" bson:"date,omitempty"`
	Blue          *GameSide           `json:"blue" bson:"blue"`
	Orange        *GameSide           `json:"orange" bson:"orange"`
	BallchasingID string              `json:"ballchasing,omitempty" bson:"ballchasing,omitempty"`
}

// GameSide .
type GameSide struct {
	Goals   int            `json:"goals" bson:"goals"`
	Winner  bool           `json:"winner" bson:"winner"`
	Team    *Team          `json:"team" bson:"team"`
	Players []*PlayerStats `json:"players" bson:"players"`
}

// PlayerStats .
type PlayerStats struct {
	Player *Player                  `json:"player" bson:"player"`
	Stats  *ballchasing.PlayerStats `json:"stats" bson:"stats"`
	Rating float64                  `json:"rating" bson:"rating"`
}
