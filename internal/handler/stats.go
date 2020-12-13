package handler

import (
	"encoding/json"
	"net/http"
	"net/url"
	"time"

	"github.com/octanegg/zsr/octane/filter"
	"github.com/octanegg/zsr/octane/pipelines"
	"go.mongodb.org/mongo-driver/bson"
)

func (h *handler) GetPlayersStats(w http.ResponseWriter, r *http.Request) {
	v := r.URL.Query()

	var having bson.M
	if len(v) == 1 && len(v["mode"]) > 0 {
		having = bson.M{"games": bson.M{"$gt": 50}}
	}

	pipeline := pipelines.PlayerAggregate(statsFilter(v), having)
	data, err := h.Octane.Statlines().Pipeline(pipeline.Pipeline, pipeline.Decode)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Error{time.Now(), err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}

func statsFilter(v url.Values) bson.M {
	return filter.New(
		filter.Strings("game.match.event.tier", v["tier"]),
		filter.Strings("game.match.event.region", v["region"]),
		filter.Ints("game.match.event.mode", v["mode"]),
		filter.ObjectIDs("game.match.event._id", v["event"]),
		filter.Ints("game.match.stage._id", v["stage"]),
		filter.ObjectIDs("player._id", v["player"]),
		filter.ObjectIDs("team._id", v["team"]),
		filter.ObjectIDs("opponent._id", v["opponent"]),
		filter.Dates("game.date", v.Get("before"), v.Get("after")),
		filter.Bool("winner", v.Get("winner")),
	)
}
