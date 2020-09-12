package admin

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/octanegg/core/internal/config"
	"github.com/octanegg/core/octane"
	"github.com/octanegg/racer"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// BallchasingLinkage .
type BallchasingLinkage struct {
	Match       string `json:"match"`
	Game        int    `json:"game"`
	Ballchasing string `json:"ballchasing"`
	SwapTeams   bool   `json:"swap_teams"`
	Propogate   bool   `json:"propogate"`
}

// AccountError .
type AccountError struct {
	Player    string `json:"player"`
	AccountID string `json:"account_id"`
	Platform  string `json:"platform"`
}

func (h *handler) LinkBallchasing(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get(config.HeaderContentType) != config.HeaderApplicationJSON {
		w.WriteHeader(http.StatusUnsupportedMediaType)
		json.NewEncoder(w).Encode(Error{time.Now(), config.ErrInvalidContentType})
		return
	}

	var linkage BallchasingLinkage
	if err := json.NewDecoder(r.Body).Decode(&linkage); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Error{time.Now(), err.Error()})
		return
	}

	oid, err := primitive.ObjectIDFromHex(linkage.Match)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Error{time.Now(), err.Error()})
		return
	}

	replay, err := h.Racer.GetReplay(linkage.Ballchasing)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Error{time.Now(), err.Error()})
		return
	}

	match, err := h.Octane.FindMatch(&oid)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Error{time.Now(), err.Error()})
		return
	}

	id := primitive.NewObjectID()
	game := &octane.Game{
		ID:            &id,
		OctaneID:      match.OctaneID,
		Number:        linkage.Game,
		MatchID:       &oid,
		EventID:       match.EventID,
		Map:           replay.MapName,
		Duration:      300 + replay.OvertimeSeconds,
		Mode:          match.Mode,
		Date:          match.Date,
		BallchasingID: linkage.Ballchasing,
		Blue: &octane.GameSide{
			Goals: replay.Blue.Stats.Core.Goals,
			Team: &octane.TeamStats{
				ID: match.Blue.Team,
			},
		},
		Orange: &octane.GameSide{
			Goals: replay.Orange.Stats.Core.Goals,
			Team: &octane.TeamStats{
				ID: match.Orange.Team,
			},
		},
	}

	bluePlayers, account := h.getPlayers(replay.Blue.Players)
	if account != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(err)
		return
	}

	orangePlayers, account := h.getPlayers(replay.Orange.Players)
	if account != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(err)
		return
	}

	game.Blue.Players = bluePlayers
	game.Orange.Players = orangePlayers
	game.Blue.Winner = game.Blue.Goals > game.Orange.Goals
	game.Orange.Winner = game.Orange.Goals > game.Blue.Goals

	if linkage.SwapTeams {
		game.Blue, game.Orange = game.Orange, game.Blue
	}

	if err = h.upsertGame(game, match, true); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Error{time.Now(), err.Error()})
		return
	}

	// TODO: Remove when old database is no more
	if linkage.Propogate {
		blue, err := h.Octane.FindTeam(game.Blue.Team.ID)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(Error{time.Now(), err.Error()})
		}

		orange, err := h.Octane.FindTeam(game.Orange.Team.ID)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(Error{time.Now(), err.Error()})
		}

		m := map[*primitive.ObjectID]string{
			game.Blue.Team.ID:   blue.Name,
			game.Orange.Team.ID: orange.Name,
		}

		for _, player := range append(game.Blue.Players, game.Orange.Players...) {
			p, err := h.Octane.FindPlayer(player.Player)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(Error{time.Now(), err.Error()})
			}
			m[player.Player] = p.Tag
		}

		if err = h.Deprecated.Propogate(game, m); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(Error{time.Now(), err.Error()})
		}
	}

	w.WriteHeader(http.StatusOK)
}

func (h *handler) getPlayers(players []racer.Player) ([]*octane.PlayerStats, *AccountError) {
	var p []*octane.PlayerStats
	for _, player := range players {
		res, err := h.Octane.FindPlayers(bson.M{config.ParamAccountID: player.ID.ID}, nil, nil)
		if err != nil || len(res.Data) == 0 {
			return nil, &AccountError{player.Name, player.ID.ID, player.ID.Platform}
		}

		data := res.Data[0].(octane.Player)
		p = append(p, &octane.PlayerStats{
			Player: data.ID,
			Stats:  player.Stats,
		})
	}
	return p, nil
}
