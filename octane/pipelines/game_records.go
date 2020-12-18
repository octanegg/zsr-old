package pipelines

import (
	"fmt"
	"time"

	"github.com/octanegg/zsr/octane"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// GamePlayerRecords .
func GamePlayerRecords(filter bson.M, stat string) *Pipeline {
	pipeline := New(
		Match(filter),
		Project(bson.M{
			"game":     "$game",
			"date":     "$date",
			"team":     "$team",
			"opponent": "$opponent",
			"winner":   "$winner",
			"player":   "$player",
			"stat":     fmt.Sprintf("$stats.core.%s", stat),
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

// GameTeamRecords .
func GameTeamRecords(filter bson.M, stat string) *Pipeline {
	pipeline := New(
		Match(filter),
		Group(bson.M{
			"_id": bson.M{
				"game": "$game._id",
				"team": "$team._id",
			},
			"game": bson.M{
				"$first": "$game",
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
		Project(bson.M{
			"game":     "$game",
			"team":     "$team",
			"opponent": "$opponent",
			"winner":   "$winner",
			"stat":     "$stat",
		}),
		Sort("stat", true),
		Limit(25),
	)

	return &Pipeline{
		Pipeline: pipeline,
		Decode: func(cursor *mongo.Cursor) (interface{}, error) {
			var team struct {
				Team     *octane.Team `json:"team,omitempty" bson:"team,omitempty"`
				Game     *octane.Game `json:"game,omitempty" bson:"game,omitempty"`
				Opponent *octane.Team `json:"opponent,omitempty" bson:"opponent,omitempty"`
				Winner   bool         `json:"winner,omitempty" bson:"winner,omitempty"`
				Stat     int          `json:"stat,omitempty" bson:"stat,omitempty"`
			}
			if err := cursor.Decode(&team); err != nil {
				return nil, err
			}
			return team, nil
		},
	}
}

// GameTotalRecords .
func GameTotalRecords(filter bson.M, stat string) *Pipeline {
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
			"stat": bson.M{
				"$add": bson.A{
					fmt.Sprintf("$blue.stats.core.%s", stat),
					fmt.Sprintf("$orange.stats.core.%s", stat),
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
				Match    *octane.Match    `json:"match,omitempty" bson:"match,omitempty"`
				Map      string           `json:"map,omitempty" bson:"map,omitempty"`
				Duration int              `json:"duration,omitempty" bson:"duration,omitempty"`
				Date     *time.Time       `json:"date,omitempty" bson:"date,omitempty"`
				Blue     *octane.GameSide `json:"blue,omitempty" bson:"blue,omitempty"`
				Orange   *octane.GameSide `json:"orange,omitempty" bson:"orange,omitempty"`
				Stat     int              `json:"stat,omitempty" bson:"stat,omitempty"`
			}
			if err := cursor.Decode(&team); err != nil {
				return nil, err
			}
			return team, nil
		},
	}
}

// GameDifferentialRecords .
func GameDifferentialRecords(filter bson.M, stat string) *Pipeline {
	pipeline := New(
		Match(filter),
		Project(bson.M{
			"_id":           "$_id",
			"match":         "$match",
			"map":           "$map",
			"duration":      "$duration",
			"date":          "$date",
			"blue.winner":   "$blue.winner",
			"orange.winner": "$orange.winner",
			"blue.team":     "$blue.team",
			"orange.team":   "$orange.team",
			"stat": bson.M{
				"$cond": bson.A{
					bson.M{
						"$gt": bson.A{
							fmt.Sprintf("$blue.stats.core.%s", stat),
							fmt.Sprintf("$orange.stats.core.%s", stat),
						},
					},
					bson.M{
						"$subtract": bson.A{
							fmt.Sprintf("$blue.stats.core.%s", stat),
							fmt.Sprintf("$orange.stats.core.%s", stat),
						},
					},
					bson.M{
						"$subtract": bson.A{
							fmt.Sprintf("$orange.stats.core.%s", stat),
							fmt.Sprintf("$blue.stats.core.%s", stat),
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
				Match    *octane.Match    `json:"match,omitempty" bson:"match,omitempty"`
				Map      string           `json:"map,omitempty" bson:"map,omitempty"`
				Duration int              `json:"duration,omitempty" bson:"duration,omitempty"`
				Date     *time.Time       `json:"date,omitempty" bson:"date,omitempty"`
				Blue     *octane.GameSide `json:"blue,omitempty" bson:"blue,omitempty"`
				Orange   *octane.GameSide `json:"orange,omitempty" bson:"orange,omitempty"`
				Stat     int              `json:"stat,omitempty" bson:"stat,omitempty"`
			}
			if err := cursor.Decode(&team); err != nil {
				return nil, err
			}
			return team, nil
		},
	}
}
