package records

import "strings"

var (
	validStats = []string{"score", "goals", "assists", "saves", "shots", "rating"}
)

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
