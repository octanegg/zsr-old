package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/octanegg/zsr/internal/config"
	"github.com/octanegg/zsr/octane"
	"github.com/octanegg/zsr/octane/collection"
	"github.com/octanegg/zsr/octane/filter"
	"github.com/octanegg/zsr/octane/helper"
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
	re := regexp.MustCompile("^[0-9a-fA-F]{24}$")

	filter := bson.M{"slug": mux.Vars(r)["_id"]}
	if re.MatchString(mux.Vars(r)["_id"]) {
		id, err := primitive.ObjectIDFromHex(mux.Vars(r)["_id"])
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(Error{time.Now(), err.Error()})
			return
		}
		filter = bson.M{"_id": id}
	}

	data, err := h.Octane.Events().FindOne(filter)
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

func (h *handler) CreateEvent(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get(config.HeaderApiKey) != os.Getenv(config.EnvApiKey) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Error{time.Now(), err.Error()})
		return
	}

	var event octane.Event
	if err := json.Unmarshal(body, &event); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Error{time.Now(), err.Error()})
		return

	}

	id := primitive.NewObjectID()
	event.ID = &id
	event.Slug = helper.EventSlug(&event)

	if _, err := h.Octane.Events().InsertOne(event); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Error{time.Now(), err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(struct {
		ID string `json:"_id"`
	}{id.Hex()})
}

func (h *handler) UpdateEvent(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get(config.HeaderApiKey) != os.Getenv(config.EnvApiKey) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Error{time.Now(), err.Error()})
		return
	}

	var event octane.Event
	if err := json.Unmarshal(body, &event); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Error{time.Now(), err.Error()})
		return

	}

	event.Slug = helper.EventSlug(&event)
	id := event.ID
	event.ID = nil

	if _, err := h.Octane.Events().UpdateOne(bson.M{"_id": id}, bson.M{"$set": event}); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Error{time.Now(), err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)

	helper.UpdateEvent(h.Octane, id, id)
}

func (h *handler) GetEventParticipants(w http.ResponseWriter, r *http.Request) {
	re := regexp.MustCompile("^[0-9a-fA-F]{24}$")

	f := bson.M{"event.slug": mux.Vars(r)["_id"]}
	if re.MatchString(mux.Vars(r)["_id"]) {
		f = filter.New(
			filter.ObjectIDs("event._id", []string{mux.Vars(r)["_id"]}),
		)
	}

	matches, err := h.Octane.Matches().Find(f, nil, nil)
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

	participants := stats.EventParticipants(matches, stages)
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
