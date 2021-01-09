package pipelines

import (
	"github.com/octanegg/zsr/octane"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// PlayerTeams .
func PlayerTeams(filter bson.M) *Pipeline {
	pipeline := New(
		Match(filter),
		Sort("game.date", false),
		Group(bson.M{
			"_id": "$team.team._id",
			"team": bson.M{
				"$first": "$team.team",
			},
			"games": bson.M{
				"$sum": 1,
			},
			"start": bson.M{
				"$first": "$game",
			},
			"end": bson.M{
				"$last": "$game",
			},
		}),
		Match(bson.M{"games": bson.M{"$gte": 30}}),
		Sort("start.date", false),
	)

	return &Pipeline{
		Pipeline: pipeline,
		Decode: func(cursor *mongo.Cursor) (interface{}, error) {
			var team struct {
				Team  *octane.Team `json:"team,omitempty" bson:"team,omitempty"`
				Games int          `json:"games,omitempty" bson:"games,omitempty"`
				Start *octane.Game `json:"start,omitempty" bson:"start,omitempty"`
				End   *octane.Game `json:"end,omitempty" bson:"end,omitempty"`
			}
			if err := cursor.Decode(&team); err != nil {
				return nil, err
			}
			return team, nil
		},
	}
}
