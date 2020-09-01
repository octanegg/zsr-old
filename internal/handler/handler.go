package handler

import (
	"net/http"
	"time"

	"github.com/octanegg/core/octane"
)

type handler struct {
	Client octane.Client
}

// Handler .
type Handler interface {
	GetEvents(http.ResponseWriter, *http.Request)
	GetEventMatches(http.ResponseWriter, *http.Request)
	GetPlayers(http.ResponseWriter, *http.Request)
	GetTeams(http.ResponseWriter, *http.Request)
	GetMatches(http.ResponseWriter, *http.Request)
	GetMatchGames(http.ResponseWriter, *http.Request)
	GetGames(http.ResponseWriter, *http.Request)
}

// NewHandler .
func NewHandler(client octane.Client) Handler {
	return &handler{
		Client: client,
	}
}

// Error .
type Error struct {
	Timestamp time.Time `json:"timestamp"`
	Error     string    `json:"error"`
}
