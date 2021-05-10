package pipelines

import (
	"fmt"
	"time"

	"github.com/octanegg/zsr/octane"
	"github.com/octanegg/zsr/octane/stats"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// PlayerStatsX .
func PlayerStatsX(filter, group, having bson.M, _stats []string) *Pipeline {
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

func playerStatsMapping(s []string) bson.M {
	mapping := bson.M{}
	for _, stat := range s {
		_, statMapping := stats.PlayerStatMapping(stat)
		if statMapping == "" {
			continue
		}

		mapping[stat] = fmt.Sprintf("$%s", statMapping)
	}

	return mapping

}
