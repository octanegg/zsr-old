package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/octanegg/zsr/octane/filter"
)

// MAX_SEARCH_RESULTS .
const MAX_SEARCH_RESULTS = 50

func (h *handler) Search(w http.ResponseWriter, r *http.Request) {
	type searchResults struct {
		Events  []interface{} `json:"events"`
		Players []interface{} `json:"players"`
		Teams   []interface{} `json:"teams"`
	}

	search := r.URL.Query()["input"]

	events, err := h.Octane.Events().Find(filter.New(filter.Or(filter.FuzzyStrings("name", search), filter.FuzzyStrings("groups", search))), nil, nil)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Error{time.Now(), err.Error()})
		return
	}

	players, err := h.Octane.Players().Find(filter.New(filter.FuzzyStrings("tag", search)), nil, nil)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Error{time.Now(), err.Error()})
		return
	}

	teams, err := h.Octane.Teams().Find(filter.New(filter.FuzzyStrings("name", search)), nil, nil)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Error{time.Now(), err.Error()})
		return
	}

	results := searchResults{
		Events:  []interface{}{},
		Players: []interface{}{},
		Teams:   []interface{}{},
	}
	if len(events)+len(players)+len(teams) <= MAX_SEARCH_RESULTS {
		results.Events = events
		results.Players = players
		results.Teams = teams
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(struct {
		Results interface{} `json:"results"`
	}{results})
}
