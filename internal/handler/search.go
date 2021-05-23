package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/octanegg/zsr/octane"
	"go.mongodb.org/mongo-driver/bson"
)

func (h *handler) Search(w http.ResponseWriter, r *http.Request) {
	filter := bson.M{}

	v := r.URL.Query()
	if v.Get("relevant") != "" {
		filter["relevant"] = true
	}

	events, err := h.Octane.Events().Find(bson.M{}, nil, nil)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Error{time.Now(), err.Error()})
		return
	}

	players, err := h.Octane.Players().Find(filter, nil, nil)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Error{time.Now(), err.Error()})
		return
	}

	teams, err := h.Octane.Teams().Find(filter, nil, nil)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Error{time.Now(), err.Error()})
		return
	}

	type SearchItem struct {
		Type   string   `json:"type,omitempty"`
		ID     string   `json:"id,omitempty"`
		Label  string   `json:"label,omitempty"`
		Groups []string `json:"groups,omitempty"`
		Image  string   `json:"image,omitempty"`
	}

	searchItems := []*SearchItem{}
	for _, e := range events {
		event := e.(octane.Event)
		searchItems = append(searchItems, &SearchItem{
			Type:   "event",
			ID:     event.Slug,
			Label:  event.Name,
			Groups: event.Groups,
			Image:  event.Image,
		})
	}

	for _, p := range players {
		player := p.(octane.Player)
		searchItems = append(searchItems, &SearchItem{
			Type:  "player",
			ID:    player.Slug,
			Label: player.Tag,
			Image: player.Country,
		})
	}

	for _, t := range teams {
		team := t.(octane.Team)
		searchItems = append(searchItems, &SearchItem{
			Type:  "team",
			ID:    team.Slug,
			Label: team.Name,
			Image: team.Image,
		})
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(struct {
		SearchList []*SearchItem `json:"searchList"`
	}{searchItems})
}
