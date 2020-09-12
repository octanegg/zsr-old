package admin

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/octanegg/core/internal/config"
	"github.com/octanegg/core/internal/deprecated"
	"github.com/octanegg/core/octane"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (h *handler) ImportMatches(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get(config.HeaderContentType) != config.HeaderApplicationJSON {
		w.WriteHeader(http.StatusUnsupportedMediaType)
		json.NewEncoder(w).Encode(Error{time.Now(), config.ErrInvalidContentType})
		return
	}

	var linkages []deprecated.EventLinkage
	if err := json.NewDecoder(r.Body).Decode(&linkages); err != nil {
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

		matches, err := h.Deprecated.GetMatches(&linkage)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(Error{time.Now(), err.Error()})
			return
		}

		gameMap, err := h.Deprecated.GetGameMap(linkage.OldEvent)
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

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
}

func (h *handler) parseGame(game *deprecated.Game) *octane.Game {
	newGame := &octane.Game{
		OctaneID: game.OctaneID,
		Number:   game.Number,
		Map:      game.Map,
		Duration: game.Duration,
		Blue: &octane.GameSide{
			Goals:   game.Blue.Goals,
			Team:    &octane.TeamStats{},
			Players: []*octane.PlayerStats{},
		},
		Orange: &octane.GameSide{
			Goals:   game.Orange.Goals,
			Team:    &octane.TeamStats{},
			Players: []*octane.PlayerStats{},
		},
	}

	newGame.Blue.Winner = newGame.Blue.Goals > newGame.Orange.Goals
	newGame.Orange.Winner = newGame.Orange.Goals > newGame.Blue.Goals

	newGame.Blue.Team.ID = h.findOrInsertTeam(game.Blue.Name)
	newGame.Orange.Team.ID = h.findOrInsertTeam(game.Orange.Name)

	for _, log := range game.Blue.Players {
		newGame.Blue.Players = append(newGame.Blue.Players, &octane.PlayerStats{
			Player: h.findOrInsertPlayer(log.Player),
			Stats: struct {
				Core deprecated.Log `bson:"core"`
			}{log},
			Rating: log.Rating,
		})
	}

	for _, log := range game.Orange.Players {
		newGame.Orange.Players = append(newGame.Orange.Players, &octane.PlayerStats{
			Player: h.findOrInsertPlayer(log.Player),
			Stats: struct {
				Core deprecated.Log `bson:"core"`
			}{log},
			Rating: log.Rating,
		})
	}

	return newGame
}

func (h *handler) parseMatch(match *deprecated.Match) *octane.Match {
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
