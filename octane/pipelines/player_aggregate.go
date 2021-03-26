package pipelines

import (
	"fmt"
	"strings"
	"time"

	"github.com/octanegg/zsr/ballchasing"
	"github.com/octanegg/zsr/octane"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// PlayerAggregate .
func PlayerAggregate(filter bson.M, group interface{}, having bson.M, cluster string) *Pipeline {
	pipeline := New(
		Match(filter),
		Sort("game.date", false),
		Group(playerAggregateGroup(group)),
		Match(having),
		Project(bson.M{
			"_id":        "$_id",
			"player":     "$player",
			"team":       "$team",
			"opponents":  "$opponents",
			"events":     "$events",
			"start_date": "$start_date",
			"end_date":   "$end_date",
			"games":      "$games",
			"replays":    "$game_replays",
			"wins":       "$wins",
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
			"stats": playerAggregateStats(cluster),
		}),
		Sort("end_date", true),
	)

	return &Pipeline{
		Pipeline: pipeline,
		Decode: func(cursor *mongo.Cursor) (interface{}, error) {
			var player struct {
				Player        *octane.Player  `json:"player" bson:"player"`
				Events        []*octane.Event `json:"events" bson:"events"`
				Team          *octane.Team    `json:"team" bson:"team"`
				Opponents     []*octane.Team  `json:"opponents" bson:"opponents"`
				StartDate     *time.Time      `json:"startDate" bson:"start_date"`
				EndDate       *time.Time      `json:"endDate" bson:"end_date"`
				Games         int             `json:"games" bson:"games"`
				Replays       int             `json:"replays" bson:"replays"`
				Wins          int             `json:"wins" bson:"wins"`
				WinPercentage float64         `json:"winPercentage" bson:"win_percentage"`
				Stats         struct {
					Core        *ballchasing.PlayerCore        `json:"core" bson:"core"`
					Boost       *ballchasing.PlayerBoost       `json:"boost" bson:"boost"`
					Movement    *ballchasing.PlayerMovement    `json:"movement" bson:"movement"`
					Positioning *ballchasing.PlayerPositioning `json:"positioning" bson:"positioning"`
					Demolitions *ballchasing.PlayerDemolitions `json:"demo" bson:"demo"`
					Advanced    *octane.AdvancedStats          `json:"advanced" bson:"advanced"`
				} `json:"stats" bson:"stats"`
			}
			if err := cursor.Decode(&player); err != nil {
				return nil, err
			}
			return player, nil
		},
	}
}

func playerAggregateGroup(group interface{}) bson.M {
	m := bson.M{
		"_id": group,
		"player": bson.M{
			"$first": "$player.player",
		},
		"events": bson.M{
			"$addToSet": "$game.match.event",
		},
		"matches": bson.M{
			"$addToSet": "$game.match._id",
		},
		"start_date": bson.M{
			"$first": "$game.date",
		},
		"end_date": bson.M{
			"$last": "$game.date",
		},
		"team": bson.M{
			"$first": "$team.team",
		},
		"opponents": bson.M{
			"$addToSet": "$opponent.team",
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
					}, 1, 0,
				},
			},
		},
	}

	for groupName, group := range PlayerStats {
		for k, v := range group {
			m[k] = bson.M{
				"$sum": PlayerStatToField(groupName, v),
			}
		}
	}

	return m
}

func playerAggregateStats(cluster string) bson.M {
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

			for _, field := range FieldsToAverageAsInt {
				if strings.Contains(k, field) {
					m[v] = bson.M{
						"$toInt": bson.M{
							"$cond": bson.A{
								bson.M{"$gt": bson.A{"$game_replays", 0}},
								bson.M{
									"$divide": bson.A{fmt.Sprintf("$%s", k), "$game_replays"},
								}, 0,
							},
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

			if !isAverage {
				if cluster == "game" {
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
					m[v] = bson.M{
						"$divide": bson.A{fmt.Sprintf("$%s", k), 1},
					}
				}
			}
		}
		return m
	}

	stats := bson.M{}
	for name, group := range PlayerStats {
		stats[name] = getStatsForGroup(name, group)
	}

	return stats
}
