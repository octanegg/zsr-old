package handler

import (
	"encoding/json"
	"net/http"
	"net/url"
	"time"

	"github.com/octanegg/zsr/octane/filter"
	"go.mongodb.org/mongo-driver/bson"
)

func (h *handler) GetPlayersStats(w http.ResponseWriter, r *http.Request) {
	data, err := h.Stats.GetPlayerAggregate("player", statsFilter(r.URL.Query()))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Error{time.Now(), err.Error()})
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
		filter.ObjectIDs("player._id", v["player"]),
		filter.ObjectIDs("team._id", v["team"]),
		filter.ObjectIDs("opponent._id", v["opponent"]),
		filter.Dates("game.date", v.Get("before"), v.Get("after")),
		filter.Bool("winner", v.Get("winner")),
	)
}
