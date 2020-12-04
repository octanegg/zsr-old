package stats

import (
	"strings"

	"github.com/octanegg/zsr/octane/collection"
	"go.mongodb.org/mongo-driver/bson"
)

var (
	validStats = []string{"score", "goals", "assists", "saves", "shots", "rating"}
)

// Stats .
type Stats interface {
	GetGameRecords(bson.M, bson.M) ([]interface{}, error)
	GetPlayerAggregate(string, bson.M) ([]*PlayerAggregate, error)
}

type stats struct {
	Statlines collection.Collection
}

// New .
func New(c collection.Collection) Stats {
	return &stats{c}
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
