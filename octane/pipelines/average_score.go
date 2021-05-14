package pipelines

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func AverageScore(goals, assists, saves, shots int) *Pipeline {
	pipeline := New(
		Match(bson.M{
			"player.stats.core.goals":   goals,
			"player.stats.core.assists": assists,
			"player.stats.core.saves":   saves,
			"player.stats.core.shots":   shots,
		}),
		Group(bson.M{
			"_id": 0,
			"score": bson.M{
				"$avg": "$player.stats.core.score",
			},
		}),
	)

	return &Pipeline{
		Pipeline: pipeline,
		Decode: func(cursor *mongo.Cursor) (interface{}, error) {
			var score struct {
				Score float64 `json:"score" bson:"score"`
			}
			if err := cursor.Decode(&score); err != nil {
				return nil, err
			}
			return score, nil
		},
	}
}
