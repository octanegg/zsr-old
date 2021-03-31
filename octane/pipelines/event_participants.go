package pipelines

import (
	"github.com/octanegg/zsr/octane"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// EventParticipants .
func EventParticipants(event *primitive.ObjectID) *Pipeline {
	pipeline := New(
		Match(bson.M{
			"game.match.event._id": event,
		}),
		Group(bson.M{
			"_id": bson.M{
				"event": "$game.match.event._id",
				"team":  "$team.team._id",
			},
			"team": bson.M{
				"$first": "$team.team",
			},
			"players": bson.M{
				"$addToSet": "$player.player",
			},
		}),
		Project(bson.M{
			"team":    "$team",
			"players": "$players",
		}),
		Sort("team.name", false),
		Limit(25),
	)

	return &Pipeline{
		Pipeline: pipeline,
		Decode: func(cursor *mongo.Cursor) (interface{}, error) {
			var participants struct {
				Team    *octane.Team     `json:"team,omitempty" bson:"team,omitempty"`
				Players []*octane.Player `json:"players,omitempty" bson:"players,omitempty"`
			}
			if err := cursor.Decode(&participants); err != nil {
				return nil, err
			}
			return participants, nil
		},
	}
}
