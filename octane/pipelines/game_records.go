package pipelines

import (
	"fmt"
	"strings"
	"time"

	"github.com/octanegg/zsr/octane"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// GameRecords .
func GameRecords(filter bson.M, stat string) *Pipeline {
	query := bson.M{
		"$add": bson.A{
			fmt.Sprintf("$blue.stats.core.%s", stat[:len(stat)-len("Total")]),
			fmt.Sprintf("$orange.stats.core.%s", stat[:len(stat)-len("Total")]),
		},
	}

	if strings.Contains(stat, "Differential") {
		query = bson.M{
			"$cond": bson.A{
				bson.M{
					"$gt": bson.A{
						fmt.Sprintf("$blue.stats.core.%s", stat[:len(stat)-len("Differential")]),
						fmt.Sprintf("$orange.stats.core.%s", stat[:len(stat)-len("Differential")]),
					},
				},
				bson.M{
					"$subtract": bson.A{
						fmt.Sprintf("$blue.stats.core.%s", stat[:len(stat)-len("Differential")]),
						fmt.Sprintf("$orange.stats.core.%s", stat[:len(stat)-len("Differential")]),
					},
				},
				bson.M{
					"$subtract": bson.A{
						fmt.Sprintf("$orange.stats.core.%s", stat[:len(stat)-len("Differential")]),
						fmt.Sprintf("$blue.stats.core.%s", stat[:len(stat)-len("Differential")]),
					},
				},
			},
		}
	}

	pipeline := New(
		Match(filter),
		Project(bson.M{
			"_id":      "$_id",
			"match":    "$match",
			"map":      "$map",
			"duration": "$duration",
			"date":     "$date",
			"blue":     "$blue",
			"orange":   "$orange",
			"stat":     query,
		}),
		Sort("stat", true),
		Limit(25),
	)

	return &Pipeline{
		Pipeline: pipeline,
		Decode: func(cursor *mongo.Cursor) (interface{}, error) {
			var team struct {
				ID       *primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
				Match    *octane.Match       `json:"match,omitempty" bson:"match,omitempty"`
				Map      string              `json:"map,omitempty" bson:"map,omitempty"`
				Duration int                 `json:"duration,omitempty" bson:"duration,omitempty"`
				Date     *time.Time          `json:"date,omitempty" bson:"date,omitempty"`
				Blue     *octane.GameSide    `json:"blue,omitempty" bson:"blue,omitempty"`
				Orange   *octane.GameSide    `json:"orange,omitempty" bson:"orange,omitempty"`
				Stat     int                 `json:"stat,omitempty" bson:"stat,omitempty"`
			}
			if err := cursor.Decode(&team); err != nil {
				return nil, err
			}
			return team, nil
		},
	}
}
