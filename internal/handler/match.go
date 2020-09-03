package handler

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/octanegg/core/octane"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (h *handler) GetMatches(w http.ResponseWriter, r *http.Request) {
	matches, err := h.Client.FindMatches(
		buildGameFilter(r.URL.Query()),
		getPagination(r.URL.Query()),
		getSort(r.URL.Query()),
	)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Error{time.Now(), err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(matches)
}

func (h *handler) GetMatch(w http.ResponseWriter, r *http.Request) {
	oid, err := primitive.ObjectIDFromHex(mux.Vars(r)["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Error{time.Now(), err.Error()})
		return
	}

	match, err := h.Client.FindMatch(&oid)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Error{time.Now(), err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(match)
}

func (h *handler) PutMatch(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get(contentType) != applicationJSON {
		w.WriteHeader(http.StatusUnsupportedMediaType)
		json.NewEncoder(w).Encode(Error{time.Now(), errContentType})
		return
	}

	var match octane.Match
	if err := json.NewDecoder(r.Body).Decode(&match); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Error{time.Now(), err.Error()})
		return
	}

	id, err := h.Client.InsertMatch(&match)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Error{time.Now(), err.Error()})
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(id)
}

func (h *handler) UpdateMatch(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get(contentType) != applicationJSON {
		w.WriteHeader(http.StatusUnsupportedMediaType)
		json.NewEncoder(w).Encode(Error{time.Now(), errContentType})
		return
	}

	oid, err := primitive.ObjectIDFromHex(mux.Vars(r)["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Error{time.Now(), err.Error()})
		return
	}

	var match octane.Match
	if err := json.NewDecoder(r.Body).Decode(&match); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Error{time.Now(), err.Error()})
		return
	}

	id, err := h.Client.UpdateMatch(&oid, &match)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Error{time.Now(), err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(id)
}

func (h *handler) DeleteMatch(w http.ResponseWriter, r *http.Request) {
	oid, err := primitive.ObjectIDFromHex(mux.Vars(r)["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Error{time.Now(), err.Error()})
		return
	}

	amount, err := h.Client.DeleteMatch(&oid)
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

func buildMatchFilter(v url.Values) bson.M {
	filter := bson.M{}
	if event := v.Get("event"); event != "" {
		if i, err := primitive.ObjectIDFromHex(event); err == nil {
			filter["event"] = i
		}
	}
	if stage := v.Get("stage"); stage != "" {
		if i, err := strconv.Atoi(stage); err == nil {
			filter["stage"] = i
		}
	}
	if substage := v.Get("substage"); substage != "" {
		if i, err := strconv.Atoi(substage); err == nil {
			filter["substage"] = i
		}
	}

	return filter
}
