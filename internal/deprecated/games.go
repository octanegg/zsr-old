package deprecated

import (
	"fmt"
	"sort"
)

// Game .
type Game struct {
	OctaneID string   `json:"octane_id"`
	Number   int      `json:"number"`
	Map      string   `json:"map"`
	Duration int      `json:"duration"`
	Blue     GameTeam `json:"blue"`
	Orange   GameTeam `json:"orange"`
}

// GameTeam .
type GameTeam struct {
	Name    string `json:"name"`
	Winner  bool   `json:"winner"`
	Goals   int    `json:"goals"`
	Players []Log  `json:"players"`
}

// Log .
type Log struct {
	Player   string  `json:"player"`
	OctaneID string  `json:"-"`
	Team     string  `json:"-"`
	Number   int     `json:"-"`
	MVP      bool    `json:"mvp"`
	Score    int     `json:"score"`
	Goals    int     `json:"goals"`
	Assists  int     `json:"assists"`
	Saves    int     `json:"saves"`
	Shots    int     `json:"shots"`
	SP       float64 `json:"shooting_percentage"`
	GP       float64 `json:"goal_participation"`
	Rating   float64 `json:"-"`
}

// ResetGameContext .
type ResetGameContext struct {
	OctaneID string `json:"octane_id"`
	Number int `json:"number"`
}

// GetGamesContext .
type GetGamesContext struct {
	OctaneID string `json:"octane_id"`
	Blue string `json:"blue"`
	Orange string `json:"orange"`
}

func (d *deprecated) GetGameMap(eventID int) (map[string]map[int]*Game, error) {
	results, err := d.DB.Query(fmt.Sprintf("SELECT match_url, Map, Length, Team, Vs, TeamGoals, OppGoals, Game FROM Matches2 WHERE Event = %d GROUP BY match_url, Game ORDER BY match_url ASC, Game ASC", eventID))
	if err != nil {
		return nil, err
	}

	m := make(map[string]map[int]*Game)
	for results.Next() {
		var game Game
		var blue, orange GameTeam
		err = results.Scan(&game.OctaneID, &game.Map, &game.Duration, &blue.Name, &orange.Name, &blue.Goals, &orange.Goals, &game.Number)
		if err != nil {
			return nil, err
		}

		if blue.Goals > orange.Goals {
			blue.Winner = true
		} else if orange.Goals > blue.Goals {
			orange.Winner = true
		}

		game.Blue = blue
		game.Orange = orange

		if m[game.OctaneID] == nil {
			m[game.OctaneID] = make(map[int]*Game)
		}

		m[game.OctaneID][game.Number] = &game
	}

	results, err = d.DB.Query(fmt.Sprintf("SELECT match_url, Team, Player, Score, Goals, Assists, Saves, Shots, TeamGoals, Rating, MVP, Game FROM Logs WHERE Event = %d ORDER BY match_url ASC, Game ASC, Score DESC", eventID))
	if err != nil {
		return nil, err
	}

	for results.Next() {
		var log Log
		var tGoals *int
		var mvp *int

		err = results.Scan(&log.OctaneID, &log.Team, &log.Player, &log.Score, &log.Goals, &log.Assists, &log.Saves, &log.Shots, &tGoals, &log.Rating, &mvp, &log.Number)
		if err != nil {
			return nil, err
		}

		if log.Shots > 0 {
			log.SP = float64(log.Goals) / float64(log.Shots)
		}

		if *tGoals > 0 {
			log.GP = (float64(log.Goals) + float64(log.Assists)) / float64(*tGoals)
		}

		if *mvp == 1 {
			log.MVP = true
		}

		if m[log.OctaneID][log.Number].Blue.Name == log.Team {
			m[log.OctaneID][log.Number].Blue.Players = append(m[log.OctaneID][log.Number].Blue.Players, log)
		} else {
			m[log.OctaneID][log.Number].Orange.Players = append(m[log.OctaneID][log.Number].Orange.Players, log)
		}

	}

	return m, nil
}


func (d *deprecated) ResetGame(ctx *ResetGameContext) error {
	stmt := "DELETE FROM Matches2 WHERE match_url = ? AND Game = ?"
	_, err := d.DB.Exec(stmt, ctx.OctaneID, ctx.Number)
	if err != nil {
		return err
	}


	stmt = "DELETE FROM Logs WHERE match_url = ? AND Game = ?"
	_, err = d.DB.Exec(stmt, ctx.OctaneID, ctx.Number)
	if err != nil {
		return err
	}

	return nil
}

func (d *deprecated) GetGames(ctx *GetGamesContext) ([]*Game, error) {
	stmt := "SELECT match_url, Map, Length, Team, Vs, TeamGoals, OppGoals, Game FROM Matches2 WHERE match_url = ?"
	results, err := d.DB.Query(stmt, ctx.OctaneID)
	if err != nil {
		return nil, err
	}

	m := make(map[int]*Game)
	for results.Next() {
		var game Game
		var blue, orange GameTeam
		err = results.Scan(&game.OctaneID, &game.Map, &game.Duration, &blue.Name, &orange.Name, &blue.Goals, &orange.Goals, &game.Number)
		if err != nil {
			return nil, err
		}

		if blue.Goals > orange.Goals {
			blue.Winner = true
		} else if orange.Goals > blue.Goals {
			orange.Winner = true
		}

		game.Blue = blue
		game.Orange = orange

		if ctx.Blue != game.Blue.Name {
			game.Blue, game.Orange = game.Orange, game.Blue
		}

		m[game.Number] = &game
	}

	stmt = "SELECT match_url, Team, Player, Score, Goals, Assists, Saves, Shots, TeamGoals, Rating, MVP, Game FROM Logs WHERE match_url = ?"
	results, err = d.DB.Query(stmt, ctx.OctaneID)
	if err != nil {
		return nil, err
	}

	for results.Next() {
		var log Log
		var tGoals *int
		var mvp *int

		err = results.Scan(&log.OctaneID, &log.Team, &log.Player, &log.Score, &log.Goals, &log.Assists, &log.Saves, &log.Shots, &tGoals, &log.Rating, &mvp, &log.Number)
		if err != nil {
			return nil, err
		}

		if log.Shots > 0 {
			log.SP = float64(log.Goals) / float64(log.Shots)
		}

		if *tGoals > 0 {
			log.GP = (float64(log.Goals) + float64(log.Assists)) / float64(*tGoals)
		}

		if *mvp == 1 {
			log.MVP = true
		}

		if m[log.Number].Blue.Name == log.Team {
			m[log.Number].Blue.Players = append(m[log.Number].Blue.Players, log)
		} else {
			m[log.Number].Orange.Players = append(m[log.Number].Orange.Players, log)
		}

	}

	var games []*Game
	for _, v := range m {
		games = append(games, v)
	}

	sort.SliceStable(games, func(i, j int) bool {
		return games[i].Number < games[j].Number
	})

	return games, nil
}