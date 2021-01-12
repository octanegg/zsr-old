package pipelines

import (
	"github.com/octanegg/zsr/octane"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// TeamAggregate .
func TeamAggregate(filter bson.M, having bson.M) *Pipeline {
	pipeline := New(
		Match(filter),
		Group(bson.M{
			"_id": "$team.team._id",
			"team": bson.M{
				"$first": "$team.team",
			},
			"games": bson.M{
				"$sum": 1,
			},
			"mode": bson.M{
				"$first": "$game.match.event.mode",
			},
			"wins": bson.M{
				"$sum": bson.M{
					"$cond": bson.A{
						bson.M{
							"$eq": bson.A{
								"$team.winner", true,
							},
						}, 1, 0,
					},
				},
			},
			"score_total": bson.M{
				"$sum": "$stats.player.core.score",
			},
			"goals_total": bson.M{
				"$sum": "$stats.player.core.goals",
			},
			"assists_total": bson.M{
				"$sum": "$stats.player.core.assists",
			},
			"saves_total": bson.M{
				"$sum": "$stats.player.core.saves",
			},
			"shots_total": bson.M{
				"$sum": "$stats.player.core.shots",
			},
			"score_avg": bson.M{
				"$avg": "$stats.player.core.score",
			},
			"goals_avg": bson.M{
				"$avg": "$stats.player.core.goals",
			},
			"assists_avg": bson.M{
				"$avg": "$stats.player.core.assists",
			},
			"saves_avg": bson.M{
				"$avg": "$stats.player.core.saves",
			},
			"shots_avg": bson.M{
				"$avg": "$stats.player.core.shots",
			},
		}),
		Match(having),
		Project(bson.M{
			"_id":  "$_id",
			"team": "$team",
			"games": bson.M{
				"$toInt": bson.M{
					"$divide": bson.A{"$games", "$mode"},
				},
			},
			"wins": bson.M{
				"$toInt": bson.M{
					"$divide": bson.A{"$wins", "$mode"},
				},
			},
			"win_percentage": bson.M{
				"$divide": bson.A{
					"$wins", "$games",
				},
			},
			"averages": bson.M{
				"score":   "$score_avg",
				"goals":   "$goals_avg",
				"assists": "$assists_avg",
				"saves":   "$saves_avg",
				"shots":   "$shots_avg",
				"shootingPercentage": bson.M{
					"$cond": bson.A{
						bson.M{"$eq": bson.A{"$shots_total", 0}},
						1,
						bson.M{
							"$divide": bson.A{"$goals_total", "$shots_total"},
						},
					},
				},
			},
		}),
	)

	return &Pipeline{
		Pipeline: pipeline,
		Decode: func(cursor *mongo.Cursor) (interface{}, error) {
			var team struct {
				Team          *octane.Team `json:"team" bson:"team,omitempty"`
				Games         int          `json:"games" bson:"games"`
				Wins          int          `json:"wins" bson:"wins"`
				WinPercentage float64      `json:"win_percentage" bson:"win_percentage"`
				Averages      struct {
					Score              float64 `json:"score" bson:"score"`
					Goals              float64 `json:"goals" bson:"goals"`
					Assists            float64 `json:"assists" bson:"assists"`
					Saves              float64 `json:"saves" bson:"saves"`
					Shots              float64 `json:"shots" bson:"shots"`
					ShootingPercentage float64 `json:"shootingPercentage" bson:"shootingPercentage"`
				} `json:"averages" bson:"averages"`
			}
			if err := cursor.Decode(&team); err != nil {
				return nil, err
			}
			return team, nil
		},
	}
}
