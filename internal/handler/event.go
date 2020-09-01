package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func (h *handler) GetEvents(w http.ResponseWriter, r *http.Request) {
	events, err := h.Client.FindEvents(nil)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Error{time.Now(), err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(events)
}

func (h *handler) GetEvent(w http.ResponseWriter, r *http.Request) {
	event, err := h.Client.FindEventByID(mux.Vars(r)["id"])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Error{time.Now(), err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(event)
}
