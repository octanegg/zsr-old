package pipelines

import (
	"fmt"
	"time"

	"github.com/octanegg/zsr/octane"
	"github.com/octanegg/zsr/octane/stats"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// TeamStatsX .
func TeamStatsX(filter, group, having bson.M, _stats []string) *Pipeline {
	_group := bson.M{
		"_id": group,
		"team": bson.M{
			"$first": "$team.team",
		},
		"events": bson.M{
			"$addToSet": "$game.match.event",
		},
		"players": bson.M{
			"$addToSet": "$player.player",
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
		groupName, statMapping := stats.TeamStatMapping(stat)
		if statMapping == "" {
			continue
		}

		field := bson.M{
			"$sum": fmt.Sprintf("$player.stats.%s.%s", groupName, statMapping),
		}

		if groupName == "against" {
			field = bson.M{
				"$sum": bson.M{
					"$divide": bson.A{
						fmt.Sprintf("$opponent.stats.core.%s", statMapping),
						"$game.match.event.mode",
					},
				},
			}
		} else if groupName == "differential" {
			field = bson.M{
				"$sum": bson.M{
					"$divide": bson.A{
						bson.M{
							"$subtract": bson.A{
								fmt.Sprintf("$team.stats.core.%s", statMapping),
								fmt.Sprintf("$opponent.stats.core.%s", statMapping),
							},
						},
						"$game.match.event.mode",
					},
				},
			}

		} else if groupName == "ball" || stat == "shootingPercentage" {
			field = bson.M{
				"$sum": bson.M{
					"$divide": bson.A{
						fmt.Sprintf("$team.stats.%s.%s", groupName, statMapping),
						"$game.match.event.mode",
					},
				},
			}
		}

		_group[statMapping] = field
	}

	pipeline := New(
		Match(filter),
		Group(_group),
		Match(having),
		Project(bson.M{
			"_id":        "$_id",
			"team":       "$team",
			"events":     "$events",
			"players":    "$players",
			"opponents":  "$opponents",
			"start_date": "$start_date",
			"end_date":   "$end_date",
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
				Team      *octane.Team     `json:"team" bson:"team"`
				Events    []*octane.Event  `json:"events" bson:"events"`
				Players   []*octane.Player `json:"players" bson:"players"`
				Opponents []*octane.Team   `json:"opponents" bson:"opponents"`
				StartDate *time.Time       `json:"startDate" bson:"start_date"`
				EndDate   *time.Time       `json:"endDate" bson:"end_date"`
				Games     struct {
					Total   int `json:"total" bson:"total"`
					Replays int `json:"replays" bson:"replays"`
					Wins    int `json:"wins" bson:"wins"`
				} `json:"games" bson:"games"`
				Matches struct {
					Total   int `json:"total" bson:"total"`
					Replays int `json:"replays" bson:"replays"`
					Wins    int `json:"wins" bson:"wins"`
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

func teamStatsMapping(s []string) bson.M {
	mapping := bson.M{}
	for _, stat := range s {
		_, statMapping := stats.TeamStatMapping(stat)
		if statMapping != "" {
			mapping[stat] = fmt.Sprintf("$%s", statMapping)
		} else {
			mapping[stat] = fmt.Sprintf("$%s", stat)
		}
	}

	return mapping

}
