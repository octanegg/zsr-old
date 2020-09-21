package deprecated

import (
	"fmt"
	"time"
)

// Match .
type Match struct {
	OctaneID string `json:"octane_id"`
	Event    string `json:"event"`
	Stage    int `json:"stage"`
	Substage int `json:"substage"`
	Date     *time.Time `json:"date"`
	Format   string `json:"format"`
	Blue     *Team `json:"blue"`
	Orange   *Team `json:"orange"`
	Mode     int `json:"mode"`
	Number   int `json:"number"`
}

// Team .
type Team struct {
	Name   string `json:"name"`
	Score  int    `json:"score"`
	Winner bool   `json:"winner"`
}

// GetMatchesContext .
type GetMatchesContext struct {
	Event string    `json:"event"`
	Stage string    `json:"stage"`
}

// GetMatchContext .
type GetMatchContext struct {
	OctaneID string `json:"octane_id"`
}

// UpdateMatchContext .
type UpdateMatchContext struct {
	OctaneID   string `json:"octane_id"`
	Team1      Mapping `json:"blue"`
	Team2      Mapping `json:"orange"`
	Team1Score int `json:"blue_score"`
	Team2Score int `json:"orange_score"`
}

// Mapping .
type Mapping struct {
	Old string `json:"old"`
	New string `json:"new"`
}

func (d *deprecated) GetMatches(ctx *GetMatchesContext) ([]*Match, error) {
	query := fmt.Sprintf("SELECT match_url, Time, best_of, Team1, Team2, Team1Games, Team2Games FROM Series WHERE Event = %s AND Stage = %s", ctx.Event, ctx.Stage)
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
		match.Blue = &blue
		match.Orange = &orange

		matches = append(matches, &match)
	}

	return matches, nil
}


func (d *deprecated) GetMatch(ctx *GetMatchContext) (*Match, error) {
	query := fmt.Sprintf("SELECT match_url, Time, best_of, Team1, Team2, Team1Games, Team2Games FROM Series WHERE match_url = '%s'", ctx.OctaneID)
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

		match.Blue = &blue
		match.Orange = &orange

		matches = append(matches, &match)
	}

	if matches == nil || len(matches) == 0 {
		return nil, nil
	}

	return matches[0], nil
}

func (d *deprecated) UpdateMatches(ctxs []*UpdateMatchContext) error {
	for _, ctx := range ctxs {
		winner := ctx.Team1.New
		if ctx.Team2Score > ctx.Team1Score {
			winner = ctx.Team2.New
		}

		stmt := "UPDATE Series SET Team1 = ?, Team2 = ?, Team1Games = ?, Team2Games = ?, Result = ? WHERE match_url = ?"
		_, err := d.DB.Exec(stmt, ctx.Team1.New, ctx.Team2.New, ctx.Team1Score, ctx.Team2Score, winner, ctx.OctaneID)
		if err != nil {
			return err
		}

		if err = d.changeTeamName(ctx.OctaneID, ctx.Team1.Old, ctx.Team1.New); err != nil {
			return err
		}

		if err = d.changeTeamName(ctx.OctaneID, ctx.Team2.Old, ctx.Team2.New); err != nil {
			return err
		}
	}

	return nil
}

func (d *deprecated) changeTeamName(id, old, new string) error {
	if old != new {
		for _, field := range []string{"Team", "Vs", "Result"} {
			stmt := fmt.Sprintf("UPDATE Logs SET %s = ? WHERE %s = ? AND match_url = ?", field, field)
			if _, err := d.DB.Exec(stmt, new, old, id); err != nil {
				return err
			}

			stmt = fmt.Sprintf("UPDATE Matches2 SET %s = ? WHERE %s = ? AND match_url = ?", field, field)
			if _, err := d.DB.Exec(stmt, new, old, id); err != nil {
				return err
			}
		}
	}

	return nil
}