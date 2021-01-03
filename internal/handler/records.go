package handler

import (
	"encoding/json"
	"net/http"
	"net/url"
	"time"

	"github.com/octanegg/zsr/octane/collection"
	"github.com/octanegg/zsr/octane/filter"
	"github.com/octanegg/zsr/octane/pipelines"
	"go.mongodb.org/mongo-driver/bson"
)

func (h *handler) GetPlayerRecords(w http.ResponseWriter, r *http.Request) {
	v := r.URL.Query()

	typ, stat := v.Get("type"), v.Get("stat")
	if typ == "" || stat == "" {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Error{time.Now(), "invalid type or stat"})
		return
	}

	var p *pipelines.Pipeline
	if typ == "game" {
		p = pipelines.PlayerGameRecords(statlinesFilter(v), stat)
	} else {
		p = pipelines.PlayerSeriesRecords(statlinesFilter(v), stat)
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

	typ, stat := v.Get("type"), v.Get("stat")
	if typ == "" || stat == "" {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Error{time.Now(), "invalid type or stat"})
		return
	}

	var p *pipelines.Pipeline
	if typ == "game" {
		p = pipelines.TeamGameRecords(statlinesFilter(v), stat)
	} else {
		p = pipelines.TeamSeriesRecords(statlinesFilter(v), stat)
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

	stat := v.Get("stat")
	if stat == "" {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Error{time.Now(), "invalid stat"})
		return
	}

	var (
		data []interface{}
		err  error
	)

	if stat == "duration" {
		data, err = h.Octane.Games().Find(gamesFilter(v), bson.M{"duration": -1}, &collection.Pagination{
			Page:    1,
			PerPage: 25,
		})
	} else {
		p := pipelines.GameRecords(gamesFilter(v), stat)
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

	stat := v.Get("stat")
	if stat == "" {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Error{time.Now(), "invalid stat"})
		return
	}

	p := pipelines.SeriesRecords(gamesFilter(v), stat)
	data, err := h.Octane.Games().Pipeline(p.Pipeline, p.Decode)
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

func statlinesFilter(v url.Values) bson.M {
	return filter.New(
		filter.Strings("game.match.event.tier", v["tier"]),
		filter.Strings("game.match.event.region", v["region"]),
		filter.Ints("game.match.event.mode", v["mode"]),
		filter.ObjectIDs("player._id", v["player"]),
		filter.ObjectIDs("team._id", v["team"]),
		filter.ObjectIDs("opponent._id", v["opponent"]),
		filter.Dates("game.date", v.Get("before"), v.Get("after")),
		filter.Bool("winner", v.Get("winner")),
	)
}
