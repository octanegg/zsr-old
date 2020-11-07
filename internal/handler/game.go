package handler

import (
	"encoding/json"
	"net/http"
	"net/url"
	"time"

	"github.com/gorilla/mux"
	"github.com/octanegg/zsr/octane"
	"github.com/octanegg/zsr/octane/filter"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (h *handler) GetGames(w http.ResponseWriter, r *http.Request) {
	var (
		v = r.URL.Query()
		p = pagination(v)
		s = sort(v)
		f = gamesFilter(v)
	)

	data, err := h.Octane.Games().Find(f, s, p)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Error{time.Now(), err.Error()})
		return
	}

	if p != nil {
		p.PageSize = len(data)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(struct {
		Games []interface{} `json:"games"`
		*octane.Pagination
	}{data, p})
}

func (h *handler) GetGame(w http.ResponseWriter, r *http.Request) {
	id, err := primitive.ObjectIDFromHex(mux.Vars(r)["_id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Error{time.Now(), err.Error()})
		return
	}

	data, err := h.Octane.Games().FindOne(bson.M{"_id": id})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Error{time.Now(), err.Error()})
		return
	}

	if data == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}

func gamesFilter(v url.Values) bson.M {
	return filter.New(
		filter.ObjectIDs("match.event._id", v["event"]),
		filter.Strings("match.event.tier", v["tier"]),
		filter.Strings("match.event.region", v["region"]),
		filter.Ints("match.event.mode", v["mode"]),
		filter.Ints("match.stage._id", v["stage"]),
		filter.Ints("match.substage", v["substage"]),
		filter.ObjectIDs("match._id", v["event"]),
		filter.BeforeDate("date", v.Get("before")),
		filter.AfterDate("date", v.Get("after")),
		filter.Or(
			filter.Strings("blue.players.player._id", v["player"]),
			filter.Strings("orange.players.player._id", v["player"]),
		),
		filter.Or(
			filter.Strings("blue.team._id", v["team"]),
			filter.Strings("orange.team._id", v["team"]),
		),
	)
}
