package deprecated

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/octanegg/zsr/internal/config"
	"github.com/octanegg/zsr/octane"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// EventLinkage .
type EventLinkage struct {
	OldEvent int    `json:"old_event"`
	OldStage int    `json:"old_stage"`
	NewEvent string `json:"new_event"`
	NewStage int    `json:"new_stage"`
}

func (h *handler) ImportMatches(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get(config.HeaderContentType) != config.HeaderApplicationJSON {
		w.WriteHeader(http.StatusUnsupportedMediaType)
		json.NewEncoder(w).Encode(Error{time.Now(), config.ErrInvalidContentType})
		return
	}

	var events []int
	if err := json.NewDecoder(r.Body).Decode(&events); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Error{time.Now(), err.Error()})
		return
	}

	linkages, err := h.Deprecated.getLinkages(events)
	if  err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Error{time.Now(), err.Error()})
		return
	}

	for _, linkage := range linkages {
		eventID, err := primitive.ObjectIDFromHex(linkage.NewEvent)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(Error{time.Now(), err.Error()})
			return
		}

		matches, err := h.Deprecated.getLinkageMatches(linkage)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(Error{time.Now(), err.Error()})
			return
		}

		gameMap, err := h.Deprecated.getGameMap(linkage.OldEvent)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(Error{time.Now(), err.Error()})
			return
		}

		for _, match := range matches {
			newMatch := h.parseMatch(match)
			newMatch.EventID = &eventID
			newMatch.Stage = linkage.NewStage

			matchID, err := h.upsertMatch(newMatch)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(Error{time.Now(), err.Error()})
				return
			}

			newMatch.ID = matchID

			if games, ok := gameMap[newMatch.OctaneID]; ok {
				for _, gameData := range games {
					game := h.parseGame(gameData)
					game.MatchID = matchID
					game.EventID = &eventID
					game.Mode = match.Mode
					game.Date = match.Date

					if err = h.upsertGame(game, newMatch, false); err != nil {
						w.WriteHeader(http.StatusInternalServerError)
						json.NewEncoder(w).Encode(Error{time.Now(), err.Error()})
						return
					}
				}
			}
		}

	}

	w.WriteHeader(http.StatusOK)
}

func (h *handler) parseGame(game *Game) *octane.Game {
	newGame := &octane.Game{
		OctaneID: game.OctaneID,
		Number:   game.Number,
		Map:      game.Map,
		Duration: game.Duration,
		Blue: &octane.GameSide{
			Goals:   game.Blue.Goals,
			Players: []*octane.PlayerStats{},
		},
		Orange: &octane.GameSide{
			Goals:   game.Orange.Goals,
			Players: []*octane.PlayerStats{},
		},
	}

	newGame.Blue.Winner = newGame.Blue.Goals > newGame.Orange.Goals
	newGame.Orange.Winner = newGame.Orange.Goals > newGame.Blue.Goals

	newGame.Blue.Team = h.findOrInsertTeam(game.Blue.Name)
	newGame.Orange.Team = h.findOrInsertTeam(game.Orange.Name)

	for _, log := range game.Blue.Players {
		newGame.Blue.Players = append(newGame.Blue.Players, &octane.PlayerStats{
			Player: h.findOrInsertPlayer(log.Player),
			Stats: struct {
				Core Log `bson:"core"`
			}{log},
			Rating: log.Rating,
		})
	}

	for _, log := range game.Orange.Players {
		newGame.Orange.Players = append(newGame.Orange.Players, &octane.PlayerStats{
			Player: h.findOrInsertPlayer(log.Player),
			Stats: struct {
				Core Log `bson:"core"`
			}{log},
			Rating: log.Rating,
		})
	}

	return newGame
}

func (h *handler) parseMatch(match *Match) *octane.Match {
	newMatch := octane.Match{
		OctaneID: match.OctaneID,
		Date:     match.Date,
		Format:   match.Format,
		Mode:     match.Mode,
		Number:   match.Number,
		Blue: &octane.MatchSide{
			Score:  match.Blue.Score,
			Winner: match.Blue.Winner,
		},
		Orange: &octane.MatchSide{
			Score:  match.Orange.Score,
			Winner: match.Orange.Winner,
		},
	}

	if match.Blue.Name != "" && match.Orange.Name != "" {
		newMatch.Blue.Team = h.findOrInsertTeam(match.Blue.Name)
		newMatch.Orange.Team = h.findOrInsertTeam(match.Orange.Name)
	}

	return &newMatch
}

