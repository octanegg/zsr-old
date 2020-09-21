package handler

import (
	"net/http"
)

func (h *handler) GetMatches(w http.ResponseWriter, r *http.Request) {
	h.Get(w, r, h.Octane.FindMatches)
}

func (h *handler) GetMatch(w http.ResponseWriter, r *http.Request) {
	h.GetID(w, r, h.Octane.FindMatches)
}

func (h *handler) PutMatch(w http.ResponseWriter, r *http.Request) {
	h.Put(w, r, h.Octane.InsertMatchWithReader)
}

func (h *handler) UpdateMatch(w http.ResponseWriter, r *http.Request) {
	h.Update(w, r, h.Octane.UpdateMatchWithReader)
}

func (h *handler) DeleteMatch(w http.ResponseWriter, r *http.Request) {
	h.Delete(w, r, h.Octane.DeleteMatch)
}
