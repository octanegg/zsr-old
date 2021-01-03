package pipelines

import (
	"fmt"
	"time"

	"github.com/octanegg/zsr/octane"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// PlayerGameRecords .
func PlayerGameRecords(filter bson.M, stat string) *Pipeline {
	pipeline := New(
		Match(filter),
		Project(bson.M{
			"game":     "$game",
			"date":     "$date",
			"team":     "$team",
			"opponent": "$opponent",
			"winner":   "$winner",
			"player":   "$player",
			"stat":     fmt.Sprintf("$stats.player.core.%s", stat),
		}),
		Sort("stat", true),
		Limit(25),
	)

	return &Pipeline{
		Pipeline: pipeline,
		Decode: func(cursor *mongo.Cursor) (interface{}, error) {
			var player struct {
				Game     *octane.Game   `json:"game,omitempty" bson:"game,omitempty"`
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

// PlayerSeriesRecords .
func PlayerSeriesRecords(filter bson.M, stat string) *Pipeline {
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
				op: fmt.Sprintf("$stats.player.core.%s", stat),
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
