package deprecated

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/octanegg/zsr/internal/config"
	"github.com/octanegg/zsr/octane"
)

// Error .
type Error struct {
	Timestamp time.Time `json:"timestamp"`
	Error     string    `json:"error"`
}

type handler struct {
	Deprecated Deprecated
	Octane     octane.Client
}

// Handler .
type Handler interface {
	UpdateMatch(http.ResponseWriter, *http.Request)
	GetMatch(http.ResponseWriter, *http.Request)
	GetMatches(http.ResponseWriter, *http.Request)
	DeleteGame(http.ResponseWriter, *http.Request)
	GetGames(http.ResponseWriter, *http.Request)
	InsertGame(http.ResponseWriter, *http.Request)
	ImportMatches(http.ResponseWriter, *http.Request)
}

// NewHandler .
func NewHandler(d Deprecated, o octane.Client) Handler {
	return &handler{d, o}
}

func (h *handler) UpdateMatch(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get(config.HeaderContentType) != config.HeaderApplicationJSON {
		w.WriteHeader(http.StatusUnsupportedMediaType)
		json.NewEncoder(w).Encode(Error{time.Now(), config.ErrInvalidContentType})
		return
	}

	var ctx []*UpdateMatchContext
	if err := json.NewDecoder(r.Body).Decode(&ctx); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Error{time.Now(), err.Error()})
		return
	}

	if err := h.Deprecated.UpdateMatches(ctx); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Error{time.Now(), err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *handler) GetMatch(w http.ResponseWriter, r *http.Request) {
	match, err := h.Deprecated.GetMatch(&GetMatchContext{
		OctaneID: mux.Vars(r)["id"],
	})

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Error{time.Now(), err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(match)
}

func (h *handler) GetMatches(w http.ResponseWriter, r *http.Request) {
	matches, err := h.Deprecated.GetMatches(&GetMatchesContext{
		Event: mux.Vars(r)["event"],
		Stage: mux.Vars(r)["stage"],
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Error{time.Now(), err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(matches)
}

func (h *handler) DeleteGame(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get(config.HeaderContentType) != config.HeaderApplicationJSON {
		w.WriteHeader(http.StatusUnsupportedMediaType)
		json.NewEncoder(w).Encode(Error{time.Now(), config.ErrInvalidContentType})
		return
	}

	var ctx DeleteGameContext
	if err := json.NewDecoder(r.Body).Decode(&ctx); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Error{time.Now(), err.Error()})
		return
	}

	if err := h.Deprecated.DeleteGame(&ctx); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Error{time.Now(), err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *handler) GetGames(w http.ResponseWriter, r *http.Request) {
	games, err := h.Deprecated.GetGames(&GetGamesContext{
		OctaneID: mux.Vars(r)["match"],
		Blue:     mux.Vars(r)["blue"],
		Orange:   mux.Vars(r)["orange"],
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Error{time.Now(), err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(games)
}

func (h *handler) InsertGame(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get(config.HeaderContentType) != config.HeaderApplicationJSON {
		w.WriteHeader(http.StatusUnsupportedMediaType)
		json.NewEncoder(w).Encode(Error{time.Now(), config.ErrInvalidContentType})
		return
	}

	var game Game
	if err := json.NewDecoder(r.Body).Decode(&game); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Error{time.Now(), err.Error()})
		return
	}

	if err := h.Deprecated.InsertGame(&game); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Error{time.Now(), err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
}
