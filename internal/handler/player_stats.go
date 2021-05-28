package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

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
