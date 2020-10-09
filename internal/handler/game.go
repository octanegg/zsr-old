package handler

import (
	"encoding/json"
	"net/http"
	"net/url"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (h *handler) GetGames(w http.ResponseWriter, r *http.Request) {
	v := r.URL.Query()
	data, err := h.Octane.FindGames(gameFilters(v), getPagination(v), getSort(v))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Error{time.Now(), err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}

func (h *handler) GetGame(w http.ResponseWriter, r *http.Request) {
	id, err := primitive.ObjectIDFromHex(mux.Vars(r)["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Error{time.Now(), err.Error()})
		return
	}

	data, err := h.Octane.FindGames(bson.M{"_id": id}, nil, nil)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Error{time.Now(), err.Error()})
		return
	}

	if len(data.Data) == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data.Data[0])
}

func gameFilters(v url.Values) bson.M {
	filter := bson.M{}
	if vals, ok := v["event"]; ok {
		filter["match.event._id"] = bson.M{"$in": toObjectIDs(vals)}
	}
	if vals, ok := v["tier"]; ok {
		filter["match.event.tier"] = bson.M{"$in": vals}
	}
	if vals, ok := v["region"]; ok {
		filter["match.event.region"] = bson.M{"$in": vals}
	}
	if vals, ok := v["mode"]; ok {
		filter["match.event.mode"] = bson.M{"$in": toInts(vals)}
	}
	if vals, ok := v["stage"]; ok {
		filter["match.stage._id"] = bson.M{"$in": toInts(vals)}
	}
	if vals, ok := v["substage"]; ok {
		filter["match.substage"] = bson.M{"$in": toInts(vals)}
	}
	if vals, ok := v["match"]; ok {
		filter["match._id"] = bson.M{"$in": toObjectIDs(vals)}
	}
	if t, err := time.Parse("2006-01-02T03:04:05Z", v.Get("before")); err == nil {
		filter["start_date"] = bson.M{"$lte": t}
	}
	if t, err := time.Parse("2006-01-02T03:04:05Z", v.Get("after")); err == nil {
		filter["start_date"] = bson.M{"$gte": t}
	}

	return filter
}
