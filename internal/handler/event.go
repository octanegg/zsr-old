package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (h *handler) GetEvents(w http.ResponseWriter, r *http.Request) {
	var (
		events interface{}
		err    error
	)

	if vars := mux.Vars(r); len(vars) > 0 {
		events, err = h.Client.FindEventByID(vars["id"])
	} else {
		events, err = h.Client.FindEvents(nil)
	}

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Error{time.Now(), err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(events)
}

func (h *handler) GetEventMatches(w http.ResponseWriter, r *http.Request) {
	oid, err := primitive.ObjectIDFromHex(mux.Vars(r)["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Error{time.Now(), err.Error()})
	}

	filter := bson.M{
		"event": oid,
	}

	matches, err := h.Client.FindMatches(filter)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Error{time.Now(), err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(matches)
}
