package handler

import (
	"net/http"
)

func (h *handler) GetPlayers(w http.ResponseWriter, r *http.Request) {
	h.Get(w, r, h.Octane.FindPlayers)
}

func (h *handler) GetPlayer(w http.ResponseWriter, r *http.Request) {
	h.GetID(w, r, h.Octane.FindPlayers)
}

func (h *handler) PutPlayer(w http.ResponseWriter, r *http.Request) {
	h.Put(w, r, h.Octane.InsertPlayerWithReader)
}

func (h *handler) UpdatePlayer(w http.ResponseWriter, r *http.Request) {
	h.Update(w, r, h.Octane.UpdatePlayerWithReader)
}

func (h *handler) DeletePlayer(w http.ResponseWriter, r *http.Request) {
	h.Delete(w, r, h.Octane.DeletePlayer)
}
