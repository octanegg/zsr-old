package pipelines

import (
	"fmt"
	"strings"
	"time"

	"github.com/octanegg/zsr/octane"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// SeriesRecords .
func SeriesRecords(filter bson.M, stat string) *Pipeline {
	var query bson.M
	if strings.Contains(stat, "Total") {
		query = bson.M{
			"$sum": bson.M{
				"$add": bson.A{
					fmt.Sprintf("$blue.stats.core.%s", stat[:len(stat)-len("Total")]),
					fmt.Sprintf("$orange.stats.core.%s", stat[:len(stat)-len("Total")]),
				},
			},
		}

	} else if strings.Contains(stat, "Differential") {
		query = bson.M{
			"$sum": bson.M{
				"$subtract": bson.A{
					fmt.Sprintf("$blue.stats.core.%s", stat[:len(stat)-len("Differential")]),
					fmt.Sprintf("$orange.stats.core.%s", stat[:len(stat)-len("Differential")]),
				},
			},
		}
	} else {
		query = bson.M{
			"$sum": "$duration",
		}
	}

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
			"stat": query,
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
