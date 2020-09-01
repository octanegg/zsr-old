package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func (h *handler) GetMatches(w http.ResponseWriter, r *http.Request) {
	var (
		matches interface{}
		err     error
	)

	if vars := mux.Vars(r); len(vars) > 0 {
		matches, err = h.Client.FindMatchByID(vars["id"])
	} else {
		matches, err = h.Client.FindMatches(nil)
	}

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Error{time.Now(), err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(matches)
}
