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

func (h *handler) GetPlayers(w http.ResponseWriter, r *http.Request) {
	var (
		v = r.URL.Query()
		p = pagination(v)
		s = sort(v)
		f = playersFilter(v)
	)

	data, err := h.Octane.Players().Find(f, s, p)
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
		Players []interface{} `json:"players"`
		*octane.Pagination
	}{data, p})
}

func (h *handler) GetPlayer(w http.ResponseWriter, r *http.Request) {
	id, err := primitive.ObjectIDFromHex(mux.Vars(r)["_id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Error{time.Now(), err.Error()})
		return
	}

	data, err := h.Octane.Players().FindOne(bson.M{"_id": id})
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

func playersFilter(v url.Values) bson.M {
	return filter.New(
		filter.Strings("country", v["country"]),
		filter.Strings("tag", v["tag"]),
		filter.ObjectIDs("team", v["team"]),
	)
}
