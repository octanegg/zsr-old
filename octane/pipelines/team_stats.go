package pipelines

import (
	"fmt"
	"strings"
	"time"

	"github.com/octanegg/zsr/octane"
	"github.com/octanegg/zsr/octane/stats"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// TeamStats .
func TeamStats(filter, group, having bson.M, _stats []string) *Pipeline {
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
		"game_seconds": bson.M{
			"$sum": bson.M{
				"$divide": bson.A{"$game.duration", "$game.match.event.mode"},
			},
		},
		"game_replay_seconds": bson.M{
			"$sum": bson.M{
				"$cond": bson.A{
					bson.M{
						"$ifNull": bson.A{"$game.ballchasing", false},
					}, bson.M{
						"$divide": bson.A{"$game.duration", "$game.match.event.mode"},
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
					}, "$game.match._id", "$$REMOVE",
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
					}, "$game.match._id", "$$REMOVE",
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
		} else if stat == "shootingPercentageAgainst" {
			_group["opponent_goals"] = bson.M{
				"$sum": bson.M{
					"$divide": bson.A{
						"$opponent.stats.core.goals",
						"$game.match.event.mode",
					},
				},
			}
			_group["opponent_shots"] = bson.M{
				"$sum": bson.M{
					"$divide": bson.A{
						"$opponent.stats.core.shots",
						"$game.match.event.mode",
					},
				},
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
				"total":          "$games",
				"replays":        "$game_replays",
				"wins":           "$game_wins",
				"seconds":        "$game_seconds",
				"replay_seconds": "$game_replay_seconds",
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
					Total         float64 `json:"total" bson:"total"`
					Replays       float64 `json:"replays" bson:"replays"`
					Wins          float64 `json:"wins" bson:"wins"`
					Seconds       float64 `json:"seconds" bson:"seconds"`
					ReplaySeconds float64 `json:"replaySeconds" bson:"replay_seconds"`
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

func teamStatsMapping(s []string) bson.M {
	mapping := bson.M{}
	for _, stat := range s {
		if stat == "shootingPercentage" {
			mapping[stat] = bson.M{
				"$multiply": bson.A{
					bson.M{
						"$divide": bson.A{"$goals", "$shots"},
					},
					100,
				},
			}
		} else if stat == "shootingPercentageAgainst" {
			mapping[stat] = bson.M{
				"$multiply": bson.A{
					bson.M{
						"$divide": bson.A{"$opponent_goals", "$opponent_shots"},
					},
					100,
				},
			}
		} else {
			_stat := stat

			var statType string
			if strings.Contains(stat, "Against") {
				_stat = stat[:len(stat)-7]
				statType = "against"
			} else if strings.Contains(stat, "Differential") {
				_stat = stat[:len(stat)-12]
				statType = "differential"
			}

			_, statMapping := stats.TeamStatMapping(_stat)
			if statMapping == "" {
				mapping[stat] = fmt.Sprintf("$%s", stat)
			} else if statType == "against" {
				mapping[stat] = fmt.Sprintf("$%s_against", statMapping)
			} else if statType == "differential" {
				mapping[stat] = fmt.Sprintf("$%s_differential", statMapping)
			} else {
				mapping[stat] = fmt.Sprintf("$%s", statMapping)
			}
		}
	}

	return mapping

}
