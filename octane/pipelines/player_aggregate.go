package pipelines

import (
	"github.com/octanegg/zsr/octane"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// PlayerAggregate .
func PlayerAggregate(filter bson.M, group interface{}, having bson.M) *Pipeline {
	pipeline := New(
		Match(filter),
		Group(bson.M{
			"_id": group,
			"player": bson.M{
				"$first": "$player",
			},
			"events": bson.M{
				"$addToSet": "$game.match.event",
			},
			"teams": bson.M{
				"$addToSet": "$team.team",
			},
			"games": bson.M{
				"$sum": 1,
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
			"rating_total": bson.M{
				"$sum": "$stats.player.core.rating",
			},
			"team_goals_total": bson.M{
				"$sum": "$stats.team.core.goals",
			},
		}),
		Match(having),
		Project(bson.M{
			"_id":    "$_id",
			"player": "$player",
			"teams":  "$teams",
			"events": "$events",
			"games":  "$games",
			"wins":   "$wins",
			"win_percentage": bson.M{
				"$divide": bson.A{
					"$wins", "$games",
				},
			},
			"averages": bson.M{
				"score": bson.M{
					"$divide": bson.A{"$score_total", "$games"},
				},
				"goals": bson.M{
					"$divide": bson.A{"$goals_total", "$games"},
				},
				"assists": bson.M{
					"$divide": bson.A{"$assists_total", "$games"},
				},
				"saves": bson.M{
					"$divide": bson.A{"$saves_total", "$games"},
				},
				"shots": bson.M{
					"$divide": bson.A{"$shots_total", "$games"},
				},
				"shootingPercentage": bson.M{
					"$cond": bson.A{
						bson.M{"$eq": bson.A{"$shots_total", 0}},
						1,
						bson.M{
							"$divide": bson.A{"$goals_total", "$shots_total"},
						},
					},
				},
				"goalParticipation": bson.M{
					"$cond": bson.A{
						bson.M{"$eq": bson.A{"$team_goals_total", 0}},
						1,
						bson.M{
							"$divide": bson.A{
								bson.M{
									"$add": bson.A{"$goals_total", "$assists_total"},
								},
								"$team_goals_total",
							},
						},
					},
				},
				"rating": bson.M{
					"$divide": bson.A{"$rating_total", "$games"},
				},
			},
		}),
	)

	return &Pipeline{
		Pipeline: pipeline,
		Decode: func(cursor *mongo.Cursor) (interface{}, error) {
			var player struct {
				Player        *octane.Player  `json:"player" bson:"player"`
				Events        []*octane.Event `json:"events" bson:"events"`
				Teams         []*octane.Team  `json:"teams" bson:"teams"`
				Games         int             `json:"games" bson:"games"`
				Wins          int             `json:"wins" bson:"wins"`
				WinPercentage float64         `json:"win_percentage" bson:"win_percentage"`
				Averages      struct {
					Score              float64 `json:"score" bson:"score"`
					Goals              float64 `json:"goals" bson:"goals"`
					Assists            float64 `json:"assists" bson:"assists"`
					Saves              float64 `json:"saves" bson:"saves"`
					Shots              float64 `json:"shots" bson:"shots"`
					ShootingPercentage float64 `json:"shootingPercentage" bson:"shootingPercentage"`
					GoalParticipation  float64 `json:"goalParticipation" bson:"goalParticipation"`
					Rating             float64 `json:"rating" bson:"rating"`
				} `json:"averages" bson:"averages"`
			}
			if err := cursor.Decode(&player); err != nil {
				return nil, err
			}
			return player, nil
		},
	}
}
