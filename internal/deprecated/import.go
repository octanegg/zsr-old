package deprecated

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/octanegg/zsr/ballchasing"
	"github.com/octanegg/zsr/internal/config"
	"github.com/octanegg/zsr/octane"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/sync/errgroup"
)

const (
	linkagesSQL = "SELECT old_event, old_stage, new_event, new_stage FROM mapping"
)

// EventLinkage .
type EventLinkage struct {
	OldEvent int    `json:"old_event"`
	OldStage int    `json:"old_stage"`
	NewEvent primitive.ObjectID `json:"new_event"`
	NewStage int    `json:"new_stage"`
}

func (h *handler) Import(w http.ResponseWriter, r *http.Request) {
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
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Error{time.Now(), err.Error()})
		return
	}

	errs, _ := errgroup.WithContext(context.TODO())
	for _, linkage := range linkages {
		errs.Go(func() error {
			return h.singleImport(linkage)
		})
	}
	
	if err = errs.Wait(); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Error{time.Now(), err.Error()})
		return
	}
	
	w.WriteHeader(http.StatusOK)
}

func (h *handler) singleImport(linkage *EventLinkage) error {
	event, err := h.Octane.FindEvent(&linkage.NewEvent)
	if err != nil {
		return err
	}
	

	// Reset Matches
	_, err = h.Octane.DeleteMatch(bson.M{"event._id": linkage.NewEvent, "stage._id": linkage.NewStage})
	if err != nil {
		return err
	}

	// Reset Games
	_, err = h.Octane.DeleteGame(bson.M{"match.event._id": linkage.NewEvent, "match.stage._id": linkage.NewStage})
	if err != nil {
		return err
	}

	matches, err := h.getMatches(linkage, event)
	if err != nil {
		return err
	}

	for _, match := range matches {
		_, err := h.Octane.InsertMatch(match)
		if err != nil {
			return err
		}

		games, err := h.getGames(match)
		if err != nil {
			return err
		}

		_, err = h.Octane.InsertGames(games)
		if err != nil {
			return err
		}
	}

	return nil
}

func (d *deprecated) getLinkages(events []int) ([]*EventLinkage, error) {
	stmt := linkagesSQL
	if len(events) > 0 {
		stmt += fmt.Sprintf(" WHERE old_event IN (%s)", strings.Join(intsToStrings(events), ","))
	}

	results, err := d.DB.Query(stmt)
	if err != nil {
		return nil, err
	}

	var linkages []*EventLinkage
	for results.Next() {
		var linkage EventLinkage
		var newEvent string
		err = results.Scan(&linkage.OldEvent, &linkage.OldStage, &newEvent, &linkage.NewStage)
		if err != nil {
			return nil, err
		}

		if linkage.NewEvent, err = primitive.ObjectIDFromHex(newEvent); err != nil {
			return nil, err
		}

		linkages = append(linkages, &linkage)
	}

	return linkages, nil
}

func (h *handler) getMatches(linkage *EventLinkage, event *octane.Event) ([]*octane.Match, error) {
	oldMatches, err := h.Deprecated.GetMatches(&GetMatchesContext{
		Event: strconv.Itoa(linkage.OldEvent), 
		Stage: strconv.Itoa(linkage.OldStage),
	})
	if err != nil {
		return nil, err
	}

	var newMatches []*octane.Match
	for _, match := range oldMatches {
		id := primitive.NewObjectID()
		newMatch := &octane.Match{
			ID: &id,
			OctaneID: match.OctaneID,
			Date:     match.Date,
			Format:   match.Format,
			Number:   match.Number,
			Blue: &octane.MatchSide{
				Score:  match.Blue.Score,
				Winner: match.Blue.Winner,
			},
			Orange: &octane.MatchSide{
				Score:  match.Orange.Score,
				Winner: match.Orange.Winner,
			},
			Event: &octane.Event{
				ID:     event.ID,
				Name:   event.Name,
				Mode:   event.Mode,
				Region: event.Region,
				Tier:   event.Tier,
			},
			Stage:  &octane.Stage{
				ID:     linkage.NewStage,
				Name:   event.Stages[linkage.NewStage].Name,
				Format: event.Stages[linkage.NewStage].Format,
			},
		}

		if event.Stages[linkage.NewStage].Qualifier {
			newMatch.Stage.Qualifier = true
		}

		if match.Blue.Name != "" && match.Orange.Name != "" {
			newMatch.Blue.Team = h.findOrInsertTeam(match.Blue.Name)
			newMatch.Orange.Team = h.findOrInsertTeam(match.Orange.Name)
		}

		newMatches = append(newMatches, newMatch)
	}

	return newMatches, nil
}

