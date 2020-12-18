package pipelines

import (
	"fmt"
	"time"

	"github.com/octanegg/zsr/octane"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// SeriesPlayerRecords .
func SeriesPlayerRecords(filter bson.M, stat string) *Pipeline {
	op := "$sum"
	if stat == "rating" {
		op = "$avg"
	}

	pipeline := New(
		Match(filter),
		Group(bson.M{
			"_id": bson.M{
				"match":  "$game.match._id",
				"player": "$player._id",
			},
			"match": bson.M{
				"$first": "$game.match",
			},
			"date": bson.M{
				"$first": "$game.date",
			},
			"team": bson.M{
				"$first": "$team",
			},
			"opponent": bson.M{
				"$first": "$opponent",
			},
			"winner": bson.M{
				"$first": "$winner",
			},
			"player": bson.M{
				"$first": "$player",
			},
			"stat": bson.M{
				op: fmt.Sprintf("$stats.core.%s", stat),
			},
		}),
		Sort("stat", true),
		Limit(25),
	)

	return &Pipeline{
		Pipeline: pipeline,
		Decode: func(cursor *mongo.Cursor) (interface{}, error) {
			var player struct {
				Match    *octane.Match  `json:"match,omitempty" bson:"match,omitempty"`
				Date     *time.Time     `json:"date,omitempty" bson:"date,omitempty"`
				Team     *octane.Team   `json:"team,omitempty" bson:"team,omitempty"`
				Opponent *octane.Team   `json:"opponent,omitempty" bson:"opponent,omitempty"`
				Winner   bool           `json:"winner,omitempty" bson:"winner,omitempty"`
				Player   *octane.Player `json:"player,omitempty" bson:"player,omitempty"`
				Stat     float64        `json:"stat,omitempty" bson:"stat,omitempty"`
			}
			if err := cursor.Decode(&player); err != nil {
				return nil, err
			}
			return player, nil
		},
	}
}

// SeriesTeamRecords .
func SeriesTeamRecords(filter bson.M, stat string) *Pipeline {
	pipeline := New(
		Match(filter),
		Group(bson.M{
			"_id": bson.M{
				"match":  "$game.match._id",
				"player": "$team._id",
			},
			"match": bson.M{
				"$first": "$game.match",
			},
			"date": bson.M{
				"$first": "$game.date",
			},
			"team": bson.M{
				"$first": "$team",
			},
			"opponent": bson.M{
				"$first": "$opponent",
			},
			"winner": bson.M{
				"$first": "$winner",
			},
			"stat": bson.M{
				"$sum": fmt.Sprintf("$stats.core.%s", stat),
			},
		}),
		Sort("stat", true),
		Limit(25),
	)

	return &Pipeline{
		Pipeline: pipeline,
		Decode: func(cursor *mongo.Cursor) (interface{}, error) {
			var team struct {
				Match    *octane.Match `json:"match,omitempty" bson:"match,omitempty"`
				Date     *time.Time    `json:"date,omitempty" bson:"date,omitempty"`
				Team     *octane.Team  `json:"team,omitempty" bson:"team,omitempty"`
				Opponent *octane.Team  `json:"opponent,omitempty" bson:"opponent,omitempty"`
				Winner   bool          `json:"winner,omitempty" bson:"winner,omitempty"`
				Stat     int           `json:"stat,omitempty" bson:"stat,omitempty"`
			}
			if err := cursor.Decode(&team); err != nil {
				return nil, err
			}
			return team, nil
		},
	}
}

// SeriesTotalRecords .
func SeriesTotalRecords(filter bson.M, stat string) *Pipeline {
	pipeline := New(
		Match(filter),
		Group(bson.M{
			"_id": bson.M{
				"match": "$match._id",
			},
			"match": bson.M{
				"$first": "$match",
			},
			"date": bson.M{
				"$first": "$date",
			},
			"blue": bson.M{
				"$last": "$blue",
			},
			"orange": bson.M{
				"$last": "$orange",
			},
			"stat": bson.M{
				"$sum": bson.M{
					"$add": bson.A{
						fmt.Sprintf("$blue.stats.core.%s", stat),
						fmt.Sprintf("$orange.stats.core.%s", stat),
					},
				},
			},
		}),
		Sort("stat", true),
		Limit(25),
	)

	return &Pipeline{
		Pipeline: pipeline,
		Decode: func(cursor *mongo.Cursor) (interface{}, error) {
			var team struct {
				Match  *octane.Match    `json:"match,omitempty" bson:"match,omitempty"`
				Date   *time.Time       `json:"date,omitempty" bson:"date,omitempty"`
				Blue   *octane.GameSide `json:"blue,omitempty" bson:"blue,omitempty"`
				Orange *octane.GameSide `json:"orange,omitempty" bson:"orange,omitempty"`
				Stat   int              `json:"stat,omitempty" bson:"stat,omitempty"`
			}
			if err := cursor.Decode(&team); err != nil {
				return nil, err
			}
			return team, nil
		},
	}
}

// SeriesDifferentialRecords .
func SeriesDifferentialRecords(filter bson.M, stat string) *Pipeline {
	pipeline := New(
		Match(filter),
		Group(bson.M{
			"_id": bson.M{
				"match": "$match._id",
			},
			"match": bson.M{
				"$first": "$match",
			},
			"date": bson.M{
				"$first": "$date",
			},
			"blue": bson.M{
				"$last": "$blue",
			},
			"orange": bson.M{
				"$last": "$orange",
			},
			"stat": bson.M{
				"$sum": bson.M{
					"$subtract": bson.A{
						fmt.Sprintf("$blue.stats.core.%s", stat),
						fmt.Sprintf("$orange.stats.core.%s", stat),
					},
				},
			},
		}),
		Project(bson.M{
			"_id":           "$_id",
			"match":         "$match",
			"date":          "$date",
			"blue.winner":   "$blue.winner",
			"orange.winner": "$orange.winner",
			"blue.team":     "$blue.team",
			"orange.team":   "$orange.team",
			"stat": bson.M{
				"$cond": bson.A{
					bson.M{
						"$gt": bson.A{"$stat", 0},
					},
					"$stat",
					bson.M{
						"$multiply": bson.A{
							"$stat",
							-1,
						},
					},
				},
			},
		}),
		Sort("stat", true),
		Limit(25),
	)

	return &Pipeline{
		Pipeline: pipeline,
		Decode: func(cursor *mongo.Cursor) (interface{}, error) {
			var team struct {
				Match  *octane.Match    `json:"match,omitempty" bson:"match,omitempty"`
				Date   *time.Time       `json:"date,omitempty" bson:"date,omitempty"`
				Blue   *octane.GameSide `json:"blue,omitempty" bson:"blue,omitempty"`
				Orange *octane.GameSide `json:"orange,omitempty" bson:"orange,omitempty"`
				Stat   int              `json:"stat,omitempty" bson:"stat,omitempty"`
			}
			if err := cursor.Decode(&team); err != nil {
				return nil, err
			}
			return team, nil
		},
	}
}
