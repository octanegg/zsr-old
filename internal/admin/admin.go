package admin

import (
	"fmt"
	"net/http"
	"time"

	"github.com/octanegg/core/internal/config"
	"github.com/octanegg/core/internal/deprecated"
	"github.com/octanegg/core/octane"
	"github.com/octanegg/racer"
	"github.com/octanegg/slimline"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Error .
type Error struct {
	Timestamp time.Time `json:"timestamp"`
	Error     string    `json:"error"`
}

type handler struct {
	Octane     octane.Client
	Racer      racer.Racer
	Slimline   slimline.Slimline
	Deprecated deprecated.Deprecated
}

// Handler .
type Handler interface {
	LinkBallchasing(http.ResponseWriter, *http.Request)
	ImportMatches(http.ResponseWriter, *http.Request)
	UpdateMatch(http.ResponseWriter, *http.Request)
	GetMatch(http.ResponseWriter, *http.Request)
	ResetGame(http.ResponseWriter, *http.Request)
}

// New .
func New(o octane.Client, r racer.Racer, s slimline.Slimline, d deprecated.Deprecated) Handler {
	return &handler{o, r, s, d}
}

func (h *handler) upsertGame(newGame *octane.Game, match *octane.Match, update bool) error {
	data, err := h.Octane.FindGames(bson.M{config.ParamMatch: newGame.MatchID, config.ParamNumber: newGame.Number}, nil, nil)
	if err != nil {
		return err
	}

	if len(data.Data) == 0 {
		if _, err = h.Octane.InsertGame(newGame); err != nil {
			return fmt.Errorf("error inserting game - %s", err.Error())
		}
	} else {
		game := data.Data[0].(octane.Game)
		if _, err = h.Octane.UpdateGame(game.ID, newGame); err != nil {
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
		if _, err := h.Octane.UpdateMatch(match.ID, match); err != nil {
			return fmt.Errorf("error updating match score - %s", err.Error())
		}
	}

	return nil
}

// TODO: Move away from octane ID later
func (h *handler) upsertMatch(newMatch *octane.Match) (*primitive.ObjectID, error) {
	data, err := h.Octane.FindMatches(bson.M{config.ParamOctaneID: newMatch.OctaneID}, nil, nil)
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
	if _, err = h.Octane.UpdateMatch(id, newMatch); err != nil {
		return nil, fmt.Errorf("error updating match - %s", err.Error())
	}

	return id, nil
}

func (h *handler) findOrInsertTeam(name string) *primitive.ObjectID {
	teams, err := h.Octane.FindTeams(bson.M{config.ParamName: name}, nil, nil)
	if err != nil || len(teams.Data) == 0 {
		team, _ := h.Octane.InsertTeam(&octane.Team{
			Name: name,
		})
		return team
	}
	return teams.Data[0].(octane.Team).ID
}

func (h *handler) findOrInsertPlayer(tag string) *primitive.ObjectID {
	players, err := h.Octane.FindPlayers(bson.M{config.ParamTag: tag}, nil, nil)
	if err != nil || len(players.Data) == 0 {
		player, _ := h.Octane.InsertPlayer(&octane.Player{
			Tag: tag,
		})
		return player
	}
	return players.Data[0].(octane.Player).ID
}
