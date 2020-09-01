package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func (h *handler) GetGames(w http.ResponseWriter, r *http.Request) {
	var (
		games interface{}
		err   error
	)

	if vars := mux.Vars(r); len(vars) > 0 {
		games, err = h.Client.FindGameByID(vars["id"])
	} else {
		games, err = h.Client.FindGames(nil)
	}

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Error{time.Now(), err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(games)
}
