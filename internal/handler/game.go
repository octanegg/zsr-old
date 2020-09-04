package handler

import (
	"net/http"
)

func (h *handler) GetGames(w http.ResponseWriter, r *http.Request) {
	h.Get(w, r, h.Client.FindGames)
}

func (h *handler) GetGame(w http.ResponseWriter, r *http.Request) {
	h.GetID(w, r, h.Client.FindGame)
}

func (h *handler) PutGame(w http.ResponseWriter, r *http.Request) {
	h.Put(w, r, h.Client.InsertGame)
}

func (h *handler) UpdateGame(w http.ResponseWriter, r *http.Request) {
	h.Update(w, r, h.Client.UpdateGame)
}

func (h *handler) DeleteGame(w http.ResponseWriter, r *http.Request) {
	h.Delete(w, r, h.Client.DeleteGame)

}
