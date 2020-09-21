package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/octanegg/core/internal/deprecated"
	"github.com/octanegg/core/internal/handler"
	"github.com/rs/cors"
)

func routes(h handler.Handler, d deprecated.Handler) http.Handler {
	r := mux.NewRouter()

	// health
	r.HandleFunc("/health", h.Health).
		Methods(http.MethodGet)

	// events
	e := r.PathPrefix("/events").Subrouter()
	e.HandleFunc("", h.GetEvents).
		Methods(http.MethodPost)
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
		Methods(http.MethodPost)
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
		Methods(http.MethodPost)
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
		Methods(http.MethodPost)
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
		Methods(http.MethodPost)
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
	s := r.PathPrefix("/deprecated").Subrouter()
	s.HandleFunc("/matches", d.UpdateMatch).Methods(http.MethodPost)
	s.HandleFunc("/matches/{id}", d.GetMatch).Methods(http.MethodGet)
	s.HandleFunc("/matches/{event}/{stage}", d.GetMatches).Methods(http.MethodGet)
	s.HandleFunc("/games", d.InsertGame).Methods(http.MethodPut)
	s.HandleFunc("/games", d.DeleteGame).Methods(http.MethodDelete)
	s.HandleFunc("/games/{match}/{blue}/{orange}", d.GetGames).Methods(http.MethodGet)
	s.HandleFunc("/import", d.ImportMatches).Methods(http.MethodPost)

	return cors.AllowAll().Handler(r)
}
