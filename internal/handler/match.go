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
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (h *handler) GetMatches(w http.ResponseWriter, r *http.Request) {
	var (
		v = r.URL.Query()
		p = pagination(v)
		s = sort(v)
		f = matchesFilter(v)
	)

	data, err := h.Octane.Matches().Find(f, s, p)
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
		Matches []interface{} `json:"matches"`
		*collection.Pagination
	}{data, p})
}

func (h *handler) GetMatch(w http.ResponseWriter, r *http.Request) {
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

	data, err := h.Octane.Matches().FindOne(filter)
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

func (h *handler) DeleteMatch(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get(config.HeaderApiKey) != os.Getenv(config.EnvApiKey) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	id, err := primitive.ObjectIDFromHex(mux.Vars(r)["_id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Error{time.Now(), err.Error()})
		return
	}

	if _, err := h.Octane.Matches().Delete(bson.M{"_id": id}); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Error{time.Now(), err.Error()})
		return
	}

	if _, err := h.Octane.Games().Delete(bson.M{"match._id": id}); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Error{time.Now(), err.Error()})
		return
	}

	if _, err := h.Octane.Statlines().Delete(bson.M{"game.match._id": id}); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Error{time.Now(), err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *handler) CreateMatch(w http.ResponseWriter, r *http.Request) {
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

	var match octane.Match
	if err := json.Unmarshal(body, &match); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Error{time.Now(), err.Error()})
		return

	}

	id := primitive.NewObjectID()
	match.ID = &id
	match.Slug = helper.MatchSlug(&match)

	if _, err := h.Octane.Matches().InsertOne(match); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Error{time.Now(), err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(struct {
		ID string `json:"_id"`
	}{id.Hex()})
}

func (h *handler) UpdateMatch(w http.ResponseWriter, r *http.Request) {
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

	var match octane.Match
	if err := json.Unmarshal(body, &match); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Error{time.Now(), err.Error()})
		return

	}

	set := bson.M{
		"number": match.Number,
		"date":   match.Date,
	}

	unset := bson.M{}

	if match.Blue != nil {
		set["blue.score"] = match.Blue.Score
		if match.Blue.Team != nil {
			set["blue.team.team"] = match.Blue.Team.Team
		}
		if match.Orange != nil {
			set["blue.winner"] = match.Blue.Score > match.Orange.Score
		}
	} else {
		unset["blue"] = ""
	}

	if match.Orange != nil {
		set["orange.score"] = match.Orange.Score
		if match.Orange.Team != nil {
			set["orange.team.team"] = match.Orange.Team.Team
		}
		if match.Blue != nil {
			set["orange.winner"] = match.Orange.Score > match.Blue.Score
		}
	} else {
		unset["orange"] = ""
	}

	if match.Format != nil {
		set["format"] = match.Format
	} else {
		unset["format"] = ""
	}

	set["slug"] = helper.MatchSlug(&match)

	update := bson.M{"$set": set}
	if len(unset) > 0 {
		update["$unset"] = unset
	}

	if _, err := h.Octane.Matches().UpdateOne(bson.M{"_id": match.ID}, update); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Error{time.Now(), err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)

	helper.UpdateMatch(h.Octane, match.ID, match.ID)
}

func (h *handler) UpdateMatches(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get(config.HeaderApiKey) != os.Getenv(config.EnvApiKey) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var matches []octane.Match
	if err := json.Unmarshal(body, &matches); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return

	}

	w.WriteHeader(http.StatusOK)

	for _, match := range matches {
		if match.ID != nil {
			set := bson.M{
				"number": match.Number,
				"date":   match.Date,
			}

			unset := bson.M{}

			if match.Blue != nil {
				set["blue.score"] = match.Blue.Score
				if match.Blue.Team != nil {
					set["blue.team.team"] = match.Blue.Team.Team
				}
				if match.Orange != nil {
					set["blue.winner"] = match.Blue.Score > match.Orange.Score
				}
			} else {
				unset["blue"] = ""
			}

			if match.Orange != nil {
				set["orange.score"] = match.Orange.Score
				if match.Blue.Team != nil {
					set["orange.team.team"] = match.Orange.Team.Team
				}
				if match.Blue != nil {
					set["orange.winner"] = match.Orange.Score > match.Blue.Score
				}
			} else {
				unset["orange"] = ""
			}

			if match.Format != nil {
				set["format"] = match.Format
			} else {
				unset["format"] = ""
			}

			set["slug"] = helper.MatchSlug(&match)

			update := bson.M{"$set": set}

			if len(unset) > 0 {
				update["$unset"] = unset
			}

			if _, err := h.Octane.Matches().UpdateOne(bson.M{"_id": match.ID}, update); err != nil {
				continue
			}
		} else {
			id := primitive.NewObjectID()
			match.ID = &id
			match.Slug = helper.MatchSlug(&match)

			if _, err := h.Octane.Matches().InsertOne(match); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(Error{time.Now(), err.Error()})
				return
			}
		}

		helper.UpdateMatch(h.Octane, match.ID, match.ID)
	}
}

