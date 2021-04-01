package handler

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/octanegg/zsr/octane"
	"github.com/octanegg/zsr/octane/collection"
	"github.com/octanegg/zsr/octane/filter"
	"github.com/octanegg/zsr/octane/stats"
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
		filter.ObjectIDs("event._id", []string{mux.Vars(r)["_id"]}),
	)

	matches, err := h.Octane.Matches().Find(filter, nil, nil)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Error{time.Now(), err.Error()})
		return
	}

	stages := []int{}
	for _, stage := range r.URL.Query()["stage"] {
		i, _ := strconv.Atoi(stage)
		stages = append(stages, i)
	}

	participants := stats.GetEventParticipants(matches, stages)
	if participants == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(struct {
		Participants []*octane.Participant `json:"participants"`
	}{participants})
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
