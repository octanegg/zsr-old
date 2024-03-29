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
		*collection.Pagination
	}{data, p})
}

func (h *handler) GetPlayer(w http.ResponseWriter, r *http.Request) {
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

	data, err := h.Octane.Players().FindOne(filter)
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

func (h *handler) CreatePlayer(w http.ResponseWriter, r *http.Request) {
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

	var player octane.Player
	if err := json.Unmarshal(body, &player); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Error{time.Now(), err.Error()})
		return

	}

	id := primitive.NewObjectID()
	player.ID = &id
	player.Slug = helper.PlayerSlug(&player)

	if _, err := h.Octane.Players().InsertOne(player); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Error{time.Now(), err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(struct {
		ID string `json:"_id"`
	}{id.Hex()})
}

func (h *handler) UpdatePlayer(w http.ResponseWriter, r *http.Request) {
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

	var player octane.Player
	if err := json.Unmarshal(body, &player); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Error{time.Now(), err.Error()})
		return

	}

	player.Slug = helper.PlayerSlug(&player)
	id := player.ID
	player.ID = nil

	update := bson.M{"$set": player}
	unset := bson.M{}

	if player.Name == "" {
		unset["name"] = ""
	}

	if player.Accounts == nil || len(player.Accounts) == 0 {
		unset["accounts"] = ""
	}

	if player.Team == nil {
		unset["team"] = ""
	}

	if !player.Substitute {
		unset["substitute"] = ""
	}

	if !player.Coach {
		unset["coach"] = ""
	}

	if !player.Relevant {
		unset["relevant"] = ""
	}

	if len(unset) > 0 {
		update["$unset"] = unset
	}

	if _, err := h.Octane.Players().UpdateOne(bson.M{"_id": id}, update); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Error{time.Now(), err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)

	helper.UpdatePlayer(h.Octane, id, id)
}

func (h *handler) MergePlayers(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get(config.HeaderApiKey) != os.Getenv(config.EnvApiKey) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	newID, err := primitive.ObjectIDFromHex(mux.Vars(r)["_id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Error{time.Now(), err.Error()})
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Error{time.Now(), err.Error()})
		return
	}

	var merge struct {
		ID *primitive.ObjectID `json:"_id" bson:"_id"`
	}
	if err := json.Unmarshal(body, &merge); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Error{time.Now(), err.Error()})
		return

	}

	if err := helper.UpdatePlayer(h.Octane, merge.ID, &newID); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Error{time.Now(), err.Error()})
		return
	}

	if _, err := h.Octane.Players().Delete(bson.M{"_id": merge.ID}); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Error{time.Now(), err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
}

func playersFilter(v url.Values) bson.M {
	return filter.New(
		filter.Strings("country", v["country"]),
		filter.FuzzyStrings("tag", v["tag"]),
		filter.Or(
			filter.ObjectIDs("team._id", v["team"]),
			filter.Strings("team.slug", v["team"]),
		),
		filter.Bool("relevant", v.Get("relevant")),
	)
}
