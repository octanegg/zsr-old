package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
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

func (h *handler) GetGames(w http.ResponseWriter, r *http.Request) {
	var (
		v = r.URL.Query()
		p = pagination(v)
		s = sort(v)
		f = gamesFilter(v)
	)

	data, err := h.Octane.Games().Find(f, s, p)
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
		Games []interface{} `json:"games"`
		*collection.Pagination
	}{data, p})
}

func (h *handler) GetGame(w http.ResponseWriter, r *http.Request) {
	id, err := primitive.ObjectIDFromHex(mux.Vars(r)["_id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Error{time.Now(), err.Error()})
		return
	}

	data, err := h.Octane.Games().FindOne(bson.M{"_id": id})
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

func (h *handler) DeleteGame(w http.ResponseWriter, r *http.Request) {
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

	g, err := h.Octane.Games().FindOne(bson.M{"_id": id})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Error{time.Now(), err.Error()})
		return
	}
	game := g.(octane.Game)

	if _, err := h.Octane.Games().Delete(bson.M{"_id": id}); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Error{time.Now(), err.Error()})
		return
	}

	if _, err := h.Octane.Statlines().Delete(bson.M{"game._id": id}); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Error{time.Now(), err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)

	helper.UpdateMatchAggregate(h.Octane, game.Match.ID)
}

func (h *handler) CreateGame(w http.ResponseWriter, r *http.Request) {
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

	var game *octane.Game = &octane.Game{}
	if err := json.Unmarshal(body, game); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Error{time.Now(), err.Error()})
		return
	}

	if game.BallchasingID != "" {
		game, err = helper.UseBallchasing(h.Octane, game)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(Error{time.Now(), err.Error()})
			return
		}
	} else {
		for _, player := range append(game.Blue.Players, game.Orange.Players...) {
			if player.Stats.Core.Score == 0 {
				player.Stats.Core.Score, err = helper.AverageScore(h.Octane, player.Stats.Core)
				if err != nil {
					w.WriteHeader(http.StatusBadRequest)
					json.NewEncoder(w).Encode(Error{time.Now(), err.Error()})
					return
				}
			}
		}

		game.Blue.Team.Stats = helper.PlayerStatsToTeamStats(game.Blue.Players)
		game.Orange.Team.Stats = helper.PlayerStatsToTeamStats(game.Orange.Players)
		game.Blue.Winner = game.Blue.Team.Stats.Core.Goals > game.Orange.Team.Stats.Core.Goals
		game.Orange.Winner = game.Orange.Team.Stats.Core.Goals > game.Blue.Team.Stats.Core.Goals
		game.Overtime = game.Duration > 300

		for _, player := range game.Blue.Players {
			player.Advanced = &octane.AdvancedStats{}
			if game.Blue.Team.Stats.Core.Goals > 0 {
				player.Advanced.GoalParticipation = float64(player.Stats.Core.Goals+player.Stats.Core.Assists) / float64(game.Blue.Team.Stats.Core.Goals) * 100
			}

			player.Advanced.Rating = helper.Rating(player)
		}

		for _, player := range game.Orange.Players {
			player.Advanced = &octane.AdvancedStats{}
			if game.Orange.Team.Stats.Core.Goals > 0 {
				player.Advanced.GoalParticipation = float64(player.Stats.Core.Goals+player.Stats.Core.Assists) / float64(game.Orange.Team.Stats.Core.Goals) * 100
			}

			player.Advanced.Rating = helper.Rating(player)
		}
	}

	id, err := h.Octane.Games().InsertOne(game)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Error{time.Now(), err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(struct {
		ID string `json:"_id"`
	}{id.Hex()})

	helper.UpdateGame(h.Octane, id)
	helper.UpdateMatchAggregate(h.Octane, game.Match.ID)
}

