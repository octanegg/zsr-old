package admin

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/octanegg/core/deprecated"
	"github.com/octanegg/core/internal/config"
	"github.com/octanegg/core/octane"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (h *handler) ImportMatches(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get(config.HeaderContentType) != config.HeaderApplicationJSON {
		w.WriteHeader(http.StatusUnsupportedMediaType)
		json.NewEncoder(w).Encode(Error{time.Now(), config.ErrInvalidContentType})
		return
	}

	var linkage deprecated.EventLinkage
	if err := json.NewDecoder(r.Body).Decode(&linkage); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Error{time.Now(), err.Error()})
		return
	}

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
	}

	for _, match := range matches {
		newMatch := h.parseMatch(match)
		newMatch.EventID = &eventID
		newMatch.Stage = linkage.NewStage
		if err = h.upsertMatch(newMatch); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(Error{time.Now(), err.Error()})
		}
	}

	w.WriteHeader(http.StatusOK)
}

func (h *handler) parseMatch(match *deprecated.Match) *octane.Match {
	newID := primitive.NewObjectID()
	newMatch := octane.Match{
		ID:       &newID,
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

func (h *handler) findOrInsertTeam(name string) *primitive.ObjectID {
	teams, err := h.Octane.FindTeams(bson.M{"name": name}, nil, nil)
	if err != nil || len(teams.Data) == 0 {
		team, _ := h.Octane.InsertTeam(&octane.Team{
			Name: name,
		})
		return team
	}
	return teams.Data[0].(octane.Team).ID
}

func (h *handler) upsertMatch(newMatch *octane.Match) error {
	data, err := h.Octane.FindMatches(bson.M{"octane_id": newMatch.OctaneID}, nil, nil)
	if err != nil {
		return err
	}

	if len(data.Data) == 0 {
		if _, err = h.Octane.InsertMatch(newMatch); err != nil {
			return fmt.Errorf("error inserting match - %s", err.Error())
		}
	} else if _, err = h.Octane.UpdateMatch(data.Data[0].(octane.Match).ID, newMatch); err != nil {
		return fmt.Errorf("error updating match - %s", err.Error())
	}

	return nil
}
