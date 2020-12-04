package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"github.com/octanegg/zsr/octane/filter"
	"github.com/octanegg/zsr/octane/stats"
	"go.mongodb.org/mongo-driver/bson"
)

func (h *handler) GetGameRecords(w http.ResponseWriter, r *http.Request) {
	stat := mux.Vars(r)["stat"]
	if !stats.IsValidStat(strings.ToLower(stat)) {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Error{time.Now(), fmt.Sprintf("valid stats: %s", stats.ValidStats())})
		return
	}

	sort := bson.M{fmt.Sprintf("stats.core.%s", stat): -1}
	data, err := h.Stats.GetGameRecords(recordsFilter(r.URL.Query()), sort)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Error{time.Now(), err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}

func recordsFilter(v url.Values) bson.M {
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
