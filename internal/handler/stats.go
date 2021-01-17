package handler

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/octanegg/zsr/octane/filter"
	"github.com/octanegg/zsr/octane/pipelines"
	"go.mongodb.org/mongo-driver/bson"
)

func (h *handler) GetPlayerStats(w http.ResponseWriter, r *http.Request) {
	v := r.URL.Query()

	var having bson.M
	if minGames, err := strconv.Atoi(v.Get("minGames")); err == nil {
		having = bson.M{"games": bson.M{"$gt": minGames}}
	}

	pipeline := pipelines.PlayerAggregate(statlinesFilter(v), "$player._id", having)
	data, err := h.Octane.Statlines().Pipeline(pipeline.Pipeline, pipeline.Decode)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Error{time.Now(), err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(struct {
		Stats []interface{} `json:"stats"`
	}{data})
}

func (h *handler) GetPlayerTeamStats(w http.ResponseWriter, r *http.Request) {
	v := r.URL.Query()

	var having bson.M
	if minGames, err := strconv.Atoi(v.Get("minGames")); err == nil {
		having = bson.M{"games": bson.M{"$gt": minGames}}
	}

	pipeline := pipelines.PlayerAggregate(statlinesFilter(v), "$team.team._id", having)
	data, err := h.Octane.Statlines().Pipeline(pipeline.Pipeline, pipeline.Decode)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Error{time.Now(), err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(struct {
		Stats []interface{} `json:"stats"`
	}{data})
}

func (h *handler) GetPlayerOpponentStats(w http.ResponseWriter, r *http.Request) {
	v := r.URL.Query()

	var having bson.M
	if minGames, err := strconv.Atoi(v.Get("minGames")); err == nil {
		having = bson.M{"games": bson.M{"$gt": minGames}}
	}

	pipeline := pipelines.PlayerAggregate(statlinesFilter(v), "$opponent.team._id", having)
	data, err := h.Octane.Statlines().Pipeline(pipeline.Pipeline, pipeline.Decode)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Error{time.Now(), err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(struct {
		Stats []interface{} `json:"stats"`
	}{data})
}

func (h *handler) GetPlayerEventStats(w http.ResponseWriter, r *http.Request) {
	v := r.URL.Query()

	var having bson.M
	if minGames, err := strconv.Atoi(v.Get("minGames")); err == nil {
		having = bson.M{"games": bson.M{"$gt": minGames}}
	}

	pipeline := pipelines.PlayerAggregate(statlinesFilter(v), "$game.match.event._id", having)
	data, err := h.Octane.Statlines().Pipeline(pipeline.Pipeline, pipeline.Decode)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Error{time.Now(), err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(struct {
		Stats []interface{} `json:"stats"`
	}{data})
}

func (h *handler) GetTeamStats(w http.ResponseWriter, r *http.Request) {
	v := r.URL.Query()

	var having bson.M
	if minGames, err := strconv.Atoi(v.Get("minGames")); err == nil {
		having = bson.M{"games": bson.M{"$gt": minGames}}
	}

	pipeline := pipelines.TeamAggregate(statlinesFilter(v), "$team.team._id", having)
	data, err := h.Octane.Statlines().Pipeline(pipeline.Pipeline, pipeline.Decode)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Error{time.Now(), err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(struct {
		Stats []interface{} `json:"stats"`
	}{data})
}

func (h *handler) GetTeamOpponentStats(w http.ResponseWriter, r *http.Request) {
	v := r.URL.Query()

	var having bson.M
	if minGames, err := strconv.Atoi(v.Get("minGames")); err == nil {
		having = bson.M{"games": bson.M{"$gt": minGames}}
	}

	pipeline := pipelines.TeamAggregate(statlinesFilter(v), "$opponent.team._id", having)
	data, err := h.Octane.Statlines().Pipeline(pipeline.Pipeline, pipeline.Decode)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Error{time.Now(), err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(struct {
		Stats []interface{} `json:"stats"`
	}{data})
}

func (h *handler) GetTeamEventStats(w http.ResponseWriter, r *http.Request) {
	v := r.URL.Query()

	var having bson.M
	if minGames, err := strconv.Atoi(v.Get("minGames")); err == nil {
		having = bson.M{"games": bson.M{"$gt": minGames}}
	}

	pipeline := pipelines.TeamAggregate(statlinesFilter(v), "$game.match.event._id", having)
	data, err := h.Octane.Statlines().Pipeline(pipeline.Pipeline, pipeline.Decode)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Error{time.Now(), err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(struct {
		Stats []interface{} `json:"stats"`
	}{data})
}

func statlinesFilter(v url.Values) bson.M {
	return filter.New(
		filter.Strings("game.match.event.tier", v["tier"]),
		filter.Strings("game.match.event.region", v["region"]),
		filter.Ints("game.match.event.mode", v["mode"]),
		filter.ObjectIDs("game.match.event._id", v["event"]),
		filter.Ints("game.match.stage._id", v["stage"]),
		filter.ObjectIDs("player._id", v["player"]),
		filter.Strings("player.country", v["nationality"]),
		filter.ObjectIDs("team.team._id", v["team"]),
		filter.ObjectIDs("opponent.team._id", v["opponent"]),
		filter.Dates("game.date", v.Get("before"), v.Get("after")),
		filter.Bool("team.winner", v.Get("winner")),
		filter.Ints("game.match.format.length", v["bestOf"]),
	)
}
