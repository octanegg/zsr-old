package pipelines

import (
	"github.com/octanegg/zsr/octane"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// PlayerAggregate .
func PlayerAggregate(filter bson.M, having bson.M) *Pipeline {
	pipeline := New(
		Match(filter),
		Group(bson.M{
			"_id": "$player._id",
			"player": bson.M{
				"$first": "$player",
			},
			"games": bson.M{
				"$sum": 1,
			},
			"wins": bson.M{
				"$sum": bson.M{
					"$cond": bson.A{
						bson.M{
							"$eq": bson.A{
								"$winner", true,
							},
						}, 1, 0,
					},
				},
			},
			"score_total": bson.M{
				"$sum": "$stats.core.score",
			},
			"goals_total": bson.M{
				"$sum": "$stats.core.goals",
			},
			"assists_total": bson.M{
				"$sum": "$stats.core.assists",
			},
			"saves_total": bson.M{
				"$sum": "$stats.core.saves",
			},
			"shots_total": bson.M{
				"$sum": "$stats.core.shots",
			},
			"score_avg": bson.M{
				"$avg": "$stats.core.score",
			},
			"goals_avg": bson.M{
				"$avg": "$stats.core.goals",
			},
			"assists_avg": bson.M{
				"$avg": "$stats.core.assists",
			},
			"saves_avg": bson.M{
				"$avg": "$stats.core.saves",
			},
			"shots_avg": bson.M{
				"$avg": "$stats.core.shots",
			},
			"rating_avg": bson.M{
				"$avg": "$stats.core.rating",
			},
		}),
		Match(having),
		Project(bson.M{
			"_id":    "$_id",
			"player": "$player",
			"team":   "$team",
			"event":  "$event",
			"games":  "$games",
			"wins":   "$wins",
			"win_percentage": bson.M{
				"$divide": bson.A{
					"$wins", "$games",
				},
			},
			"totals": bson.M{
				"score":   "$score_total",
				"goals":   "$goals_total",
				"assists": "$assists_total",
				"saves":   "$saves_total",
				"shots":   "$shots_total",
			},
			"averages": bson.M{
				"score":   "$score_avg",
				"goals":   "$goals_avg",
				"assists": "$assists_avg",
				"saves":   "$saves_avg",
				"shots":   "$shots_avg",
				"rating":  "$rating_avg",
			},
		}),
	)

	return &Pipeline{
		Pipeline: pipeline,
		Decode: func(cursor *mongo.Cursor) (interface{}, error) {
			var player struct {
				Player        *octane.Player `json:"player" bson:"player,omitempty"`
				Games         int            `json:"games" bson:"games"`
				Wins          int            `json:"wins" bson:"wins"`
				WinPercentage float64        `json:"win_percentage" bson:"win_percentage"`
				Totals        struct {
					Score   int     `json:"score" bson:"score"`
					Goals   int     `json:"goals" bson:"goals"`
					Assists int     `json:"assists" bson:"assists"`
					Saves   int     `json:"saves" bson:"saves"`
					Shots   int     `json:"shots" bson:"shots"`
					Rating  float64 `json:"rating" bson:"rating"`
				} `json:"totals" bson:"totals"`
				Averages struct {
					Score   float64 `json:"score" bson:"score"`
					Goals   float64 `json:"goals" bson:"goals"`
					Assists float64 `json:"assists" bson:"assists"`
					Saves   float64 `json:"saves" bson:"saves"`
					Shots   float64 `json:"shots" bson:"shots"`
					Rating  float64 `json:"rating" bson:"rating"`
				} `json:"averages" bson:"averages"`
			}
			if err := cursor.Decode(&player); err != nil {
				return nil, err
			}
			return player, nil
		},
	}
}
