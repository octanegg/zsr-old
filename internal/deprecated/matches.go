package deprecated

import (
	"fmt"
	"strconv"
	"time"
)

// EventLinkage .
type EventLinkage struct {
	OldEvent int    `json:"old_event"`
	OldStage int    `json:"old_stage"`
	NewEvent string `json:"new_event"`
	NewStage int    `json:"new_stage"`
}

// Match .
type Match struct {
	OctaneID string
	Event    string
	Stage    int
	Substage int
	Date     *time.Time
	Format   string
	Blue     *Team
	Orange   *Team
	Mode     int
	Number   int
}

// Team .
type Team struct {
	Name   string `bson:"name"`
	Score  int    `bson:"score"`
	Winner bool   `bson:"winner"`
}

func (d *deprecated) GetMatches(l *EventLinkage) ([]*Match, error) {
	query := fmt.Sprintf("SELECT match_url, Time, best_of, Team1, Team2, Team1Games, Team2Games FROM Series WHERE Event = %d AND Stage = %d", l.OldEvent, l.OldStage+1)
	results, err := d.DB.Query(query)
	if err != nil {
		return nil, err
	}

	var matches []*Match
	for results.Next() {
		var match Match
		var blue, orange Team
		err = results.Scan(&match.OctaneID, &match.Date, &match.Format, &blue.Name, &orange.Name, &blue.Score, &orange.Score)
		if err != nil {
			return nil, err
		}

		blue.Winner = blue.Score > orange.Score
		orange.Winner = orange.Score > blue.Score

		match.Event = l.NewEvent
		match.Stage = l.NewStage
		match.Mode = 3
		i, _ := strconv.Atoi(match.OctaneID[5:7])
		match.Number = i

		match.Blue = &blue
		match.Orange = &orange

		matches = append(matches, &match)
	}

	return matches, nil
}