func (h *handler) UpdateGame(w http.ResponseWriter, r *http.Request) {
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

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Error{time.Now(), err.Error()})
		return
	}

	var game *octane.Game = &octane.Game{}
	if err := json.Unmarshal(body, game); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Error{time.Now(), err.Error()})
		return
	}

	g, err := h.Octane.Games().FindOne(bson.M{"_id": id})
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Error{time.Now(), err.Error()})
		return
	}
	existingGame := g.(octane.Game)

	if game.BallchasingID != existingGame.BallchasingID {
		game, err = helper.UseBallchasing(h.Octane, game)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(Error{time.Now(), err.Error()})
			return
		}
	} else {
		for _, player := range append(game.Blue.Players, game.Orange.Players...) {
			if player.Stats.Core.Score == 0 {
				player.Stats.Core.Score, err = helper.AverageScore(h.Octane, player.Stats.Core)
				if err != nil {
					w.WriteHeader(http.StatusBadRequest)
					json.NewEncoder(w).Encode(Error{time.Now(), err.Error()})
					return
				}
			}
		}

		game.Blue.Team.Stats = helper.PlayerStatsToTeamStats(game.Blue.Players)
		game.Orange.Team.Stats = helper.PlayerStatsToTeamStats(game.Orange.Players)
		game.Blue.Winner = game.Blue.Team.Stats.Core.Goals > game.Orange.Team.Stats.Core.Goals
		game.Orange.Winner = game.Orange.Team.Stats.Core.Goals > game.Blue.Team.Stats.Core.Goals
		game.Overtime = game.Duration > 300

		for _, player := range game.Blue.Players {
			player.Advanced = &octane.AdvancedStats{}
			if game.Blue.Team.Stats.Core.Goals > 0 {
				player.Advanced.GoalParticipation = float64(player.Stats.Core.Goals+player.Stats.Core.Assists) / float64(game.Blue.Team.Stats.Core.Goals) * 100
			}

			player.Advanced.Rating = helper.Rating(player)
		}

		for _, player := range game.Orange.Players {
			player.Advanced = &octane.AdvancedStats{}
			if game.Orange.Team.Stats.Core.Goals > 0 {
				player.Advanced.GoalParticipation = float64(player.Stats.Core.Goals+player.Stats.Core.Assists) / float64(game.Orange.Team.Stats.Core.Goals) * 100
			}

			player.Advanced.Rating = helper.Rating(player)
		}
	}

	if _, err := h.Octane.Games().UpdateOne(bson.M{"_id": id}, bson.M{"$set": game}); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Error{time.Now(), err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)

	helper.UpdateGame(h.Octane, &id)
	helper.UpdateMatchAggregate(h.Octane, game.Match.ID)
}

func gamesFilter(v url.Values) bson.M {
	return filter.New(
		filter.Strings("match.event.tier", v["tier"]),
		filter.Strings("match.event.region", v["region"]),
		filter.Ints("match.event.mode", v["mode"]),
		filter.Ints("match.stage._id", v["stage"]),
		filter.Ints("match.substage", v["substage"]),
		filter.ObjectIDs("match._id", v["match"]),
		filter.Dates("date", v.Get("before"), v.Get("after")),
		filter.Ints("match.format.length", v["bestOf"]),
		filter.Strings("match.event.groups", v["group"]),
		filter.Bool("match.stage.qualifier", v.Get("qualifier")),
		filter.Bool("match.stage.lan", v.Get("lan")),
		filter.Bool("overtime", v.Get("overtime")),
		filter.ExplicitAnd(
			filter.Or(
				filter.ObjectIDs("match.event._id", v["event"]),
				filter.Strings("match.event.slug", v["event"]),
			),
			filter.Or(
				filter.ElemMatch("blue.players", filter.ObjectIDs("player._id", v["player"])),
				filter.ElemMatch("orange.players", filter.ObjectIDs("player._id", v["player"])),
			),
			filter.Or(
				filter.And(filter.ElemMatch("blue.players", filter.ObjectIDs("player._id", v["player"])), filter.ObjectIDs("orange.team.team._id", v["opponent"])),
				filter.And(filter.ElemMatch("orange.players", filter.ObjectIDs("player._id", v["player"])), filter.ObjectIDs("blue.team.team._id", v["opponent"])),
			),
			filter.Or(
				filter.ObjectIDs("blue.team.team._id", v["team"]),
				filter.ObjectIDs("orange.team.team._id", v["team"]),
			),
			filter.Or(
				filter.And(filter.ObjectIDs("blue.team.team._id", v["team"]), filter.ObjectIDs("orange.team.team._id", v["opponent"])),
				filter.And(filter.ObjectIDs("orange.team.team._id", v["team"]), filter.ObjectIDs("blue.team.team._id", v["opponent"])),
			),
		),
	)
}
