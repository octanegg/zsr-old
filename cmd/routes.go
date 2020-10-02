package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/octanegg/core/internal/deprecated"
	"github.com/octanegg/core/internal/handler"
	"github.com/rs/cors"
	"github.com/urfave/negroni"
)

func routes(h handler.Handler, d deprecated.Handler) http.Handler {
	r := mux.NewRouter()

	// health
	r.HandleFunc("/health", h.Health).
		Methods(http.MethodGet)

	// events
	e := r.PathPrefix("/events").Subrouter()
	e.HandleFunc("", h.GetEvents).
		Methods(http.MethodGet)

	// matches
	m := r.PathPrefix("/matches").Subrouter()
	m.HandleFunc("", h.GetMatches).
		Methods(http.MethodGet)

	// games
	g := r.PathPrefix("/games").Subrouter()
	g.HandleFunc("", h.GetGames).
		Methods(http.MethodGet)

	// players
	p := r.PathPrefix("/players").Subrouter()
	p.HandleFunc("", h.GetPlayers).
		Methods(http.MethodGet)
	// teams
	t := r.PathPrefix("/teams").Subrouter()
	t.HandleFunc("", h.GetTeams).
		Methods(http.MethodGet)

	// TODO: Stats endpoints

	// TODO: News endpoints

	// admin
	r.HandleFunc("/import", d.ImportMatches).Methods(http.MethodPost)
	s := mux.NewRouter().PathPrefix("/deprecated").Subrouter()
	s.HandleFunc("/matches", d.UpdateMatch).Methods(http.MethodPost)
	s.HandleFunc("/matches/{id}", d.GetMatch).Methods(http.MethodGet)
	s.HandleFunc("/matches/{event}/{stage}", d.GetMatches).Methods(http.MethodGet)
	s.HandleFunc("/games", d.InsertGame).Methods(http.MethodPut)
	s.HandleFunc("/games", d.DeleteGame).Methods(http.MethodDelete)
	s.HandleFunc("/games/{match}/{blue}/{orange}", d.GetGames).Methods(http.MethodGet)
	r.PathPrefix("/deprecated").Handler(negroni.New(
		negroni.HandlerFunc(jwtMiddleware),
		negroni.HandlerFunc(scopeMiddleware),
		negroni.Wrap(s),
	))

	return cors.AllowAll().Handler(r)
}
