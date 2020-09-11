package deprecated

import "fmt"

// Game .
type Game struct {
	OctaneID string   `bson:"octane_id"`
	Number   int      `bson:"number"`
	Map      string   `bson:"map"`
	Duration int      `bson:"duration"`
	Blue     GameTeam `bson:"blue"`
	Orange   GameTeam `bson:"orange"`
}

// GameTeam .
type GameTeam struct {
	Name    string `bson:"name"`
	Winner  bool   `bson:"winner"`
	Goals   int    `bson:"goals"`
	Players []Log  `bson:"players"`
}

// Log .
type Log struct {
	Player   string  `bson:"-"`
	OctaneID string  `bson:"-"`
	Team     string  `bson:"-"`
	Number   int     `bson:"-"`
	MVP      bool    `bson:"mvp"`
	Score    int     `bson:"score"`
	Goals    int     `bson:"goals"`
	Assists  int     `bson:"assists"`
	Saves    int     `bson:"saves"`
	Shots    int     `bson:"shots"`
	SP       float64 `bson:"shooting_percentage"`
	GP       float64 `bson:"goal_participation"`
	Rating   float64 `bson:"-"`
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
