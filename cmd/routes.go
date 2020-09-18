package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/octanegg/core/internal/admin"
	"github.com/octanegg/core/internal/handler"
	"github.com/rs/cors"
)

func routes(h handler.Handler, a admin.Handler) http.Handler {
	r := mux.NewRouter()

	// health
	r.HandleFunc("/health", h.Health).
		Methods(http.MethodGet)

	// events
	e := r.PathPrefix("/events").Subrouter()
	e.HandleFunc("", h.GetEvents).
		Methods(http.MethodGet)
	e.HandleFunc("", h.PutEvent).
		Methods(http.MethodPut)
	e.HandleFunc("/{id}", h.GetEvent).
		Methods(http.MethodGet)
	e.HandleFunc("/{id}", h.UpdateEvent).
		Methods(http.MethodPost)
	e.HandleFunc("/{id}", h.DeleteEvent).
		Methods(http.MethodDelete)

	// matches
	m := r.PathPrefix("/matches").Subrouter()
	m.HandleFunc("", h.GetMatches).
		Methods(http.MethodGet)
	m.HandleFunc("", h.PutMatch).
		Methods(http.MethodPut)
	m.HandleFunc("/{id}", h.GetMatch).
		Methods(http.MethodGet)
	m.HandleFunc("/{id}", h.UpdateMatch).
		Methods(http.MethodPost)
	m.HandleFunc("/{id}", h.DeleteMatch).
		Methods(http.MethodDelete)

	// games
	g := r.PathPrefix("/games").Subrouter()
	g.HandleFunc("", h.GetGames).
		Methods(http.MethodGet)
	g.HandleFunc("", h.PutGame).
		Methods(http.MethodPut)
	g.HandleFunc("/{id}", h.GetGame).
		Methods(http.MethodGet)
	g.HandleFunc("/{id}", h.UpdateGame).
		Methods(http.MethodPost)
	g.HandleFunc("/{id}", h.DeleteGame).
		Methods(http.MethodDelete)

	// players
	p := r.PathPrefix("/players").Subrouter()
	p.HandleFunc("", h.GetPlayers).
		Methods(http.MethodGet)
	p.HandleFunc("", h.PutPlayer).
		Methods(http.MethodPut)
	p.HandleFunc("/{id}", h.GetPlayer).
		Methods(http.MethodGet)
	p.HandleFunc("/{id}", h.UpdatePlayer).
		Methods(http.MethodPost)
	p.HandleFunc("/{id}", h.DeletePlayer).
		Methods(http.MethodDelete)

	// teams
	t := r.PathPrefix("/teams").Subrouter()
	t.HandleFunc("", h.GetTeams).
		Methods(http.MethodGet)
	t.HandleFunc("", h.PutTeam).
		Methods(http.MethodPut)
	t.HandleFunc("/{id}", h.GetTeam).
		Methods(http.MethodGet)
	t.HandleFunc("/{id}", h.UpdateTeam).
		Methods(http.MethodPost)
	t.HandleFunc("/{id}", h.DeleteTeam).
		Methods(http.MethodDelete)

	// TODO: Stats endpoints

	// TODO: News endpoints

	// TODO: Authentication endpoints

	// admin
	s := r.PathPrefix("/admin").Subrouter()
	s.HandleFunc("/link-ballchasing", a.LinkBallchasing).Methods(http.MethodPost)
	s.HandleFunc("/import-matches", a.ImportMatches).Methods(http.MethodPost)
	s.HandleFunc("/update-match", a.UpdateMatch).Methods(http.MethodPost)
	s.HandleFunc("/get-match/{id}", a.GetMatch).Methods(http.MethodGet)
	s.HandleFunc("/reset-game", a.ResetGame).Methods(http.MethodPost)

	return cors.Default().Handler(r)
}
