package pipelines

import (
	"time"

	"github.com/octanegg/zsr/octane"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// TeamAggregate .
func TeamAggregate(filter bson.M, group interface{}, having bson.M) *Pipeline {
	pipeline := New(
		Match(filter),
		Sort("game.date", false),
		Group(bson.M{
			"_id": group,
			"team": bson.M{
				"$first": "$team.team",
			},
			"events": bson.M{
				"$addToSet": "$game.match.event",
			},
			"opponents": bson.M{
				"$addToSet": "$opponent.team",
			},
			"start_date": bson.M{
				"$first": "$game.date",
			},
			"end_date": bson.M{
				"$last": "$game.date",
			},
			"players": bson.M{
				"$addToSet": "$player",
			},
			"games": bson.M{
				"$sum": bson.M{
					"$divide": bson.A{1, "$game.match.event.mode"},
				},
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
						}, bson.M{
							"$divide": bson.A{1, "$game.match.event.mode"},
						}, 0,
					},
				},
			},
			"score_total": bson.M{
				"$sum": bson.M{
					"$divide": bson.A{"$stats.team.core.score", "$game.match.event.mode"},
				},
			},
			"goals_total": bson.M{
				"$sum": bson.M{
					"$divide": bson.A{"$stats.team.core.goals", "$game.match.event.mode"},
				},
			},
			"assists_total": bson.M{
				"$sum": bson.M{
					"$divide": bson.A{"$stats.team.core.assists", "$game.match.event.mode"},
				},
			},
			"saves_total": bson.M{
				"$sum": bson.M{
					"$divide": bson.A{"$stats.team.core.saves", "$game.match.event.mode"},
				},
			},
			"shots_total": bson.M{
				"$sum": bson.M{
					"$divide": bson.A{"$stats.team.core.shots", "$game.match.event.mode"},
				},
			},
			"opp_score_total": bson.M{
				"$sum": bson.M{
					"$divide": bson.A{"$stats.opponent.core.score", "$game.match.event.mode"},
				},
			},
			"opp_goals_total": bson.M{
				"$sum": bson.M{
					"$divide": bson.A{"$stats.opponent.core.goals", "$game.match.event.mode"},
				},
			},
			"opp_assists_total": bson.M{
				"$sum": bson.M{
					"$divide": bson.A{"$stats.opponent.core.assists", "$game.match.event.mode"},
				},
			},
			"opp_saves_total": bson.M{
				"$sum": bson.M{
					"$divide": bson.A{"$stats.opponent.core.saves", "$game.match.event.mode"},
				},
			},
			"opp_shots_total": bson.M{
				"$sum": bson.M{
					"$divide": bson.A{"$stats.opponent.core.shots", "$game.match.event.mode"},
				},
			},
		}),
		Match(having),
		Project(bson.M{
			"_id":        "$_id",
			"team":       "$team",
			"players":    "$players",
			"events":     "$events",
			"opponents":  "$opponents",
			"start_date": "$start_date",
			"end_date":   "$end_date",
			"games": bson.M{
				"$toInt": "$games",
			},
			"wins": bson.M{
				"$toInt": "$wins",
			},
			"win_percentage": bson.M{
				"$divide": bson.A{
					"$wins", "$games",
				},
			},
			"stats": bson.M{
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
					"savePercentage": bson.M{
						"$divide": bson.A{
							"$saves_total", bson.M{
								"$sum": bson.A{"$saves_total", "$opp_goals_total"},
							},
						},
					},
				},
				"against": bson.M{
					"score": bson.M{
						"$divide": bson.A{"$opp_score_total", "$games"},
					},
					"goals": bson.M{
						"$divide": bson.A{"$opp_goals_total", "$games"},
					},
					"assists": bson.M{
						"$divide": bson.A{"$opp_assists_total", "$games"},
					},
					"saves": bson.M{
						"$divide": bson.A{"$opp_saves_total", "$games"},
					},
					"shots": bson.M{
						"$divide": bson.A{"$opp_shots_total", "$games"},
					},
					"shootingPercentage": bson.M{
						"$cond": bson.A{
							bson.M{"$eq": bson.A{"$opp_shots_total", 0}},
							1,
							bson.M{
								"$divide": bson.A{"$opp_goals_total", "$opp_shots_total"},
							},
						},
					},
					"savePercentage": bson.M{
						"$divide": bson.A{
							"$opp_saves_total", bson.M{
								"$sum": bson.A{"$opp_saves_total", "$goals_total"},
							},
						},
					},
				},
				"totals": bson.M{
					"score":   "$score_total",
					"goals":   "$goals_total",
					"assists": "$assists_total",
					"saves":   "$saves_total",
					"shots":   "$shots_total",
				},
				"differentials": bson.M{
					"score": bson.M{
						"$divide": bson.A{
							bson.M{
								"$subtract": bson.A{"$score_total", "$opp_score_total"},
							},
							"$games",
						},
					},
					"goals": bson.M{
						"$divide": bson.A{
							bson.M{
								"$subtract": bson.A{"$goals_total", "$opp_goals_total"},
							},
							"$games",
						},
					},
					"assists": bson.M{
						"$divide": bson.A{
							bson.M{
								"$subtract": bson.A{"$assists_total", "$opp_assists_total"},
							},
							"$games",
						},
					},
					"saves": bson.M{
						"$divide": bson.A{
							bson.M{
								"$subtract": bson.A{"$saves_total", "$opp_saves_total"},
							},
							"$games",
						},
					},
					"shots": bson.M{
						"$divide": bson.A{
							bson.M{
								"$subtract": bson.A{"$shots_total", "$opp_shots_total"},
							},
							"$games",
						},
					},
				},
			},
		}),
		Sort("end_date", true),
	)

	return &Pipeline{
		Pipeline: pipeline,
		Decode: func(cursor *mongo.Cursor) (interface{}, error) {
			var team struct {
				Team          *octane.Team     `json:"team" bson:"team"`
				Players       []*octane.Player `json:"players" bson:"players"`
				Events        []*octane.Event  `json:"events" bson:"events"`
				Opponents     []*octane.Team   `json:"opponents" bson:"opponents"`
				StartDate     *time.Time       `json:"startDate" bson:"start_date"`
				EndDate       *time.Time       `json:"endDate" bson:"end_date"`
				Games         int              `json:"games" bson:"games"`
				Wins          int              `json:"wins" bson:"wins"`
				WinPercentage float64          `json:"winPercentage" bson:"win_percentage"`
				Stats         struct {
					Averages struct {
						Score              float64 `json:"score" bson:"score"`
						Goals              float64 `json:"goals" bson:"goals"`
						Assists            float64 `json:"assists" bson:"assists"`
						Saves              float64 `json:"saves" bson:"saves"`
						Shots              float64 `json:"shots" bson:"shots"`
						ShootingPercentage float64 `json:"shootingPercentage" bson:"shootingPercentage"`
						SavePercentage     float64 `json:"savePercentage" bson:"savePercentage"`
					} `json:"averages" bson:"averages"`
					Against struct {
						Score              float64 `json:"score" bson:"score"`
						Goals              float64 `json:"goals" bson:"goals"`
						Assists            float64 `json:"assists" bson:"assists"`
						Saves              float64 `json:"saves" bson:"saves"`
						Shots              float64 `json:"shots" bson:"shots"`
						ShootingPercentage float64 `json:"shootingPercentage" bson:"shootingPercentage"`
						SavePercentage     float64 `json:"savePercentage" bson:"savePercentage"`
					} `json:"against" bson:"against"`
					Totals struct {
						Score   float64 `json:"score" bson:"score"`
						Goals   float64 `json:"goals" bson:"goals"`
						Assists float64 `json:"assists" bson:"assists"`
						Saves   float64 `json:"saves" bson:"saves"`
						Shots   float64 `json:"shots" bson:"shots"`
					} `json:"totals" bson:"totals"`
					Differentials struct {
						Score   float64 `json:"score" bson:"score"`
						Goals   float64 `json:"goals" bson:"goals"`
						Assists float64 `json:"assists" bson:"assists"`
						Saves   float64 `json:"saves" bson:"saves"`
						Shots   float64 `json:"shots" bson:"shots"`
					} `json:"differentials" bson:"differentials"`
				} `json:"stats" bson:"stats"`
			}
			if err := cursor.Decode(&team); err != nil {
				return nil, err
			}
			return team, nil
		},
	}
}
