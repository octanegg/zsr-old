package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func (h *handler) GetPlayers(w http.ResponseWriter, r *http.Request) {
	var (
		players interface{}
		err     error
	)

	if vars := mux.Vars(r); len(vars) > 0 {
		players, err = h.Client.FindPlayerByID(vars["id"])
	} else {
		players, err = h.Client.FindPlayers(nil)
	}

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Error{time.Now(), err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(players)
}