func (d *deprecated) getLinkages(events []int) ([]*EventLinkage, error) {
	stmt := "SELECT old_event, old_stage, new_event, new_stage FROM mapping"
	if len(events) > 0 {
		stmt += " WHERE old_event IN ("
		for i, event := range events {
			stmt += strconv.Itoa(event)
			if i != len(events) - 1 {
				stmt += ","
			} 
		}
		stmt += ")"
	}
	
	results, err := d.DB.Query(stmt)
	if err != nil {
		return nil, err
	}

	var linkages []*EventLinkage
	for results.Next() {
		var linkage EventLinkage
		err = results.Scan(&linkage.OldEvent, &linkage.OldStage, &linkage.NewEvent, &linkage.NewStage)
		if err != nil {
			return nil, err
		}

		linkages = append(linkages, &linkage)
	}

	return linkages, nil
}

func (d *deprecated) getGameMap(eventID int) (map[string]map[int]*Game, error) {
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

func (d *deprecated) getLinkageMatches(l *EventLinkage) ([]*Match, error) {
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

func (h *handler) upsertGame(newGame *octane.Game, match *octane.Match, update bool) error {
	data, err := h.Octane.FindGames(bson.M{"match": newGame.MatchID, "number": newGame.Number}, nil, nil)
	if err != nil {
		return err
	}

	if len(data.Data) == 0 {
		if _, err = h.Octane.InsertGame(newGame); err != nil {
			return fmt.Errorf("error inserting game - %s", err.Error())
		}
	} else {
		game := data.Data[0].(octane.Game)
		newGame.ID = game.ID
		if _, err = h.Octane.ReplaceGame(game.ID, newGame); err != nil {
			return fmt.Errorf("error updating game - %s", err.Error())
		}
		if game.Blue.Winner {
			match.Blue.Score--
		} else if game.Orange.Winner {
			match.Orange.Score--
		}
	}

	if newGame.Blue.Winner {
		match.Blue.Score++
	} else if newGame.Orange.Winner {
		match.Orange.Score++
	}

	match.Blue.Winner = match.Blue.Score > match.Orange.Score
	match.Orange.Winner = match.Orange.Score > match.Blue.Score

	if update {
		if _, err := h.Octane.ReplaceMatch(match.ID, match); err != nil {
			return fmt.Errorf("error updating match score - %s", err.Error())
		}
	}

	return nil
}

// TODO: Move away from octane ID later
func (h *handler) upsertMatch(newMatch *octane.Match) (*primitive.ObjectID, error) {
	data, err := h.Octane.FindMatches(bson.M{"octane_id": newMatch.OctaneID}, nil, nil)
	if err != nil {
		return nil, err
	}

	if len(data.Data) == 0 {
		id, err := h.Octane.InsertMatch(newMatch)
		if err != nil {
			return nil, fmt.Errorf("error inserting match - %s", err.Error())
		}
		return id, nil
	}

	id := data.Data[0].(octane.Match).ID
	newMatch.ID = id
	if _, err = h.Octane.ReplaceMatch(id, newMatch); err != nil {
		return nil, fmt.Errorf("error updating match - %s", err.Error())
	}

	return id, nil
}

func (h *handler) findOrInsertTeam(name string) *octane.Team {
	teams, err := h.Octane.FindTeams(bson.M{"name": name}, nil, nil)
	if err != nil || len(teams.Data) == 0 {
		team, _ := h.Octane.InsertTeam(&octane.Team{
			Name: name,
		})
		return &octane.Team{
			ID:   team,
			Name: name,
		}
	}

	team := teams.Data[0].(octane.Team)
	return &octane.Team{
		ID: team.ID,
		Name: team.Name,
	}
}

func (h *handler) findOrInsertPlayer(tag string) *octane.Player {
	players, err := h.Octane.FindPlayers(bson.M{"tag": tag}, nil, nil)
	if err != nil || len(players.Data) == 0 {
		player, _ := h.Octane.InsertPlayer(&octane.Player{
			Tag: tag,
		})
		return &octane.Player{
			ID:  player,
			Tag: tag,
		}
	}

	player := players.Data[0].(octane.Player)
	return &octane.Player{
		ID: player.ID,
		Tag: player.Tag,
	}
}
