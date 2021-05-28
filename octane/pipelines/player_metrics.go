package pipelines

import (
	"fmt"
	"time"

	"github.com/octanegg/zsr/octane/stats"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// PlayerMetrics .
func PlayerMetrics(filter bson.M, _stats []string) *Pipeline {
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
			"$sum": 1,
		},
		"game_replays": bson.M{
			"$sum": bson.M{
				"$cond": bson.A{
					bson.M{
						"$ifNull": bson.A{"$game.ballchasing", false},
					}, 1, 0,
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
					}, 1, 0,
				},
			},
		},
		"game_seconds": bson.M{
			"$sum": "$game.duration",
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
				"$sum": "$player.stats.core.goals",
			}
			_group["shots"] = bson.M{
				"$sum": "$player.stats.core.shots",
			}
		} else if stat == "goalParticipation" {
			_group["goals"] = bson.M{
				"$sum": "$player.stats.core.goals",
			}
			_group["assists"] = bson.M{
				"$sum": "$player.stats.core.assists",
			}
			_group["team_goals"] = bson.M{
				"$sum": "$team.stats.core.goals",
			}
		} else {
			groupName, statMapping := stats.PlayerStatMapping(stat)
			if statMapping == "" {
				continue
			}

			field := fmt.Sprintf("$player.stats.%s.%s", groupName, statMapping)
			if groupName == "advanced" {
				field = fmt.Sprintf("$player.%s.%s", groupName, statMapping)
			}

			_group[statMapping] = bson.M{
				"$sum": field,
			}
		}
	}

	pipeline := New(
		Match(filter),
		Group(_group),
		Project(bson.M{
			"_id":    "$_id",
			"player": "$player",
			"date":   "$date",
			"games": bson.M{
				"total":   "$games",
				"replays": "$game_replays",
				"wins":    "$game_wins",
				"seconds": "$game_seconds",
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
			"stats": playerStatsMapping(_stats),
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
					Seconds float64 `json:"seconds" bson:"seconds"`
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
