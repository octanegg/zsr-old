package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/octanegg/zsr/octane/pipelines"
)

func (h *handler) GetPlayerRecords(w http.ResponseWriter, r *http.Request) {
	v := r.URL.Query()

	var p *pipelines.Pipeline
	if v.Get("type") == "game" {
		p = pipelines.PlayerGameRecords(statlinesFilter(v), v.Get("stat"))
	} else {
		p = pipelines.PlayerSeriesRecords(statlinesFilter(v), v.Get("stat"))
	}

	data, err := h.Octane.Statlines().Pipeline(p.Pipeline, p.Decode)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Error{time.Now(), err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(struct {
		Records []interface{} `json:"records"`
	}{data})
}

func (h *handler) GetTeamRecords(w http.ResponseWriter, r *http.Request) {
	v := r.URL.Query()

	var p *pipelines.Pipeline
	if v.Get("type") == "game" {
		p = pipelines.TeamGameRecords(statlinesFilter(v), v.Get("stat"))
	} else {
		p = pipelines.TeamSeriesRecords(statlinesFilter(v), v.Get("stat"))
	}

	data, err := h.Octane.Statlines().Pipeline(p.Pipeline, p.Decode)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Error{time.Now(), err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(struct {
		Records []interface{} `json:"records"`
	}{data})
}

func (h *handler) GetGameRecords(w http.ResponseWriter, r *http.Request) {
	v := r.URL.Query()

	var (
		data []interface{}
		err  error
	)

	if v.Get("stat") == "duration" {
		p := pipelines.GameDurationRecords(gamesFilter(v))
		data, err = h.Octane.Games().Pipeline(p.Pipeline, p.Decode)
	} else {
		p := pipelines.GameRecords(gamesFilter(v), v.Get("stat"))
		data, err = h.Octane.Games().Pipeline(p.Pipeline, p.Decode)
	}

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Error{time.Now(), err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(struct {
		Records []interface{} `json:"records"`
	}{data})
}

func (h *handler) GetSeriesRecords(w http.ResponseWriter, r *http.Request) {
	v := r.URL.Query()

	var (
		data []interface{}
		err  error
	)

	if v.Get("stat") == "duration" {
		p := pipelines.SeriesDurationRecords(gamesFilter(v))
		data, err = h.Octane.Games().Pipeline(p.Pipeline, p.Decode)
	} else {
		p := pipelines.SeriesRecords(gamesFilter(v), v.Get("stat"))
		data, err = h.Octane.Games().Pipeline(p.Pipeline, p.Decode)
	}

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Error{time.Now(), err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(struct {
		Records []interface{} `json:"records"`
	}{data})
}
