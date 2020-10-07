package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/octanegg/zsr/internal/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (h *handler) GetEvents(w http.ResponseWriter, r *http.Request) {
	v := r.URL.Query()
	filter := getBasicFilters(v)

	dates := bson.M{}
	if val := v.Get(config.ParamBefore); val != "" {
		if t, err := time.Parse("2006-01-02T03:04:05Z", val); err == nil {
			dates["$lte"] = t
		}
	}
	if val := v.Get(config.ParamAfter); val != "" {
		if t, err := time.Parse("2006-01-02T03:04:05Z", val); err == nil {
			dates["$gte"] = t
		}
	}
	if len(dates) > 0 {
		filter[config.ParamStartDate] = dates
	}

	data, err := h.Octane.FindEvents(filter, getPagination(v), getSort(v))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Error{time.Now(), err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}


func (h *handler) GetEvent(w http.ResponseWriter, r *http.Request) {
	id, err := primitive.ObjectIDFromHex(mux.Vars(r)["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Error{time.Now(), err.Error()})
		return
	}

	data, err := h.Octane.FindEvents(bson.M{"_id": id}, nil, nil)
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