func (h *handler) GetMatchGames(w http.ResponseWriter, r *http.Request) {
	re := regexp.MustCompile("^[0-9a-fA-F]{24}$")

	filter := bson.M{"match.slug": mux.Vars(r)["_id"]}
	if re.MatchString(mux.Vars(r)["_id"]) {
		id, err := primitive.ObjectIDFromHex(mux.Vars(r)["_id"])
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(Error{time.Now(), err.Error()})
			return
		}
		filter = bson.M{"match._id": id}
	}

	data, err := h.Octane.Games().Find(filter, nil, nil)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Error{time.Now(), err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(struct {
		Games []interface{} `json:"games"`
	}{data})
}

func (h *handler) GetMatchGame(w http.ResponseWriter, r *http.Request) {
	re := regexp.MustCompile("^[0-9a-fA-F]{24}$")

	filter := bson.M{"match.slug": mux.Vars(r)["_id"]}
	if re.MatchString(mux.Vars(r)["_id"]) {
		id, err := primitive.ObjectIDFromHex(mux.Vars(r)["_id"])
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(Error{time.Now(), err.Error()})
			return
		}
		filter = bson.M{"match._id": id}
	}

	filter["number"], _ = strconv.Atoi(mux.Vars(r)["number"])

	data, err := h.Octane.Games().FindOne(filter)
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

func matchesFilter(v url.Values) bson.M {
	return filter.New(
		filter.Strings("event.tier", v["tier"]),
		filter.Strings("event.region", v["region"]),
		filter.Ints("event.mode", v["mode"]),
		filter.Ints("stage._id", v["stage"]),
		filter.Ints("substage", v["substage"]),
		filter.Strings("event.groups", v["group"]),
		filter.Dates("date", v.Get("before"), v.Get("after")),
		filter.Ints("format.length", v["bestOf"]),
		filter.Bool("reverse_sweep", v.Get("reverseSweep")),
		filter.Bool("reverse_sweep_attempt", v.Get("reverseSweepAttempt")),
		filter.Bool("stage.qualifier", v.Get("qualifier")),
		filter.Bool("stage.lan", v.Get("lan")),
		filter.If(v.Get("tbd") == "false", filter.ExplicitAnd(
			filter.Exists("blue.team"),
			filter.Exists("orange.team"),
		)),
		filter.ExplicitAnd(
			filter.Or(
				filter.ObjectIDs("event._id", v["event"]),
				filter.Strings("event.slug", v["event"]),
			),
			filter.Or(
				filter.ElemMatch("blue.players", filter.ObjectIDs("player._id", v["player"])),
				filter.ElemMatch("blue.players", filter.Strings("player.slug", v["player"])),
				filter.ElemMatch("orange.players", filter.ObjectIDs("player._id", v["player"])),
				filter.ElemMatch("orange.players", filter.Strings("player.slug", v["player"])),
			),
			filter.Or(
				filter.And(filter.ElemMatch("blue.players", filter.ObjectIDs("player._id", v["player"])), filter.ObjectIDs("orange.team.team._id", v["opponent"])),
				filter.And(filter.ElemMatch("blue.players", filter.Strings("player.slug", v["player"])), filter.Strings("orange.team.team.slug", v["opponent"])),
				filter.And(filter.ElemMatch("orange.players", filter.ObjectIDs("player._id", v["player"])), filter.ObjectIDs("blue.team.team._id", v["opponent"])),
				filter.And(filter.ElemMatch("orange.players", filter.Strings("player.slug", v["player"])), filter.Strings("blue.team.team.slug", v["opponent"])),
			),
			filter.Or(
				filter.ObjectIDs("blue.team.team._id", v["team"]),
				filter.Strings("blue.team.team.slug", v["team"]),
				filter.ObjectIDs("orange.team.team._id", v["team"]),
				filter.Strings("orange.team.team.slug", v["team"]),
			),
			filter.Or(
				filter.And(filter.ObjectIDs("blue.team.team._id", v["team"]), filter.ObjectIDs("orange.team.team._id", v["opponent"])),
				filter.And(filter.Strings("blue.team.team.slug", v["team"]), filter.Strings("orange.team.team.slug", v["opponent"])),
				filter.And(filter.ObjectIDs("orange.team.team._id", v["team"]), filter.ObjectIDs("blue.team.team._id", v["opponent"])),
				filter.And(filter.Strings("orange.team.team.slug", v["team"]), filter.Strings("blue.team.team.slug", v["opponent"])),
			),
		),
	)
}
