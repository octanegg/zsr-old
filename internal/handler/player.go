package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/octanegg/core/octane"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (h *handler) GetPlayers(w http.ResponseWriter, r *http.Request) {
	page, perPage := getPaginationDetails(r.URL.Query())
	// TODO: Player filters
	players, err := h.Client.FindPlayers(nil, page, perPage)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Error{time.Now(), err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(players)
}

func (h *handler) GetPlayer(w http.ResponseWriter, r *http.Request) {
	oid, err := primitive.ObjectIDFromHex(mux.Vars(r)["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Error{time.Now(), err.Error()})
		return
	}

	player, err := h.Client.FindPlayer(&oid)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Error{time.Now(), err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(player)
}

func (h *handler) PutPlayer(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get(contentType) != applicationJSON {
		w.WriteHeader(http.StatusUnsupportedMediaType)
		json.NewEncoder(w).Encode(Error{time.Now(), errContentType})
		return
	}

	var player octane.Player
	if err := json.NewDecoder(r.Body).Decode(&player); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Error{time.Now(), err.Error()})
		return
	}

	id, err := h.Client.InsertPlayer(&player)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Error{time.Now(), err.Error()})
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(id)
}

func (h *handler) UpdatePlayer(w http.ResponseWriter, r *http.Request) {
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

	var player octane.Player
	if err := json.NewDecoder(r.Body).Decode(&player); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Error{time.Now(), err.Error()})
		return
	}

	id, err := h.Client.UpdatePlayer(&oid, &player)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Error{time.Now(), err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(id)
}

func (h *handler) DeletePlayer(w http.ResponseWriter, r *http.Request) {
	oid, err := primitive.ObjectIDFromHex(mux.Vars(r)["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Error{time.Now(), err.Error()})
		return
	}

	amount, err := h.Client.DeletePlayer(&oid)
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
