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

func (h *handler) GetPlayers(w http.ResponseWriter, r *http.Request) {
	v := r.URL.Query()
	data, err := h.Octane.FindPlayers(playerFilters(v), getPagination(v), getSort(v))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Error{time.Now(), err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}

func (h *handler) GetPlayer(w http.ResponseWriter, r *http.Request) {
	id, err := primitive.ObjectIDFromHex(mux.Vars(r)["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Error{time.Now(), err.Error()})
		return
	}

	data, err := h.Octane.FindPlayers(bson.M{"_id": id}, nil, nil)
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

func playerFilters(v url.Values) bson.M {
	filter := bson.M{}
	if vals, ok := v["country"]; ok {
		filter["country"] = bson.M{"$in": vals}
	}
	if vals, ok := v["tag"]; ok {
		filter["tag"] = bson.M{"$in": vals}
	}
	if vals, ok := v["team"]; ok {
		filter["team"] = bson.M{"$in": toObjectIDs(vals)}
	}

	return filter
}