func (h *handler) getGames(match *octane.Match) ([]interface{}, error) {
	oldGames, err := h.Deprecated.GetGames(&GetGamesContext{
		OctaneID: match.OctaneID,
		Blue: match.Blue.Team.Name,
		Orange: match.Orange.Team.Name,
	})
	if err != nil {
		return nil, err
	}

	var newGames []interface{}
	for _, game := range oldGames {
		if game.Blue.Name == match.Orange.Team.Name {
			game.Blue, game.Orange = game.Orange, game.Blue
		}

		id := primitive.NewObjectID()
		newGame := &octane.Game{
			ID: &id,
			OctaneID: game.OctaneID,
			Number:   game.Number,
			Map:      game.Map,
			Duration: game.Duration,
			Blue: &octane.GameSide{
				Team: match.Blue.Team,
				Goals:   game.Blue.Goals,
				Players: h.toPlayers(game.Blue.Players),
			},
			Orange: &octane.GameSide{
				Team: match.Orange.Team,
				Goals:   game.Orange.Goals,
				Players: h.toPlayers(game.Orange.Players),
			},
			Match: &octane.Match{
				ID:     match.ID,
				Format: match.Format,
				Event:  match.Event,
				Stage:  match.Stage,
			},
			Date: match.Date,
		}

		newGame.Blue.Winner = newGame.Blue.Goals > newGame.Orange.Goals
		newGame.Orange.Winner = newGame.Orange.Goals > newGame.Blue.Goals

		newGames = append(newGames, newGame)
	}

	return newGames, nil
}

func (h *handler) toPlayers(logs []Log) []*octane.PlayerStats {
	var players []*octane.PlayerStats
	for _, log := range logs {
		players = append(players, &octane.PlayerStats{
			Player: h.findOrInsertPlayer(log.Player),
			Stats: &ballchasing.PlayerStats{
				Core: &ballchasing.PlayerCore{
					Score:              log.Score,
					Goals:              log.Goals,
					Assists:            log.Assists,
					Saves:              log.Saves,
					Shots:              log.Shots,
					Mvp:                log.MVP,
					ShootingPercentage: log.SP,
					GoalParticipation:  log.GP,
					Rating:             log.Rating,
				},
			},
		})
	}

	return players
}

func intsToStrings(a []int) []string {
	b := make([]string, len(a))
	for i, n := range a {
		b[i] = strconv.Itoa(n)
	}
	return b
}

func (h *handler) findOrInsertTeam(name string) *octane.Team {
	teams, err := h.Octane.FindTeams(bson.M{"name": name}, nil, nil)
	if err != nil || len(teams.Teams) == 0 {
		team, _ := h.Octane.InsertTeam(&octane.Team{
			Name: name,
		})
		return &octane.Team{
			ID:   team,
			Name: name,
		}
	}

	return &octane.Team{
		ID:   teams.Teams[0].ID,
		Name: teams.Teams[0].Name,
	}
}

func (h *handler) findOrInsertPlayer(tag string) *octane.Player {
	players, err := h.Octane.FindPlayers(bson.M{"tag": tag}, nil, nil)
	if err != nil || len(players.Players) == 0 {
		player, _ := h.Octane.InsertPlayer(&octane.Player{
			Tag: tag,
		})
		return &octane.Player{
			ID:  player,
			Tag: tag,
		}
	}

	return &octane.Player{
		ID:  players.Players[0].ID,
		Tag: players.Players[0].Tag,
	}
}