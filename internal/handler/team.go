package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (h *handler) GetTeams(w http.ResponseWriter, r *http.Request) {
	page, perPage := getPaginationDetails(r.URL.Query())
	// TODO: Team filters
	teams, err := h.Client.FindTeams(nil, page, perPage)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Error{time.Now(), err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(teams)
}

func (h *handler) GetTeam(w http.ResponseWriter, r *http.Request) {
	oid, err := primitive.ObjectIDFromHex(mux.Vars(r)["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Error{time.Now(), err.Error()})
		return
	}

	team, err := h.Client.FindTeam(&oid)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Error{time.Now(), err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(team)
}
