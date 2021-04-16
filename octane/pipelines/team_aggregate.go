package pipelines

import (
	"fmt"
	"strings"
	"time"

	"github.com/octanegg/zsr/octane"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// TeamAggregate .
func TeamAggregate(filter bson.M, group interface{}, having bson.M, cluster string) *Pipeline {
	pipeline := New(
		Match(filter),
		Sort("game.date", false),
		Group(teamAggregateGroup(group)),
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
			"replays": bson.M{
				"$toInt": "$game_replays",
			},
			"wins": bson.M{
				"$toInt": "$wins",
			},
			"win_percentage": bson.M{
				"$multiply": bson.A{
					bson.M{
						"$divide": bson.A{
							"$wins", "$games",
						},
					},
					100,
				},
			},
			"stats": teamAggregateStats(cluster),
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
				Replays       int              `json:"replays" bson:"replays"`
				Wins          int              `json:"wins" bson:"wins"`
				WinPercentage float64          `json:"winPercentage" bson:"win_percentage"`
				Stats         struct {
					Core         *octane.TeamCore        `json:"core" bson:"core"`
					Against      *octane.TeamCore        `json:"against" bson:"against"`
					Differential *octane.TeamCore        `json:"differential" bson:"differential"`
					Boost        *octane.TeamBoost       `json:"boost" bson:"boost"`
					Movement     *octane.TeamMovement    `json:"movement" bson:"movement"`
					Positioning  *octane.TeamPositioning `json:"positioning" bson:"positioning"`
					Demolitions  *octane.TeamDemolitions `json:"demo" bson:"demo"`
					Ball         *octane.TeamBall        `json:"ball" bson:"ball"`
				} `json:"stats" bson:"stats"`
			}
			if err := cursor.Decode(&team); err != nil {
				return nil, err
			}
			return team, nil
		},
	}
}

func teamAggregateGroup(group interface{}) bson.M {
	m := bson.M{
		"_id": group,
		"team": bson.M{
			"$first": "$team.team",
		},
		"events": bson.M{
			"$addToSet": "$game.match.event",
		},
		"matches": bson.M{
			"$addToSet": "$game.match._id",
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
			"$addToSet": "$player.player",
		},
		"games": bson.M{
			"$sum": bson.M{
				"$divide": bson.A{1, "$game.match.event.mode"},
			},
		},
		"mode": bson.M{
			"$first": "$game.match.event.mode",
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
		"match_replays": bson.M{
			"$addToSet": bson.M{
				"$cond": bson.A{
					bson.M{
						"$ifNull": bson.A{"$game.ballchasing", false},
					}, "$game.match._id", bson.TypeNull,
				},
			},
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
	}

	for groupName, group := range TeamStats {
		for k, v := range group {
			m[k] = bson.M{
				"$sum": bson.M{
					"$divide": bson.A{TeamStatToField(groupName, v), "$game.match.event.mode"},
				},
			}
		}
	}

	for k, v := range TeamCore {
		m[fmt.Sprintf("%sOpponent", k)] = bson.M{
			"$sum": bson.M{
				"$divide": bson.A{fmt.Sprintf("$opponent.stats.core.%s", v), "$game.match.event.mode"},
			},
		}
	}

	return m
}

func teamAggregateStats(cluster string) bson.M {
	var getStatsForGroup = func(groupName string, group map[string]string) bson.M {
		m := bson.M{}
		for k, v := range group {
			var isAverage bool

			for _, field := range FieldsToAverageOverReplays {
				if strings.Contains(k, field) {
					m[v] = bson.M{
						"$cond": bson.A{
							bson.M{"$gt": bson.A{"$game_replays", 0}},
							bson.M{
								"$divide": bson.A{fmt.Sprintf("$%s", k), "$game_replays"},
							}, 0,
						},
					}
					isAverage = true
				}
			}

			for _, field := range FieldsToAverage {
				if strings.Contains(k, field) {
					m[v] = bson.M{
						"$cond": bson.A{
							bson.M{"$gt": bson.A{"$games", 0}},
							bson.M{
								"$divide": bson.A{fmt.Sprintf("$%s", k), "$games"},
							}, 0,
						},
					}
					isAverage = true
				}
			}

			if k == "shootingPercentage" {
				m[v] = bson.M{
					"$cond": bson.A{
						bson.M{"$eq": bson.A{"$shots", 0}},
						1,
						bson.M{
							"$multiply": bson.A{
								bson.M{
									"$divide": bson.A{"$goals", "$shots"},
								},
								100,
							},
						},
					},
				}
			}

			if !isAverage {
				if cluster == "total" {
					m[v] = bson.M{
						"$divide": bson.A{fmt.Sprintf("$%s", k), 1},
					}
				} else if cluster == "series" {
					if groupName != "core" && groupName != "advanced" {
						m[v] = bson.M{
							"$cond": bson.A{
								bson.M{"$gt": bson.A{bson.M{"$size": "$match_replays"}, 0}},
								bson.M{
									"$divide": bson.A{fmt.Sprintf("$%s", k), bson.M{"$size": "$match_replays"}},
								}, 0,
							},
						}
					} else {
						m[v] = bson.M{
							"$divide": bson.A{fmt.Sprintf("$%s", k), bson.M{"$size": "$matches"}},
						}
					}
				} else {
					if groupName != "core" && groupName != "advanced" {
						m[v] = bson.M{
							"$cond": bson.A{
								bson.M{"$gt": bson.A{"$game_replays", 0}},
								bson.M{
									"$divide": bson.A{fmt.Sprintf("$%s", k), "$game_replays"},
								}, 0,
							},
						}
					} else {
						m[v] = bson.M{
							"$divide": bson.A{fmt.Sprintf("$%s", k), "$games"},
						}
					}
				}
			}
		}
		return m
	}

	stats := bson.M{}
	for name, group := range TeamStats {
		stats[name] = getStatsForGroup(name, group)
	}

	against, differential := bson.M{}, bson.M{}
	for k, v := range TeamCore {
		if k == "shootingPercentage" {
			against[v] = bson.M{
				"$cond": bson.A{
					bson.M{"$eq": bson.A{"$shotsOpponent", 0}},
					1,
					bson.M{
						"$multiply": bson.A{
							bson.M{
								"$divide": bson.A{"$goalsOpponent", "$shotsOpponent"},
							},
							100,
						},
					},
				},
			}
		} else if cluster == "total" {
			differential[v] = bson.M{
				"$subtract": bson.A{fmt.Sprintf("$%s", k), fmt.Sprintf("$%sOpponent", k)},
			}
			against[v] = fmt.Sprintf("$%sOpponent", k)
		} else if cluster == "series" {
			differential[v] = bson.M{
				"$divide": bson.A{
					bson.M{
						"$subtract": bson.A{fmt.Sprintf("$%s", k), fmt.Sprintf("$%sOpponent", k)},
					},
					bson.M{
						"$size": "$matches",
					},
				},
			}
			against[v] = bson.M{
				"$divide": bson.A{fmt.Sprintf("$%sOpponent", k), bson.M{
					"$size": "$matches",
				}},
			}
		} else {
			differential[v] = bson.M{
				"$divide": bson.A{
					bson.M{
						"$subtract": bson.A{fmt.Sprintf("$%s", k), fmt.Sprintf("$%sOpponent", k)},
					},
					"$games",
				},
			}
			against[v] = bson.M{
				"$divide": bson.A{fmt.Sprintf("$%sOpponent", k), "$games"},
			}
		}
	}

	stats["against"] = against
	stats["differential"] = differential

	return stats
}
