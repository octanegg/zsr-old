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

func (h *handler) GetMatches(w http.ResponseWriter, r *http.Request) {
	var (
		v = r.URL.Query()
		p = pagination(v)
		s = sort(v)
		f = matchesFilter(v)
	)

	data, err := h.Octane.Matches().Find(f, s, p)
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
		Matches []interface{} `json:"matches"`
		*octane.Pagination
	}{data, p})
}

func (h *handler) GetMatch(w http.ResponseWriter, r *http.Request) {
	id, err := primitive.ObjectIDFromHex(mux.Vars(r)["_id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Error{time.Now(), err.Error()})
		return
	}

	data, err := h.Octane.Matches().FindOne(bson.M{"_id": id})
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

func matchesFilter(v url.Values) bson.M {
	return filter.New(
		filter.ObjectIDs("event._id", v["event"]),
		filter.Strings("event.tier", v["tier"]),
		filter.Strings("event.region", v["region"]),
		filter.Ints("event.mode", v["mode"]),
		filter.Ints("stage._id", v["stage"]),
		filter.Ints("substage", v["substage"]),
		filter.BeforeDate("date", v.Get("before")),
		filter.AfterDate("date", v.Get("after")),
	)
}
