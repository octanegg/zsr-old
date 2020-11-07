package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"github.com/octanegg/zsr/octane"
	"github.com/octanegg/zsr/octane/filter"
	"go.mongodb.org/mongo-driver/bson"
)

func (h *handler) GetRecords(w http.ResponseWriter, r *http.Request) {
	stat := mux.Vars(r)["stat"]
	if !octane.IsValidStat(strings.ToLower(stat)) {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Error{time.Now(), fmt.Sprintf("valid stats: %s", octane.ValidStats())})
		return
	}

	var (
		v = r.URL.Query()
		p = pagination(v)
		s = recordsSort(stat)
		f = recordsFilter(v)
	)

	data, err := h.Octane.Stats().Find(f, s, p)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Error{time.Now(), err.Error()})
		return
	}

	if len(data) > 50 {
		data = data[:50]
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}

func recordsSort(stat string) bson.M {
	switch stat {
	case "score":
		return bson.M{"stats.core.score": -1}
	case "goals":
		return bson.M{"stats.core.goals": -1}
	case "assists":
		return bson.M{"stats.core.assists": -1}
	case "saves":
		return bson.M{"stats.core.saves": -1}
	case "shots":
		return bson.M{"stats.core.shots": -1}
	case "rating":
		return bson.M{"stats.core.rating": -1}
	default:
		return nil
	}
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
