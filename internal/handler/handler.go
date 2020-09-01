package handler

import (
	"net/http"
	"time"

	"github.com/octanegg/core/octane"
)

type handler struct {
	Client octane.Client
}

type Handler interface {
	GetEvent(http.ResponseWriter, *http.Request)
	GetEvents(http.ResponseWriter, *http.Request)
}

// NewHandler .
func NewHandler(client octane.Client) Handler {
	return &handler{
		Client: client,
	}
}

type Error struct {
	Timestamp time.Time `json:"timestamp"`
	Error     string    `json:"error"`
}
