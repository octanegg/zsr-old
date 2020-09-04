package handler

import (
	"net/http"
)

func (h *handler) GetPlayers(w http.ResponseWriter, r *http.Request) {
	h.Get(w, r, h.Client.FindPlayers)
}

func (h *handler) GetPlayer(w http.ResponseWriter, r *http.Request) {
	h.GetID(w, r, h.Client.FindPlayer)
}

func (h *handler) PutPlayer(w http.ResponseWriter, r *http.Request) {
	h.Put(w, r, h.Client.InsertPlayer)
}

func (h *handler) UpdatePlayer(w http.ResponseWriter, r *http.Request) {
	h.Update(w, r, h.Client.UpdatePlayer)
}

func (h *handler) DeletePlayer(w http.ResponseWriter, r *http.Request) {
	h.Delete(w, r, h.Client.DeletePlayer)
}
