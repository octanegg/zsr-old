package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/octanegg/zsr/octane"
	"github.com/octanegg/zsr/octane/filter"
	"go.mongodb.org/mongo-driver/bson"
)

func (h *handler) Search(w http.ResponseWriter, r *http.Request) {
	v := r.URL.Query()
	if v.Get("searchEvents") != "" || v.Get("searchPlayers") != "" || v.Get("searchTeams") != "" || v.Get("searchOpponents") != "" {
		h.SearchAdvanced(w, r)
		return
	}

	filter := bson.M{}
	if v.Get("relevant") != "" {
		filter["relevant"] = true
	}

	events, err := h.Octane.Events().Find(bson.M{}, nil, nil)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Error{time.Now(), err.Error()})
		return
	}

	players, err := h.Octane.Players().Find(filter, nil, nil)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Error{time.Now(), err.Error()})
		return
	}

	teams, err := h.Octane.Teams().Find(filter, nil, nil)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Error{time.Now(), err.Error()})
		return
	}

	type SearchItem struct {
		Type   string   `json:"type,omitempty"`
		ID     string   `json:"id,omitempty"`
		Label  string   `json:"label,omitempty"`
		Groups []string `json:"groups,omitempty"`
		Image  string   `json:"image,omitempty"`
	}

	searchItems := []*SearchItem{}
	for _, e := range events {
		event := e.(octane.Event)
		searchItems = append(searchItems, &SearchItem{
			Type:   "event",
			ID:     event.Slug,
			Label:  event.Name,
			Groups: event.Groups,
			Image:  event.Image,
		})
	}

	for _, p := range players {
		player := p.(octane.Player)
		searchItems = append(searchItems, &SearchItem{
			Type:  "player",
			ID:    player.Slug,
			Label: player.Tag,
			Image: player.Country,
		})
	}

	for _, t := range teams {
		team := t.(octane.Team)
		searchItems = append(searchItems, &SearchItem{
			Type:  "team",
			ID:    team.Slug,
			Label: team.Name,
			Image: team.Image,
		})
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(struct {
		SearchList []*SearchItem `json:"searchList"`
	}{searchItems})
}

func (h *handler) SearchAdvanced(w http.ResponseWriter, r *http.Request) {
	var (
		events    = []*octane.Event{}
		teams     = []*octane.Team{}
		opponents = []*octane.Team{}
		players   = []*octane.Player{}
	)

	v := r.URL.Query()

	if v.Get("searchEvents") != "" {
		eventFilter := filter.New(
			filter.ExplicitAnd(
				filter.Or(
					filter.Strings("team.team.slug", v["team"]),
					filter.ObjectIDs("team.team._id", v["team"]),
				),
				filter.Or(
					filter.Strings("opponent.team.slug", v["opponent"]),
					filter.ObjectIDs("opponent.team._id", v["opponent"]),
				),
				filter.Or(
					filter.Strings("player.player.slug", v["player"]),
					filter.ObjectIDs("player.player._id", v["player"]),
				),
			),
		)

		dEvents, err := h.Octane.Statlines().Distinct("game.match.event", eventFilter)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(Error{time.Now(), err.Error()})
			return
		}

		for _, t := range dEvents {
			bytes, _ := bson.Marshal(t.(bson.D).Map())
			var event *octane.Event
			bson.Unmarshal(bytes, &event)
			events = append(events, event)
		}
	}

	if v.Get("searchTeams") != "" {
		teamFilter := filter.New(
			filter.ExplicitAnd(
				filter.Or(
					filter.Strings("game.match.event.slug", v["event"]),
					filter.ObjectIDs("game.match.event._id", v["event"]),
				),
				filter.Or(
					filter.Strings("opponent.team.slug", v["opponent"]),
					filter.ObjectIDs("opponent.team._id", v["opponent"]),
				),
				filter.Or(
					filter.Strings("player.player.slug", v["player"]),
					filter.ObjectIDs("player.player._id", v["player"]),
				),
			),
		)

		dTeams, err := h.Octane.Statlines().Distinct("team.team", teamFilter)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(Error{time.Now(), err.Error()})
			return
		}

		for _, t := range dTeams {
			bytes, _ := bson.Marshal(t.(bson.D).Map())
			var team *octane.Team
			bson.Unmarshal(bytes, &team)
			teams = append(teams, team)
		}
	}

	if v.Get("searchOpponents") != "" {
		opponentFilter := filter.New(
			filter.ExplicitAnd(
				filter.Or(
					filter.Strings("game.match.event.slug", v["event"]),
					filter.ObjectIDs("game.match.event._id", v["event"]),
				),
				filter.Or(
					filter.Strings("team.team.slug", v["team"]),
					filter.ObjectIDs("team.team._id", v["team"]),
				),
				filter.Or(
					filter.Strings("player.player.slug", v["player"]),
					filter.ObjectIDs("player.player._id", v["player"]),
				),
			),
		)

		dOpponents, err := h.Octane.Statlines().Distinct("opponent.team", opponentFilter)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(Error{time.Now(), err.Error()})
			return
		}

		for _, t := range dOpponents {
			bytes, _ := bson.Marshal(t.(bson.D).Map())
			var opponent *octane.Team
			bson.Unmarshal(bytes, &opponent)
			opponents = append(opponents, opponent)
		}
	}

	if v.Get("searchPlayers") != "" {
		playerFilter := filter.New(
			filter.ExplicitAnd(
				filter.Or(
					filter.Strings("game.match.event.slug", v["event"]),
					filter.ObjectIDs("game.match.event._id", v["event"]),
				),
				filter.Or(
					filter.Strings("team.team.slug", v["team"]),
					filter.ObjectIDs("team.team._id", v["team"]),
				),
				filter.Or(
					filter.Strings("opponent.team.slug", v["opponent"]),
					filter.ObjectIDs("opponent.team._id", v["opponent"]),
				),
			),
		)

		dPlayers, err := h.Octane.Statlines().Distinct("player.player", playerFilter)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(Error{time.Now(), err.Error()})
			return
		}

		for _, t := range dPlayers {
			bytes, _ := bson.Marshal(t.(bson.D).Map())
			var player *octane.Player
			bson.Unmarshal(bytes, &player)
			players = append(players, player)
		}
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(struct {
		Events    []*octane.Event  `json:"events"`
		Teams     []*octane.Team   `json:"teams"`
		Opponents []*octane.Team   `json:"opponents"`
		Players   []*octane.Player `json:"players"`
	}{events, teams, opponents, players})
}
