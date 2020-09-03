package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/octanegg/core/internal/handler"
)

func routes(h handler.Handler) *mux.Router {
	r := mux.NewRouter()

	// health
	r.HandleFunc("/health", h.Health).
		Methods(http.MethodGet)

	// events
	r.HandleFunc("/events", h.GetEvents).
		Methods(http.MethodGet)
	r.HandleFunc("/events", h.PutEvent).
		Methods(http.MethodPut)
	r.HandleFunc("/events/{id}", h.GetEvent).
		Methods(http.MethodGet)
	r.HandleFunc("/events/{id}", h.UpdateEvent).
		Methods(http.MethodPost)
	r.HandleFunc("/events/{id}", h.DeleteEvent).
		Methods(http.MethodDelete)

	// matches
	r.HandleFunc("/matches", h.GetMatches).
		Methods(http.MethodGet)
	r.HandleFunc("/matches", h.PutMatch).
		Methods(http.MethodPut)
	r.HandleFunc("/matches/{id}", h.GetMatch).
		Methods(http.MethodGet)
	r.HandleFunc("/matches/{id}", h.UpdateMatch).
		Methods(http.MethodPost)
	r.HandleFunc("/matches/{id}", h.DeleteMatch).
		Methods(http.MethodDelete)

	// games
	r.HandleFunc("/games", h.GetGames).
		Methods(http.MethodGet)
	r.HandleFunc("/games", h.PutGame).
		Methods(http.MethodPut)
	r.HandleFunc("/games/{id}", h.GetGame).
		Methods(http.MethodGet)
	r.HandleFunc("/games/{id}", h.UpdateGame).
		Methods(http.MethodPost)
	r.HandleFunc("/games/{id}", h.DeleteGame).
		Methods(http.MethodDelete)

	// players
	r.HandleFunc("/players", h.GetPlayers).
		Methods(http.MethodGet)
	r.HandleFunc("/players", h.PutPlayer).
		Methods(http.MethodPut)
	r.HandleFunc("/players/{id}", h.GetPlayer).
		Methods(http.MethodGet)
	r.HandleFunc("/players/{id}", h.UpdatePlayer).
		Methods(http.MethodPost)
	r.HandleFunc("/players/{id}", h.DeletePlayer).
		Methods(http.MethodDelete)

	// teams
	r.HandleFunc("/teams", h.GetTeams).
		Methods(http.MethodGet)
	r.HandleFunc("/teams", h.PutTeam).
		Methods(http.MethodPut)
	r.HandleFunc("/teams/{id}", h.GetTeam).
		Methods(http.MethodGet)
	r.HandleFunc("/teams/{id}", h.UpdateTeam).
		Methods(http.MethodPost)
	r.HandleFunc("/teams/{id}", h.DeleteTeam).
		Methods(http.MethodDelete)

	// TODO: Stats endpoints

	return r
}
