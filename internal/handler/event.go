package handler

import (
	"encoding/json"
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

func (h *handler) GetEvents(w http.ResponseWriter, r *http.Request) {
	events, err := h.Client.FindEvents(
		buildEventFilter(r.URL.Query()),
		getPagination(r.URL.Query()),
		getSort(r.URL.Query()),
	)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Error{time.Now(), err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(events)
}

func (h *handler) GetEvent(w http.ResponseWriter, r *http.Request) {
	oid, err := primitive.ObjectIDFromHex(mux.Vars(r)[config.ParamID])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Error{time.Now(), err.Error()})
		return
	}

	event, err := h.Client.FindEvent(&oid)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Error{time.Now(), err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(event)
}

func (h *handler) PutEvent(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get(config.HeaderContentType) != config.HeaderApplicationJSON {
		w.WriteHeader(http.StatusUnsupportedMediaType)
		json.NewEncoder(w).Encode(Error{time.Now(), config.ErrInvalidContentType})
		return
	}

	var event octane.Event
	if err := json.NewDecoder(r.Body).Decode(&event); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Error{time.Now(), err.Error()})
		return
	}

	id, err := h.Client.InsertEvent(&event)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Error{time.Now(), err.Error()})
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(id)
}

func (h *handler) UpdateEvent(w http.ResponseWriter, r *http.Request) {
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

	var event octane.Event
	if err := json.NewDecoder(r.Body).Decode(&event); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Error{time.Now(), err.Error()})
		return
	}

	id, err := h.Client.UpdateEvent(&oid, &event)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Error{time.Now(), err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(id)
}

func (h *handler) DeleteEvent(w http.ResponseWriter, r *http.Request) {
	oid, err := primitive.ObjectIDFromHex(mux.Vars(r)[config.ParamID])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Error{time.Now(), err.Error()})
		return
	}

	amount, err := h.Client.DeleteEvent(&oid)
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

func buildEventFilter(v url.Values) bson.M {
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

	return filter
}
