package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/octanegg/zsr/internal/handler"
	"github.com/rs/cors"
)

func routes(h handler.Handler) http.Handler {
	r := mux.NewRouter()

	r.Handle("/", http.FileServer(http.Dir("./docs")))

	r.HandleFunc("/health", h.Health).
		Methods(http.MethodGet)

	r.HandleFunc("/search", h.Search).
		Methods(http.MethodGet)

	e := r.PathPrefix("/events").Subrouter()
	e.HandleFunc("", h.GetEvents).
		Methods(http.MethodGet)
	e.HandleFunc("/{_id}", h.GetEvent).
		Methods(http.MethodGet)
	e.HandleFunc("/{_id}/matches", h.GetEventMatches).
		Methods(http.MethodGet)
	e.HandleFunc("/{_id}/participants", h.GetEventParticipants).
		Methods(http.MethodGet)

	m := r.PathPrefix("/matches").Subrouter()
	m.HandleFunc("", h.GetMatches).
		Methods(http.MethodGet)
	m.HandleFunc("/{_id}", h.GetMatch).
		Methods(http.MethodGet)
	m.HandleFunc("/{_id}/games", h.GetMatchGames).
		Methods(http.MethodGet)
	m.HandleFunc("/{_id}/games/{number}", h.GetMatchGame).
		Methods(http.MethodGet)

	g := r.PathPrefix("/games").Subrouter()
	g.HandleFunc("", h.GetGames).
		Methods(http.MethodGet)
	g.HandleFunc("/{_id}", h.GetGame).
		Methods(http.MethodGet)

	p := r.PathPrefix("/players").Subrouter()
	p.HandleFunc("", h.GetPlayers).
		Methods(http.MethodGet)
	p.HandleFunc("/{_id}", h.GetPlayer).
		Methods(http.MethodGet)
	p.HandleFunc("/{_id}/teams", h.GetPlayerTeams).
		Methods(http.MethodGet)
	p.HandleFunc("/{_id}/opponents", h.GetPlayerOpponents).
		Methods(http.MethodGet)

	t := r.PathPrefix("/teams").Subrouter()
	t.HandleFunc("", h.GetTeams).
		Methods(http.MethodGet)
	t.HandleFunc("/active", h.GetActiveTeams).
		Methods(http.MethodGet)
	t.HandleFunc("/{_id}", h.GetTeam).
		Methods(http.MethodGet)

	c := r.PathPrefix("/records").Subrouter()
	c.HandleFunc("/players", h.GetPlayerRecords).
		Methods(http.MethodGet)
	c.HandleFunc("/teams", h.GetTeamRecords).
		Methods(http.MethodGet)
	c.HandleFunc("/games", h.GetGameRecords).
		Methods(http.MethodGet)
	c.HandleFunc("/series", h.GetSeriesRecords).
		Methods(http.MethodGet)

	s := r.PathPrefix("/stats").Subrouter()
	s.HandleFunc("/players", h.GetPlayerStats).
		Methods(http.MethodGet)
	s.HandleFunc("/players/teams", h.GetPlayerTeamStats).
		Methods(http.MethodGet)
	s.HandleFunc("/players/opponents", h.GetPlayerOpponentStats).
		Methods(http.MethodGet)
	s.HandleFunc("/players/events", h.GetPlayerEventStats).
		Methods(http.MethodGet)

	s.HandleFunc("/teams", h.GetTeamStats).
		Methods(http.MethodGet)
	s.HandleFunc("/teams/opponents", h.GetTeamOpponentStats).
		Methods(http.MethodGet)
	s.HandleFunc("/teams/events", h.GetTeamEventStats).
		Methods(http.MethodGet)

	e.HandleFunc("", h.CreateEvent).
		Methods(http.MethodPost)
	e.HandleFunc("/{_id}", h.UpdateEvent).
		Methods(http.MethodPut)
	e.HandleFunc("/{_id}", h.DeleteEvent).
		Methods(http.MethodDelete)

	m.HandleFunc("", h.CreateMatch).
		Methods(http.MethodPost)
	m.HandleFunc("/{_id}", h.UpdateMatch).
		Methods(http.MethodPut)
	m.HandleFunc("/{_id}", h.DeleteMatch).
		Methods(http.MethodDelete)
	m.HandleFunc("", h.UpdateMatches).
		Methods(http.MethodPut)

	g.HandleFunc("", h.CreateGame).
		Methods(http.MethodPost)
	g.HandleFunc("/{_id}", h.UpdateGame).
		Methods(http.MethodPut)
	g.HandleFunc("/{_id}", h.DeleteGame).
		Methods(http.MethodDelete)

	p.HandleFunc("", h.CreatePlayer).
		Methods(http.MethodPost)
	p.HandleFunc("/{_id}", h.UpdatePlayer).
		Methods(http.MethodPut)
	p.HandleFunc("/{_id}/merge", h.MergePlayers).
		Methods(http.MethodPost)

	t.HandleFunc("", h.CreateTeam).
		Methods(http.MethodPost)
	t.HandleFunc("/{_id}", h.UpdateTeam).
		Methods(http.MethodPut)

	return cors.AllowAll().Handler(r)
}
