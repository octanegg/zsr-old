package pipelines

import (
	"fmt"
	"strings"
	"time"

	"github.com/octanegg/zsr/octane/stats"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// TeamMetrics .
func TeamMetrics(filter bson.M, _stats []string) *Pipeline {
	_group := bson.M{
		"_id": bson.M{
			"day": bson.M{
				"$dateToString": bson.M{
					"date":   "$game.date",
					"format": "%Y-%m-%d",
				},
			},
		},
		"date": bson.M{
			"$first": "$game.date",
		},
		"games": bson.M{
			"$sum": bson.M{
				"$divide": bson.A{1, "$game.match.event.mode"},
			},
		},
		"game_replays": bson.M{
			"$sum": bson.M{
				"$cond": bson.A{
					bson.M{
						"$ifNull": bson.A{"$game.ballchasing", false},
					}, bson.M{
						"$divide": bson.A{1, "$game.match.event.mode"},
					}, 0,
				},
			},
		},
		"game_wins": bson.M{
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
		"matches": bson.M{
			"$addToSet": "$game.match._id",
		},
		"match_replays": bson.M{
			"$addToSet": bson.M{
				"$cond": bson.A{
					bson.M{
						"$ifNull": bson.A{"$game.ballchasing", false},
					}, "$game.match._id", bson.TypeNull,
				},
			},
		},
		"match_wins": bson.M{
			"$addToSet": bson.M{
				"$cond": bson.A{
					bson.M{
						"$eq": bson.A{
							"$team.match_winner", true,
						},
					}, "$game.match._id", bson.TypeNull,
				},
			},
		},
	}

	for _, stat := range _stats {
		if stat == "shootingPercentage" {
			_group["goals"] = bson.M{
				"$sum": "$team.stats.core.goals",
			}
			_group["shots"] = bson.M{
				"$sum": "$team.stats.core.shots",
			}
		} else if stat == "shootingPercentageAgainst" {
			_group["opponent_goals"] = bson.M{
				"$sum": "$opponent.stats.core.goals",
			}
			_group["opponent_shots"] = bson.M{
				"$sum": "$opponent.stats.core.shots",
			}
		} else {
			stat := stat

			var statType string
			if strings.Contains(stat, "Against") {
				stat = stat[:len(stat)-7]
				statType = "against"
			} else if strings.Contains(stat, "Differential") {
				stat = stat[:len(stat)-12]
				statType = "differential"
			}

			groupName, statMapping := stats.TeamStatMapping(stat)
			if statMapping == "" {
				continue
			}

			if statType == "against" {
				_group[fmt.Sprintf("%s_against", statMapping)] = bson.M{
					"$sum": bson.M{
						"$divide": bson.A{
							fmt.Sprintf("$opponent.stats.%s.%s", groupName, statMapping),
							"$game.match.event.mode",
						},
					},
				}
			} else if statType == "differential" {
				_group[fmt.Sprintf("%s_differential", statMapping)] = bson.M{
					"$sum": bson.M{
						"$divide": bson.A{
							bson.M{
								"$subtract": bson.A{
									fmt.Sprintf("$team.stats.%s.%s", groupName, statMapping),
									fmt.Sprintf("$opponent.stats.%s.%s", groupName, statMapping),
								},
							},
							"$game.match.event.mode",
						},
					},
				}

			} else if groupName == "ball" {
				_group[statMapping] = bson.M{
					"$sum": bson.M{
						"$divide": bson.A{
							fmt.Sprintf("$team.stats.%s.%s", groupName, statMapping),
							"$game.match.event.mode",
						},
					},
				}
			} else {
				_group[statMapping] = bson.M{
					"$sum": fmt.Sprintf("$player.stats.%s.%s", groupName, statMapping),
				}
			}
		}
	}

	pipeline := New(
		Match(filter),
		Group(_group),
		Project(bson.M{
			"_id":  "$_id",
			"date": "$date",
			"games": bson.M{
				"total":   "$games",
				"replays": "$game_replays",
				"wins":    "$game_wins",
			},
			"match": bson.M{
				"total": bson.M{
					"$size": "$matches",
				},
				"replays": bson.M{
					"$size": "$match_replays",
				},
				"wins": bson.M{
					"$size": "$match_wins",
				},
			},
			"stats": teamStatsMapping(_stats),
		}),
	)

	return &Pipeline{
		Pipeline: pipeline,
		Decode: func(cursor *mongo.Cursor) (interface{}, error) {
			var player struct {
				Date  *time.Time `json:"date" bson:"date"`
				Games struct {
					Total   float64 `json:"total" bson:"total"`
					Replays float64 `json:"replays" bson:"replays"`
					Wins    float64 `json:"wins" bson:"wins"`
				} `json:"games" bson:"games"`
				Matches struct {
					Total   float64 `json:"total" bson:"total"`
					Replays float64 `json:"replays" bson:"replays"`
					Wins    float64 `json:"wins" bson:"wins"`
				} `json:"matches" bson:"match"`
				Stats map[string]float64 `json:"stats" bson:"stats"`
			}
			if err := cursor.Decode(&player); err != nil {
				return nil, err
			}
			return player, nil
		},
	}
}
