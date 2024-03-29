package pipelines

import (
	"fmt"
	"time"

	"github.com/octanegg/zsr/octane"
	"github.com/octanegg/zsr/octane/stats"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// PlayerGameRecords .
func PlayerGameRecords(filter bson.M, stat string) *Pipeline {
	project := bson.M{
		"game":     "$game",
		"date":     "$date",
		"team":     "$team.team",
		"opponent": "$opponent.team",
		"winner":   "$team.winner",
		"player":   "$player.player",
	}

	groupName, statMapping := stats.PlayerStatMapping(stat)
	_stat := fmt.Sprintf("player.stats.%s.%s", groupName, statMapping)
	if groupName == "advanced" {
		_stat = fmt.Sprintf("player.%s.%s", groupName, statMapping)
	}

	project["stat"] = fmt.Sprintf("$%s", _stat)

	pipeline := New(
		Match(filter),
		Sort(_stat, true),
		Project(project),
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
				Stat     float64        `json:"stat" bson:"stat"`
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
	group := bson.M{
		"_id": bson.M{
			"match":  "$game.match._id",
			"player": "$player.player._id",
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
		"player": bson.M{
			"$first": "$player.player",
		},
	}

	op := "$sum"
	if stat == "rating" {
		op = "$avg"
	}

	groupName, statMapping := stats.PlayerStatMapping(stat)
	group["stat"] = bson.M{
		op: fmt.Sprintf("$player.stats.%s.%s", groupName, statMapping),
	}
	if groupName == "advanced" {
		group["stat"] = bson.M{
			op: fmt.Sprintf("$player.%s.%s", groupName, statMapping),
		}
	}

	pipeline := New(
		Match(filter),
		Group(group),
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
				Stat     float64        `json:"stat" bson:"stat"`
			}
			if err := cursor.Decode(&player); err != nil {
				return nil, err
			}
			return player, nil
		},
	}
}
