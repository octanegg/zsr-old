package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"regexp"
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

func (h *handler) GetTeams(w http.ResponseWriter, r *http.Request) {
	var (
		v = r.URL.Query()
		p = pagination(v)
		s = sort(v)
		f = teamsFilter(v)
	)

	data, err := h.Octane.Teams().Find(f, s, p)
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
		Teams []interface{} `json:"teams"`
		*collection.Pagination
	}{data, p})
}

func (h *handler) GetTeam(w http.ResponseWriter, r *http.Request) {
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

	data, err := h.Octane.Teams().FindOne(filter)
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

func (h *handler) CreateTeam(w http.ResponseWriter, r *http.Request) {
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

	var team octane.Team
	if err := json.Unmarshal(body, &team); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Error{time.Now(), err.Error()})
		return

	}

	id := primitive.NewObjectID()
	team.ID = &id
	team.Slug = helper.TeamSlug(&team)

	if _, err := h.Octane.Teams().InsertOne(team); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Error{time.Now(), err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(struct {
		ID string `json:"_id"`
	}{id.Hex()})
}

func (h *handler) UpdateTeam(w http.ResponseWriter, r *http.Request) {
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

	var team octane.Team
	if err := json.Unmarshal(body, &team); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Error{time.Now(), err.Error()})
		return

	}

	team.Slug = helper.TeamSlug(&team)
	id := team.ID
	team.ID = nil

	if _, err := h.Octane.Teams().UpdateOne(bson.M{"_id": id}, bson.M{"$set": team}); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Error{time.Now(), err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)

	helper.UpdateTeam(h.Octane, id, id)
}

func (h *handler) GetActiveTeams(w http.ResponseWriter, r *http.Request) {
	players, err := h.Octane.Players().Find(bson.M{"team": bson.M{"$exists": true}}, nil, nil)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Error{time.Now(), err.Error()})
		return
	}

	participants := stats.ActiveTeams(players, r.URL.Query()["region"])
	if participants == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(struct {
		Teams []*octane.Participant `json:"teams"`
	}{participants})
}

func teamsFilter(v url.Values) bson.M {
	return filter.New(
		filter.FuzzyStrings("name", v["name"]),
		filter.Bool("relevant", v.Get("relevant")),
	)
}
