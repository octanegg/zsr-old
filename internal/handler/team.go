package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func (h *handler) GetTeams(w http.ResponseWriter, r *http.Request) {
	var (
		teams interface{}
		err   error
	)

	if vars := mux.Vars(r); len(vars) > 0 {
		teams, err = h.Client.FindTeamByID(vars["id"])
	} else {
		teams, err = h.Client.FindTeams(nil)
	}

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Error{time.Now(), err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(teams)
}
