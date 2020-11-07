package octane

import (
	"fmt"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	validStats = []string{"score", "goals", "assists", "saves", "shots", "rating"}
)

type records struct {
	Collection Collection
}

// Records .
type Records interface {
	Player(string, bson.M) ([]interface{}, error)
}

// Record .
type Record struct {
	ID     *primitive.ObjectID `json:"_id" bson:"_id"`
	Match  *Match              `json:"match,omitempty" bson:"match,omitempty"`
	Date   *time.Time          `json:"date,omitempty" bson:"date,omitempty"`
	Blue   *GameSide           `json:"blue,omitempty" bson:"blue,omitempty"`
	Orange *GameSide           `json:"orange,omitempty" bson:"orange,omitempty"`
	Player *PlayerStats        `json:"player,omitempty" bson:"player,omitempty"`
	Index  int 				   `json:"index,omitempty" bson:"index,omitempty"`
}

var (
	getStatLines = []bson.M{
		{
			"$project": bson.M{
				"blue.team":     true,
				"blue.score":    true,
				"blue.winner":   true,
				"orange.team":   true,
				"orange.score":  true,
				"orange.winner": true,
				"match":         true,
				"date":          true,
				"player": bson.M{
					"$concatArrays": bson.A{"$blue.players", "$orange.players"},
				},
			},
		},
		{
			"$unwind": bson.M{
				"path":                       "$player",
				"includeArrayIndex":          "index",
				"preserveNullAndEmptyArrays": true,
			},
		},
	}
)

func (r *records) Player(stat string, filter bson.M) ([]interface{}, error) {
	return r.Collection.Aggregate(playerPipeline(filter, stat, false), CursorToRecord)
}

func playerPipeline(filter bson.M, sortField string, ascending bool) []bson.M {
	pipeline := getStatLines
	if filter != nil {
		pipeline = append(pipeline, bson.M{"$match": filter})
	}

	order := -1
	if ascending {
		order = 1
	}

	pipeline = append(pipeline, bson.M{
		"$sort": bson.M{
			fmt.Sprintf("player.stats.core.%s", sortField): order,
		},
	})

	return pipeline
}

// IsValidStat .
func IsValidStat(stat string) bool {
	for _, validStat := range validStats {
		if stat == validStat {
			return true
		}
	}
	return false
}

// ValidStats .
func ValidStats() string {
	return strings.Join(validStats, ", ")
}