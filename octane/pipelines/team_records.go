package pipelines

import (
	"fmt"
	"time"

	"github.com/octanegg/zsr/octane"
	"github.com/octanegg/zsr/octane/stats"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// TeamGameRecords .
func TeamGameRecords(filter bson.M, stat string) *Pipeline {
	group := bson.M{
		"_id": bson.M{
			"game": "$game._id",
			"team": "$team.team._id",
		},
		"game": bson.M{
			"$first": "$game",
		},
		"team": bson.M{
			"$first": "$team.team",
		},
		"opponent": bson.M{
			"$first": "$opponent.team",
		},
		"winner": bson.M{
			"$first": "$team.winner",
		},
	}

	groupName, statMapping := stats.TeamStatMapping(stat)
	group["stat"] = bson.M{
		"$sum": fmt.Sprintf("$player.stats.%s.%s", groupName, statMapping),
	}
	if groupName == "advanced" {
		group["stat"] = bson.M{
			"$sum": fmt.Sprintf("$player.%s.%s", groupName, statMapping),
		}
	}

	pipeline := New(
		Match(filter),
		Group(group),
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
				Stat     float64      `json:"stat" bson:"stat"`
			}
			if err := cursor.Decode(&team); err != nil {
				return nil, err
			}
			return team, nil
		},
	}
}

// TeamSeriesRecords .
func TeamSeriesRecords(filter bson.M, stat string) *Pipeline {
	group := bson.M{
		"_id": bson.M{
			"match":  "$game.match._id",
			"player": "$team.team._id",
		},
		"match": bson.M{
			"$first": "$game.match",
		},
		"date": bson.M{
			"$first": "$game.date",
		},
		"team": bson.M{
			"$first": "$team.team",
		},
		"opponent": bson.M{
			"$first": "$opponent.team",
		},
		"winner": bson.M{
			"$first": "$team.winner",
		},
	}

	groupName, statMapping := stats.TeamStatMapping(stat)
	group["stat"] = bson.M{
		"$sum": fmt.Sprintf("$player.stats.%s.%s", groupName, statMapping),
	}
	if groupName == "advanced" {
		group["stat"] = bson.M{
			"$sum": fmt.Sprintf("$player.%s.%s", groupName, statMapping),
		}
	}

	pipeline := New(
		Match(filter),
		Group(group),
		Project(bson.M{
			"match":    "$match",
			"date":     "$date",
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
				Match    *octane.Match `json:"match,omitempty" bson:"match,omitempty"`
				Date     *time.Time    `json:"date,omitempty" bson:"date,omitempty"`
				Team     *octane.Team  `json:"team,omitempty" bson:"team,omitempty"`
				Opponent *octane.Team  `json:"opponent,omitempty" bson:"opponent,omitempty"`
				Winner   bool          `json:"winner,omitempty" bson:"winner,omitempty"`
				Stat     float64       `json:"stat" bson:"stat"`
			}
			if err := cursor.Decode(&team); err != nil {
				return nil, err
			}
			return team, nil
		},
	}
}
