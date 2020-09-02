package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/octanegg/core/internal/handler"
)

func routes(h handler.Handler) *mux.Router {
	r := mux.NewRouter()

	// events
	r.HandleFunc("/events", h.GetEvents).
		Methods(http.MethodGet)
	r.HandleFunc("/events", h.PutEvent).
		Methods(http.MethodPut)
	r.HandleFunc("/events/{id}", h.GetEvents).
		Methods(http.MethodGet)
	r.HandleFunc("/events/{id}", h.UpdateEvent).
		Methods(http.MethodPost)
	r.HandleFunc("/events/{id}", h.DeleteEvent).
		Methods(http.MethodDelete)
	r.HandleFunc("/events/{id}/matches", h.GetEventMatches).
		Methods(http.MethodGet)

	// matches
	r.HandleFunc("/matches", h.GetMatches).
		Methods(http.MethodGet)
	r.HandleFunc("/matches", h.PutMatch).
		Methods(http.MethodPut)
	r.HandleFunc("/matches/{id}", h.GetMatches).
		Methods(http.MethodGet)
	r.HandleFunc("/matches/{id}", h.UpdateMatch).
		Methods(http.MethodPost)
	r.HandleFunc("/matches/{id}", h.DeleteMatch).
		Methods(http.MethodDelete)
	r.HandleFunc("/matches/{id}/games", h.GetMatchGames).
		Methods(http.MethodGet)

	// games
	r.HandleFunc("/games", h.GetGames).
		Methods(http.MethodGet)
	r.HandleFunc("/games", h.PutGame).
		Methods(http.MethodPut)
	r.HandleFunc("/games/{id}", h.GetGames).
		Methods(http.MethodGet)
	r.HandleFunc("/games/{id}", h.UpdateGame).
		Methods(http.MethodPost)
	r.HandleFunc("/games/{id}", h.DeleteGame).
		Methods(http.MethodDelete)

	// players
	r.HandleFunc("/players", h.GetPlayers).
		Methods(http.MethodGet)
	r.HandleFunc("/players/{id}", h.GetPlayers).
		Methods(http.MethodGet)

	// teams
	r.HandleFunc("/teams", h.GetTeams).
		Methods(http.MethodGet)
	r.HandleFunc("/teams/{id}", h.GetTeams).
		Methods(http.MethodGet)

	return r
}