package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
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
