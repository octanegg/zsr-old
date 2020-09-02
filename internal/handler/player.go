package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (h *handler) GetPlayers(w http.ResponseWriter, r *http.Request) {
	var (
		players interface{}
		err     error
	)

	if vars := mux.Vars(r); len(vars) > 0 {
		oid, err := primitive.ObjectIDFromHex(vars["id"])
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(Error{time.Now(), err.Error()})
			return
		}

		players, err = h.Client.FindPlayer(&oid)
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
