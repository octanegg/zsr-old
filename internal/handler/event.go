package handler

import (
	"encoding/json"
	"net/http"
	"net/url"
	"time"

	"github.com/gorilla/mux"
	"github.com/octanegg/zsr/octane/collection"
	"github.com/octanegg/zsr/octane/filter"
	"github.com/octanegg/zsr/octane/pipelines"
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

func (h *handler) GetEventParticipants(w http.ResponseWriter, r *http.Request) {
	filter := filter.New(
		filter.ObjectIDs("game.match.event._id", []string{mux.Vars(r)["_id"]}),
		filter.Ints("game.match.stage._id", r.URL.Query()["stage"]),
	)

	pipeline := pipelines.EventParticipants(filter)
	data, err := h.Octane.Statlines().Pipeline(pipeline.Pipeline, pipeline.Decode)
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
	json.NewEncoder(w).Encode(struct {
		Participants []interface{} `json:"participants"`
	}{data})
}

func eventsFilter(v url.Values) bson.M {
	return filter.New(
		filter.Strings("name", v["name"]),
		filter.Strings("tier", v["tier"]),
		filter.Strings("region", v["region"]),
		filter.Ints("mode", v["mode"]),
		filter.Strings("groups", v["group"]),
		filter.AfterDate("end_date", v.Get("date")),
		filter.BeforeDate("start_date", v.Get("date")),
		filter.AfterDate("start_date", v.Get("after")),
		filter.BeforeDate("end_date", v.Get("before")),
	)
}
