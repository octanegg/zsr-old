package octane

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Match .
type Match struct {
	ID       primitive.ObjectID   `json:"id" bson:"_id"`
	EventID  primitive.ObjectID   `json:"event" bson:"event"`
	Stage    int                  `json:"stage" bson:"stage"`
	Substage int                  `json:"substage" bson:"substage"`
	Date     string               `json:"date" bson:"date"`
	Format   string               `json:"format" bson:"format"`
	Blue     MatchTeam            `json:"blue" bson:"blue"`
	Orange   MatchTeam            `json:"orange" bson:"orange"`
	Games    []primitive.ObjectID `json:"games" bson:"games"`
	Mode     int                  `json:"mode" bson:"mode"`
}

// MatchTeam .
type MatchTeam struct {
	Score  int  `json:"score" bson:"score"`
	Winner bool `json:"winner" bson:"winner"`
	Team   Team `json:"team" bson:"team"`
}
