package handler

import (
	"encoding/json"
	"net/http"
	"net/url"
	"time"

	"github.com/gorilla/mux"
	"github.com/octanegg/zsr/octane/collection"
	"github.com/octanegg/zsr/octane/filter"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (h *handler) GetEvents(w http.ResponseWriter, r *http.Request) {
	var (
		v = r.URL.Query()
		p = pagination(v)
		s = sort(v)
		f = eventsFilter(v)
	)

	data, err := h.Octane.Events().Find(f, s, p)
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
		Events []interface{} `json:"events"`
		*collection.Pagination
	}{data, p})
}

func (h *handler) GetEvent(w http.ResponseWriter, r *http.Request) {
	id, err := primitive.ObjectIDFromHex(mux.Vars(r)["_id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Error{time.Now(), err.Error()})
		return
	}

	data, err := h.Octane.Events().FindOne(bson.M{"_id": id})
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

func eventsFilter(v url.Values) bson.M {
	beforeField, afterField := "end_date", "start_date"
	before, after := v.Get("before"), v.Get("after")
	if date := v.Get("date"); date != "" {
		beforeField, afterField = afterField, beforeField
		before, after = date, date
	}

	return filter.New(
		filter.Strings("name", v["name"]),
		filter.Strings("tier", v["tier"]),
		filter.Strings("region", v["region"]),
		filter.Strings("mode", v["mode"]),
		filter.BeforeDate(beforeField, before),
		filter.AfterDate(afterField, after),
	)
}
