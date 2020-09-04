package handler

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strconv"
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

// Handler .
type Handler interface {
	Health(http.ResponseWriter, *http.Request)

	GetEvent(http.ResponseWriter, *http.Request)
	GetEvents(http.ResponseWriter, *http.Request)
	GetMatch(http.ResponseWriter, *http.Request)
	GetMatches(http.ResponseWriter, *http.Request)
	GetGame(http.ResponseWriter, *http.Request)
	GetGames(http.ResponseWriter, *http.Request)
	GetPlayer(http.ResponseWriter, *http.Request)
	GetPlayers(http.ResponseWriter, *http.Request)
	GetTeam(http.ResponseWriter, *http.Request)
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

func (h *handler) Get(w http.ResponseWriter, r *http.Request, do func(bson.M, *octane.Pagination, *octane.Sort) (*octane.Data, error)) {
	v := r.URL.Query()
	data, err := do(getFilter(v), getPagination(v), getSort(v))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Error{time.Now(), err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}

func (h *handler) GetID(w http.ResponseWriter, r *http.Request, do func(*primitive.ObjectID) (interface{}, error)) {
	oid, err := primitive.ObjectIDFromHex(mux.Vars(r)[config.ParamID])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Error{time.Now(), err.Error()})
		return
	}

	data, err := do(&oid)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Error{time.Now(), err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}

func (h *handler) Put(w http.ResponseWriter, r *http.Request, do func(io.ReadCloser) (*octane.ObjectID, error)) {
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

func (h *handler) Update(w http.ResponseWriter, r *http.Request, do func(*primitive.ObjectID, io.ReadCloser) (*octane.ObjectID, error)) {
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

func getFilter(v url.Values) bson.M {
	filter := bson.M{}
	if tier := v.Get(config.ParamTier); tier != "" {
		filter[config.ParamTier] = tier
	}
	if region := v.Get(config.ParamRegion); region != "" {
		filter[config.ParamRegion] = region
	}
	if mode, err := strconv.Atoi(v.Get(config.ParamMode)); err == nil {
		filter[config.ParamMode] = mode
	}

	if event, err := primitive.ObjectIDFromHex(v.Get(config.ParamEvent)); err == nil {
		filter[config.ParamEvent] = event
	}
	if stage, err := strconv.Atoi(v.Get(config.ParamStage)); err == nil {
		filter[config.ParamStage] = stage
	}
	if substage, err := strconv.Atoi(v.Get(config.ParamSubstage)); err == nil {
		filter[config.ParamSubstage] = substage
	}

	if match, err := primitive.ObjectIDFromHex(v.Get(config.ParamMatch)); err == nil {
		filter[config.ParamMatch] = match
	}

	if country := v.Get(config.ParamCountry); country != "" {
		filter[config.ParamCountry] = country
	}
	if team := v.Get(config.ParamTeam); team != "" {
		filter[config.ParamTeam] = team
	}

	return filter
}

func getPagination(v url.Values) *octane.Pagination {
	page, _ := strconv.ParseInt(v.Get(config.ParamPage), 10, 64)
	perPage, _ := strconv.ParseInt(v.Get(config.ParamPage), 10, 64)
	if page == 0 || perPage == 0 {
		return nil
	}

	return &octane.Pagination{
		Page:    page,
		PerPage: perPage,
	}
}

func getSort(v url.Values) *octane.Sort {
	var order int
	switch v.Get(config.ParamOrder) {
	case config.ParamAscending:
		order = 1
	case config.ParamDescending:
		order = -1
	default:
		return nil
	}

	return &octane.Sort{
		Field: v.Get(config.ParamSort),
		Order: order,
	}
}
