package admin

import (
	"net/http"
	"time"

	"github.com/octanegg/core/deprecated"
	"github.com/octanegg/core/octane"
	"github.com/octanegg/racer"
	"github.com/octanegg/slimline"
)

// Error .
type Error struct {
	Timestamp time.Time `json:"timestamp"`
	Error     string    `json:"error"`
}

type handler struct {
	Octane     octane.Client
	Racer      racer.Racer
	Slimline   slimline.Slimline
	Deprecated deprecated.Deprecated
}

// Handler .
type Handler interface {
	LinkBallchasing(http.ResponseWriter, *http.Request)
	ImportMatches(http.ResponseWriter, *http.Request)
}

// New .
func New(o octane.Client, r racer.Racer, s slimline.Slimline, d deprecated.Deprecated) Handler {
	return &handler{o, r, s, d}
}
