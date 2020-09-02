package handler

import (
	"net/http"
	"time"

	"github.com/octanegg/core/octane"
)

const (
	contentType     = "Content-Type"
	applicationJSON = "application/json"
	errContentType  = "Content-Type header is not application/json"
)

type handler struct {
	Client octane.Client
}

// Handler .
type Handler interface {
	Health(http.ResponseWriter, *http.Request)

	GetEvent(http.ResponseWriter, *http.Request)
	GetEvents(http.ResponseWriter, *http.Request)
	GetEventMatches(http.ResponseWriter, *http.Request)
	GetMatch(http.ResponseWriter, *http.Request)
	GetMatches(http.ResponseWriter, *http.Request)
	GetMatchGames(http.ResponseWriter, *http.Request)
	GetGame(http.ResponseWriter, *http.Request)
	GetGames(http.ResponseWriter, *http.Request)
	GetPlayer(http.ResponseWriter, *http.Request)
	GetPlayers(http.ResponseWriter, *http.Request)
	GetTeam(http.ResponseWriter, *http.Request)
	GetTeams(http.ResponseWriter, *http.Request)

	PutEvent(http.ResponseWriter, *http.Request)
	PutMatch(http.ResponseWriter, *http.Request)
	PutGame(http.ResponseWriter, *http.Request)

	UpdateEvent(http.ResponseWriter, *http.Request)
	UpdateMatch(http.ResponseWriter, *http.Request)
	UpdateGame(http.ResponseWriter, *http.Request)

	DeleteEvent(http.ResponseWriter, *http.Request)
	DeleteMatch(http.ResponseWriter, *http.Request)
	DeleteGame(http.ResponseWriter, *http.Request)
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
