package pipelines

import (
	"fmt"
	"time"

	"github.com/octanegg/zsr/octane"
	"github.com/octanegg/zsr/octane/stats"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// PlayerStats .
func PlayerStats(filter, group, having bson.M, _stats []string) *Pipeline {
	_group := bson.M{
		"_id": group,
		"player": bson.M{
			"$first": "$player.player",
		},
		"events": bson.M{
			"$addToSet": "$game.match.event",
		},
		"teams": bson.M{
			"$addToSet": "$team.team",
		},
		"opponents": bson.M{
			"$addToSet": "$opponent.team",
		},
		"start_date": bson.M{
			"$min": "$game.date",
		},
		"end_date": bson.M{
			"$max": "$game.date",
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
		Match(having),
		Project(bson.M{
			"_id":        "$_id",
			"player":     "$player",
			"events":     "$events",
			"teams":      "$teams",
			"opponents":  "$opponents",
			"start_date": "$start_date",
			"end_date":   "$end_date",
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
				Player    *octane.Player  `json:"player" bson:"player"`
				Events    []*octane.Event `json:"events" bson:"events"`
				Teams     []*octane.Team  `json:"teams" bson:"teams"`
				Opponents []*octane.Team  `json:"opponents" bson:"opponents"`
				StartDate *time.Time      `json:"startDate" bson:"start_date"`
				EndDate   *time.Time      `json:"endDate" bson:"end_date"`
				Games     struct {
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

func playerStatsMapping(s []string) bson.M {
	mapping := bson.M{}
	for _, stat := range s {
		if stat == "shootingPercentage" {
			mapping[stat] = bson.M{
				"$cond": bson.A{
					bson.M{
						"$gt": bson.A{"$shots", 0},
					},
					bson.M{
						"$multiply": bson.A{
							bson.M{
								"$divide": bson.A{"$goals", "$shots"},
							},
							100,
						},
					},
					0,
				},
			}
		} else if stat == "goalParticipation" {
			mapping[stat] = bson.M{
				"$cond": bson.A{
					bson.M{
						"$gt": bson.A{"$team_goals", 0},
					},
					bson.M{
						"$multiply": bson.A{
							bson.M{
								"$divide": bson.A{
									bson.M{"$add": bson.A{"$goals", "$assists"}},
									"$team_goals",
								},
							},
							100,
						},
					},
					0,
				},
			}
		} else {
			_, statMapping := stats.PlayerStatMapping(stat)
			if statMapping == "" {
				continue
			}

			mapping[stat] = fmt.Sprintf("$%s", statMapping)
		}
	}

	return mapping

}
