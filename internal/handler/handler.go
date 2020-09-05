package handler

import (
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/octanegg/core/internal/config"
	"github.com/octanegg/core/octane"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Error .
type Error struct {
	Timestamp time.Time `json:"timestamp"`
	Error     string    `json:"error"`
}

type handler struct {
	Client octane.Client
}

// FindContext .
type FindContext struct {
	Do         func(bson.M, *octane.Pagination, *octane.Sort) (*octane.Data, error)
	Filter     bson.M
	Pagination *octane.Pagination
	Sort       *octane.Sort
}

// Handler .
type Handler interface {
	Health(http.ResponseWriter, *http.Request)

	GetEvent(http.ResponseWriter, *http.Request)
	GetMatch(http.ResponseWriter, *http.Request)
	GetGame(http.ResponseWriter, *http.Request)
	GetPlayer(http.ResponseWriter, *http.Request)
	GetTeam(http.ResponseWriter, *http.Request)

	GetEvents(http.ResponseWriter, *http.Request)
	GetMatches(http.ResponseWriter, *http.Request)
	GetGames(http.ResponseWriter, *http.Request)
	GetPlayers(http.ResponseWriter, *http.Request)
	GetTeams(http.ResponseWriter, *http.Request)

	PutEvent(http.ResponseWriter, *http.Request)
	PutMatch(http.ResponseWriter, *http.Request)
	PutGame(http.ResponseWriter, *http.Request)
	PutPlayer(http.ResponseWriter, *http.Request)
	PutTeam(http.ResponseWriter, *http.Request)

	UpdateEvent(http.ResponseWriter, *http.Request)
	UpdateMatch(http.ResponseWriter, *http.Request)
	UpdateGame(http.ResponseWriter, *http.Request)
	UpdatePlayer(http.ResponseWriter, *http.Request)
	UpdateTeam(http.ResponseWriter, *http.Request)

	DeleteEvent(http.ResponseWriter, *http.Request)
	DeleteMatch(http.ResponseWriter, *http.Request)
	DeleteGame(http.ResponseWriter, *http.Request)
	DeletePlayer(http.ResponseWriter, *http.Request)
	DeleteTeam(http.ResponseWriter, *http.Request)
}

// NewHandler .
func NewHandler(client octane.Client) Handler {
	return &handler{
		Client: client,
	}
}

func (h *handler) Get(w http.ResponseWriter, r *http.Request, ctx *FindContext) {
	data, err := ctx.Do(ctx.Filter, ctx.Pagination, ctx.Sort)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Error{time.Now(), err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}

func (h *handler) GetID(w http.ResponseWriter, r *http.Request, do func(bson.M, *octane.Pagination, *octane.Sort) (*octane.Data, error)) {
	oid, err := primitive.ObjectIDFromHex(mux.Vars(r)[config.ParamID])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Error{time.Now(), err.Error()})
		return
	}

	data, err := do(bson.M{config.KeyID: oid}, nil, nil)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Error{time.Now(), err.Error()})
		return
	}

	if len(data.Data) == 0 {
		w.WriteHeader(http.StatusNotFound)
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(data.Data[0])
	}
}

func (h *handler) Put(w http.ResponseWriter, r *http.Request, do func(io.ReadCloser) (*primitive.ObjectID, error)) {
	defer r.Body.Close()

	if r.Header.Get(config.HeaderContentType) != config.HeaderApplicationJSON {
		w.WriteHeader(http.StatusUnsupportedMediaType)
		json.NewEncoder(w).Encode(Error{time.Now(), config.ErrInvalidContentType})
		return
	}

	id, err := do(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Error{time.Now(), err.Error()})
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(id)
}

func (h *handler) Update(w http.ResponseWriter, r *http.Request, do func(*primitive.ObjectID, io.ReadCloser) (*primitive.ObjectID, error)) {
	if r.Header.Get(config.HeaderContentType) != config.HeaderApplicationJSON {
		w.WriteHeader(http.StatusUnsupportedMediaType)
		json.NewEncoder(w).Encode(Error{time.Now(), config.ErrInvalidContentType})
		return
	}

	oid, err := primitive.ObjectIDFromHex(mux.Vars(r)[config.ParamID])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Error{time.Now(), err.Error()})
		return
	}

	id, err := do(&oid, r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Error{time.Now(), err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(id)
}

func (h *handler) Delete(w http.ResponseWriter, r *http.Request, do func(*primitive.ObjectID) (int64, error)) {
	oid, err := primitive.ObjectIDFromHex(mux.Vars(r)[config.ParamID])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Error{time.Now(), err.Error()})
		return
	}

	amount, err := do(&oid)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Error{time.Now(), err.Error()})
		return
	}

	if amount == 0 {
		w.WriteHeader(http.StatusNotModified)
	} else {
		w.WriteHeader(http.StatusNoContent)
	}
}
