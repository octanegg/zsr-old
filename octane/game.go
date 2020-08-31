package octane

import "go.mongodb.org/mongo-driver/bson/primitive"

// Game .
type Game struct {
	ID       primitive.ObjectID `json:"id" bson:"_id"`
	MatchID  primitive.ObjectID `json:"match" bson:"match"`
	EventID  primitive.ObjectID `json:"event" bson:"event"`
	Map      string             `json:"map" bson:"map"`
	Duration int                `json:"duration" bson:"duration"`
	Mode     int                `json:"mode" bson:"mode"`
	Blue     GameTeam           `json:"blue" bson:"blue"`
	Orange   GameTeam           `json:"orange" bson:"orange"`
}

// GameTeam .
type GameTeam struct {
	Goals   int          `json:"goals" bson:"goals"`
	Winner  bool         `json:"winner" bson:"winner"`
	Team    Team         `json:"team" bson:"team"`
	Players []GamePlayer `json:"players" bson:"players"`
}

// GamePlayer .
type GamePlayer struct {
	Player Player `json:"player" bson:"player"`
	Stats  Stats  `json:"stats" bson:"stats"`
}
