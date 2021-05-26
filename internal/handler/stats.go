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

	f := statlinesFilter(v)

	pipeline := pipelines.PlayerStats(f, bson.M{"player": "$player.player._id"}, having, v["stat"])
	data, err := h.Octane.Statlines().Pipeline(pipeline.Pipeline, pipeline.Decode)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Error{time.Now(), err.Error()})
		return
	}

	resp := struct {
		Stats []interface{} `json:"stats"`
	}{data}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

func (h *handler) GetPlayerTeamStats(w http.ResponseWriter, r *http.Request) {
	v := r.URL.Query()

	var having bson.M
	if minGames, err := strconv.Atoi(v.Get("minGames")); err == nil {
		having = bson.M{"games": bson.M{"$gt": minGames}}
	}

	f := statlinesFilter(v)

	pipeline := pipelines.PlayerStats(f, bson.M{"player": "$player.player._id", "team": "$team.team._id"}, having, v["stat"])
	data, err := h.Octane.Statlines().Pipeline(pipeline.Pipeline, pipeline.Decode)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Error{time.Now(), err.Error()})
		return
	}

	resp := struct {
		Stats []interface{} `json:"stats"`
	}{data}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

func (h *handler) GetPlayerOpponentStats(w http.ResponseWriter, r *http.Request) {
	v := r.URL.Query()

	var having bson.M
	if minGames, err := strconv.Atoi(v.Get("minGames")); err == nil {
		having = bson.M{"games": bson.M{"$gt": minGames}}
	}

	f := statlinesFilter(v)

	pipeline := pipelines.PlayerStats(f, bson.M{"player": "$player.player._id", "opponent": "$opponent.team._id"}, having, v["stat"])
	data, err := h.Octane.Statlines().Pipeline(pipeline.Pipeline, pipeline.Decode)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Error{time.Now(), err.Error()})
		return
	}

	resp := struct {
		Stats []interface{} `json:"stats"`
	}{data}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

func (h *handler) GetPlayerEventStats(w http.ResponseWriter, r *http.Request) {
	v := r.URL.Query()

	var having bson.M
	if minGames, err := strconv.Atoi(v.Get("minGames")); err == nil {
		having = bson.M{"games": bson.M{"$gt": minGames}}
	}

	f := statlinesFilter(v)

	pipeline := pipelines.PlayerStats(f, bson.M{"player": "$player.player._id", "event": "$game.match.event._id"}, having, v["stat"])
	data, err := h.Octane.Statlines().Pipeline(pipeline.Pipeline, pipeline.Decode)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Error{time.Now(), err.Error()})
		return
	}

	resp := struct {
		Stats []interface{} `json:"stats"`
	}{data}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

func (h *handler) GetTeamStats(w http.ResponseWriter, r *http.Request) {
	v := r.URL.Query()

	var having bson.M
	if minGames, err := strconv.Atoi(v.Get("minGames")); err == nil {
		having = bson.M{"games": bson.M{"$gt": minGames}}
	}

	f := statlinesFilter(v)

	pipeline := pipelines.TeamStats(f, bson.M{"team": "$team.team._id"}, having, v["stat"])
	data, err := h.Octane.Statlines().Pipeline(pipeline.Pipeline, pipeline.Decode)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Error{time.Now(), err.Error()})
		return
	}

	resp := struct {
		Stats []interface{} `json:"stats"`
	}{data}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

func (h *handler) GetTeamOpponentStats(w http.ResponseWriter, r *http.Request) {
	v := r.URL.Query()

	var having bson.M
	if minGames, err := strconv.Atoi(v.Get("minGames")); err == nil {
		having = bson.M{"games": bson.M{"$gt": minGames}}
	}

	f := statlinesFilter(v)

	pipeline := pipelines.TeamStats(f, bson.M{"team": "$team.team._id", "opponent": "$opponent.team._id"}, having, v["stat"])
	data, err := h.Octane.Statlines().Pipeline(pipeline.Pipeline, pipeline.Decode)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Error{time.Now(), err.Error()})
		return
	}

	resp := struct {
		Stats []interface{} `json:"stats"`
	}{data}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

func (h *handler) GetTeamEventStats(w http.ResponseWriter, r *http.Request) {
	v := r.URL.Query()

	var having bson.M
	if minGames, err := strconv.Atoi(v.Get("minGames")); err == nil {
		having = bson.M{"games": bson.M{"$gt": minGames}}
	}

	f := statlinesFilter(v)

	pipeline := pipelines.TeamStats(f, bson.M{"team": "$team.team._id", "event": "$game.match.event._id"}, having, v["stat"])
	data, err := h.Octane.Statlines().Pipeline(pipeline.Pipeline, pipeline.Decode)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Error{time.Now(), err.Error()})
		return
	}

	resp := struct {
		Stats []interface{} `json:"stats"`
	}{data}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

func statlinesFilter(v url.Values) bson.M {
	return filter.New(
		filter.Strings("game.match.event.tier", v["tier"]),
		filter.Strings("game.match.event.region", v["region"]),
		filter.Ints("game.match.event.mode", v["mode"]),
		filter.ExplicitAnd(
			filter.Or(
				filter.ObjectIDs("game.match.event._id", v["event"]),
				filter.Strings("game.match.event.slug", v["event"]),
			),
			filter.Or(
				filter.ObjectIDs("player.player._id", v["player"]),
				filter.Strings("player.player.slug", v["player"]),
			),
			filter.Or(
				filter.ObjectIDs("team.team._id", v["team"]),
				filter.Strings("team.team.slug", v["team"]),
			),
			filter.Or(
				filter.ObjectIDs("opponent.team._id", v["opponent"]),
				filter.Strings("opponent.team.slug", v["opponent"]),
			),
		),
		filter.Ints("game.match.stage._id", v["stage"]),
		filter.Strings("player.player.country", v["nationality"]),
		filter.ObjectIDs("opponent.team._id", v["opponent"]),
		filter.Dates("game.date", v.Get("before"), v.Get("after")),
		filter.Bool("team.winner", v.Get("winner")),
		filter.Ints("game.match.format.length", v["bestOf"]),
		filter.Bool("game.match.stage.qualifier", v.Get("qualifier")),
		filter.Strings("game.match.event.groups", v["group"]),
	)
}
