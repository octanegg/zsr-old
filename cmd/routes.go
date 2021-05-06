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

	// health
	r.HandleFunc("/health", h.Health).
		Methods(http.MethodGet)

	// events
	e := r.PathPrefix("/events").Subrouter()
	e.HandleFunc("", h.GetEvents).
		Methods(http.MethodGet)
	e.HandleFunc("", h.CreateEvent).
		Methods(http.MethodPost)
	e.HandleFunc("/{_id}", h.GetEvent).
		Methods(http.MethodGet)
	e.HandleFunc("/{_id}", h.UpdateEvent).
		Methods(http.MethodPut)
	e.HandleFunc("/{_id}/participants", h.GetEventParticipants).
		Methods(http.MethodGet)

	// matches
	m := r.PathPrefix("/matches").Subrouter()
	m.HandleFunc("", h.GetMatches).
		Methods(http.MethodGet)
	m.HandleFunc("", h.UpdateMatches).
		Methods(http.MethodPut)
	m.HandleFunc("", h.CreateMatch).
		Methods(http.MethodPost)
	m.HandleFunc("/{_id}", h.GetMatch).
		Methods(http.MethodGet)
	m.HandleFunc("/{_id}", h.UpdateMatch).
		Methods(http.MethodPut)

	// games
	g := r.PathPrefix("/games").Subrouter()
	g.HandleFunc("", h.GetGames).
		Methods(http.MethodGet)
	g.HandleFunc("", h.CreateGame).
		Methods(http.MethodPost)
	g.HandleFunc("/{_id}", h.GetGame).
		Methods(http.MethodGet)
	g.HandleFunc("/{_id}", h.UpdateGame).
		Methods(http.MethodPut)

	// players
	p := r.PathPrefix("/players").Subrouter()
	p.HandleFunc("", h.GetPlayers).
		Methods(http.MethodGet)
	p.HandleFunc("", h.CreatePlayer).
		Methods(http.MethodPost)
	p.HandleFunc("/{_id}", h.GetPlayer).
		Methods(http.MethodGet)
	p.HandleFunc("/{_id}", h.UpdatePlayer).
		Methods(http.MethodPut)
	p.HandleFunc("/{_id}/teams", h.GetPlayerTeams).
		Methods(http.MethodGet)
	p.HandleFunc("/{_id}/opponents", h.GetPlayerOpponents).
		Methods(http.MethodGet)

	// teams
	t := r.PathPrefix("/teams").Subrouter()
	t.HandleFunc("", h.GetTeams).
		Methods(http.MethodGet)
	t.HandleFunc("", h.CreateTeam).
		Methods(http.MethodPost)
	t.HandleFunc("/active", h.GetActiveTeams).
		Methods(http.MethodGet)
	t.HandleFunc("/{_id}", h.GetTeam).
		Methods(http.MethodGet)
	t.HandleFunc("/{_id}", h.UpdateTeam).
		Methods(http.MethodPut)

	// stats
	r.HandleFunc("/records/players", h.GetPlayerRecords).Methods(http.MethodGet)
	r.HandleFunc("/records/teams", h.GetTeamRecords).Methods(http.MethodGet)
	r.HandleFunc("/records/games", h.GetGameRecords).Methods(http.MethodGet)
	r.HandleFunc("/records/series", h.GetSeriesRecords).Methods(http.MethodGet)

	r.HandleFunc("/stats/players", h.GetPlayerStats).Methods(http.MethodGet)
	r.HandleFunc("/stats/players/teams", h.GetPlayerTeamStats).Methods(http.MethodGet)
	r.HandleFunc("/stats/players/opponents", h.GetPlayerOpponentStats).Methods(http.MethodGet)
	r.HandleFunc("/stats/players/events", h.GetPlayerEventStats).Methods(http.MethodGet)

	r.HandleFunc("/stats/teams", h.GetTeamStats).Methods(http.MethodGet)
	r.HandleFunc("/stats/teams/opponents", h.GetTeamOpponentStats).Methods(http.MethodGet)
	r.HandleFunc("/stats/teams/events", h.GetTeamEventStats).Methods(http.MethodGet)

	return cors.AllowAll().Handler(r)
}
