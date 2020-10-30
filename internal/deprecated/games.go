package deprecated

import (
	"sort"
	"strconv"
	"strings"
	"time"
)

// Metadata .
type Metadata struct {
	Event int
	Stage int
	Match int
}

// Game .
type Game struct {
	OctaneID string     `json:"octane_id"`
	Date     *time.Time `json:"date"`
	Number   int        `json:"number"`
	Map      string     `json:"map"`
	Duration int        `json:"duration"`
	Blue     GameTeam   `json:"blue"`
	Orange   GameTeam   `json:"orange"`
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

// Averages .
type Averages struct {
	Score   float64 `json:"score"`
	Goals   float64 `json:"goals"`
	Assists float64 `json:"assists"`
	Saves   float64 `json:"saves"`
	Shots   float64 `json:"shots"`
	SP      float64 `json:"shooting_percentage"`
	GP      float64 `json:"goal_participation"`
}

// DeleteGameContext .
type DeleteGameContext struct {
	OctaneID string `json:"octane_id"`
	Number   int    `json:"number"`
}

// GetGamesContext .
type GetGamesContext struct {
	OctaneID string `json:"octane_id"`
	Blue     string `json:"blue"`
	Orange   string `json:"orange"`
}

func (d *deprecated) DeleteGame(ctx *DeleteGameContext) error {
	stmt := "SELECT CASE m.Result WHEN s.Team1 THEN 1 WHEN s.Team2 THEN 2 ELSE 0 END AS Result FROM Matches2 m, Series s WHERE m.match_url = s.match_url AND m.match_url = ? AND m.Game = ?"
	row := d.DB.QueryRow(stmt, ctx.OctaneID, ctx.Number)

	stmt = "DELETE FROM Matches2 WHERE match_url = ? AND Game = ?"
	_, err := d.DB.Exec(stmt, ctx.OctaneID, ctx.Number)
	if err != nil {
		return err
	}

	stmt = "DELETE FROM Logs WHERE match_url = ? AND Game = ?"
	_, err = d.DB.Exec(stmt, ctx.OctaneID, ctx.Number)
	if err != nil {
		return err
	}

	var winner int
	if err = row.Scan(&winner); err != nil {
		return err
	}

	if winner == 1 {
		if _, err = d.DB.Exec("UPDATE `Series` SET `Team1Games` = `Team1Games` - 1, `Result` = CASE WHEN `Team1Games` > `Team2Games` THEN `Team1` ELSE `Team2` END WHERE `match_url` = ?", ctx.OctaneID); err != nil {
			return err
		}
	} else if winner == 2 {
		if _, err = d.DB.Exec("UPDATE `Series` SET `Team2Games` = `Team2Games` - 1, `Result` = CASE WHEN `Team1Games` > `Team2Games` THEN `Team1` ELSE `Team2` END WHERE `match_url` = ?", ctx.OctaneID); err != nil {
			return err
		}
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

func (d *deprecated) InsertGame(game *Game) error {
	md := getMetadata(game.OctaneID)
	blue := getTeamStats(game.Blue.Players)
	orange := getTeamStats(game.Orange.Players)

	if err := d.insertGamePlayers(md, game, blue, orange); err != nil {
		return err
	}

	if err := d.insertGameTeams(md, game, blue, orange); err != nil {
		return err
	}

	if blue.Goals > orange.Goals {
		if _, err := d.DB.Exec("UPDATE `Series` SET `Team1Games` = `Team1Games` + 1, `Result` = CASE WHEN `Team1Games` > `Team2Games` THEN `Team1` ELSE `Team2` END WHERE `match_url` = ?", game.OctaneID); err != nil {
			return err
		}
	} else if orange.Goals > blue.Goals {
		if _, err := d.DB.Exec("UPDATE `Series` SET `Team2Games` = `Team2Games` + 1, `Result` = CASE WHEN `Team1Games` > `Team2Games` THEN `Team1` ELSE `Team2` END WHERE `match_url` = ?", game.OctaneID); err != nil {
			return err
		}
	}

	return nil
}

func (d *deprecated) insertGameTeams(md *Metadata, game *Game, blue, orange *Log) error {
	stmt := "INSERT INTO Matches2(Event, Stage, `Match`, Game, Date, Map, Length, Team, Vs, Result, Winner, TeamScore, TeamGoals, TeamAssists, TeamSaves, TeamShots, OppScore, OppGoals, OppAssists, OppSaves, OppShots, match_url) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"

	var result string
	var blueWin, orangeWin int
	if orange.Goals > blue.Goals {
		result = game.Orange.Name
		orangeWin = 1
	} else if blue.Goals > orange.Goals {
		result = game.Blue.Name
		blueWin = 1
	}

	_, err := d.DB.Exec(stmt, md.Event, md.Stage, md.Match, game.Number, game.Date.Format("2006-01-02"), game.Map, game.Duration, game.Blue.Name, game.Orange.Name, result, blueWin, blue.Score, blue.Goals, blue.Assists, blue.Saves, blue.Shots, orange.Score, orange.Goals, orange.Assists, orange.Saves, orange.Shots, game.OctaneID)
	if err != nil {
		return err
	}

	_, err = d.DB.Exec(stmt, md.Event, md.Stage, md.Match, game.Number, game.Date.Format("2006-01-02"), game.Map, game.Duration, game.Orange.Name, game.Blue.Name, result, orangeWin, orange.Score, orange.Goals, orange.Assists, orange.Saves, orange.Shots, blue.Score, blue.Goals, blue.Assists, blue.Saves, blue.Shots, game.OctaneID)
	if err != nil {
		return err
	}

	return nil
}

func (d *deprecated) insertGamePlayers(md *Metadata, game *Game, blue, orange *Log) error {
	stmt := "INSERT INTO Logs(Event, Stage, `Match`, Game, Date, Map, Team, Vs, Result, Winner, Player, Score, Goals, Assists, Saves, Shots, TeamScore, TeamGoals, TeamAssists, TeamSaves, TeamShots, SP, MVP, HT, PM, SAV, Rating, match_url) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"

	avg, err := d.getAverages()
	if err != nil {
		return err
	}

	var result string
	var blueWin, orangeWin int
	if orange.Goals > blue.Goals {
		result = game.Orange.Name
		orangeWin = 1
	} else if blue.Goals > orange.Goals {
		result = game.Blue.Name
		blueWin = 1
	}

	sort.Slice(game.Blue.Players, func(i, j int) bool {
		return game.Blue.Players[i].Score > game.Blue.Players[j].Score
	})

	sort.Slice(game.Orange.Players, func(i, j int) bool {
		return game.Orange.Players[i].Score > game.Orange.Players[j].Score
	})

	for i, player := range game.Blue.Players {
		if strings.TrimSpace(player.Player) == "" {
			continue
		}

		if player.Shots > 0 {
			player.SP = float64(player.Goals) / float64(player.Shots)
		}
		if blue.Goals > 0 {
			player.GP = float64(player.Goals+player.Assists) / float64(blue.Goals)
		}
		player.Rating = getRating(avg, &player)

		var mvp int
		if i == 0 && blueWin == 1 {
			mvp = 1
		}

		var ht, pm, sav int
		if player.Goals >= 3 {
			ht = 1
		}
		if player.Assists >= 3 {
			pm = 1
		}
		if player.Saves >= 3 {
			sav = 1
		}

		if player.Score == 0 {
			player.Score = d.getAverageScore(player)
		}

		_, err := d.DB.Exec(stmt, md.Event, md.Stage, md.Match, game.Number, game.Date.Format("2006-01-02"), game.Map, game.Blue.Name, game.Orange.Name, result, blueWin, player.Player, player.Score, player.Goals, player.Assists, player.Saves, player.Shots, blue.Score, blue.Goals, blue.Assists, blue.Saves, blue.Shots, i+1, mvp, ht, pm, sav, player.Rating, game.OctaneID)
		if err != nil {
			return err
		}
	}

	for i, player := range game.Orange.Players {
		if strings.TrimSpace(player.Player) == "" {
			continue
		}

		if player.Shots > 0 {
			player.SP = float64(player.Goals) / float64(player.Shots)
		}
		if orange.Goals > 0 {
			player.GP = float64(player.Goals+player.Assists) / float64(orange.Goals)
		}
		player.Rating = getRating(avg, &player)

		var mvp int
		if i == 0 && orangeWin == 1 {
			mvp = 1
		}

		var ht, pm, sav int
		if player.Goals >= 3 {
			ht = 1
		}
		if player.Assists >= 3 {
			pm = 1
		}
		if player.Saves >= 3 {
			sav = 1
		}

		if player.Score == 0 {
			player.Score = d.getAverageScore(player)
		}

		_, err := d.DB.Exec(stmt, md.Event, md.Stage, md.Match, game.Number, game.Date.Format("2006-01-02"), game.Map, game.Orange.Name, game.Blue.Name, result, orangeWin, player.Player, player.Score, player.Goals, player.Assists, player.Saves, player.Shots, orange.Score, orange.Goals, orange.Assists, orange.Saves, orange.Shots, i+1, mvp, ht, pm, sav, player.Rating, game.OctaneID)
		if err != nil {
			return err
		}
	}

	return nil
}

func (d *deprecated) getAverages() (*Averages, error) {
	stmt := "SELECT AVG(Score), AVG(Goals), AVG(Assists), AVG(Shots), AVG(Saves), SUM(Goals) / SUM(Shots), SUM(Goals + Assists) / SUM(TeamGoals) FROM Logs, octane.Events WHERE event_id = Event AND mode = 3 AND Date > '2019-01-01'"
	res := d.DB.QueryRow(stmt)

	var avg Averages
	if err := res.Scan(&avg.Score, &avg.Goals, &avg.Assists, &avg.Saves, &avg.Shots, &avg.SP, &avg.GP); err != nil {
		return nil, err
	}

	return &avg, nil
}

func (d *deprecated) getAverageScore(log Log) int {
	stmt := "SELECT AVG(Score) FROM Logs, octane.Events WHERE event_id = Event AND mode = 3 AND Date > '2019-01-01' AND Goals = ? AND Assists = ? AND Saves = ? AND Shots = ?"
	res := d.DB.QueryRow(stmt, log.Goals, log.Assists, log.Saves, log.Shots)

	var avg float64
	if err := res.Scan(&avg); err != nil {
		return 0
	}

	return int(avg)
}

func getTeamStats(players []Log) *Log {
	log := &Log{}
	for _, player := range players {
		log.Score += player.Score
		log.Goals += player.Goals
		log.Assists += player.Assists
		log.Saves += player.Saves
		log.Shots += player.Shots
	}
	return log
}

func getMetadata(id string) *Metadata {
	md := &Metadata{}
	md.Event, _ = strconv.Atoi(id[0:3])
	md.Stage, _ = strconv.Atoi(id[3:5])
	md.Match, _ = strconv.Atoi(id[5:7])
	return md
}

func getRating(avg *Averages, log *Log) float64 {
	return (float64(log.Score)/avg.Score + float64(log.Goals)/avg.Goals + float64(log.Assists)/avg.Assists + float64(log.Saves)/avg.Saves + float64(log.Shots)/avg.Shots + log.GP/avg.GP + log.SP/avg.SP) / 7.0
}
