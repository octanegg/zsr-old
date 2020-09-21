package handler

import (
	"net/http"
)

func (h *handler) GetGames(w http.ResponseWriter, r *http.Request) {
	h.Get(w, r, h.Octane.FindGames)
}

func (h *handler) GetGame(w http.ResponseWriter, r *http.Request) {
	h.GetID(w, r, h.Octane.FindGames)
}

func (h *handler) PutGame(w http.ResponseWriter, r *http.Request) {
	h.Put(w, r, h.Octane.InsertGameWithReader)
}

func (h *handler) UpdateGame(w http.ResponseWriter, r *http.Request) {
	h.Update(w, r, h.Octane.UpdateGameWithReader)
}

func (h *handler) DeleteGame(w http.ResponseWriter, r *http.Request) {
	h.Delete(w, r, h.Octane.DeleteGame)

}
