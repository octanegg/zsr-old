package handler

import (
	"net/http"
)

func (h *handler) GetMatches(w http.ResponseWriter, r *http.Request) {
	h.Get(w, r, h.Client.FindMatches)
}

func (h *handler) GetMatch(w http.ResponseWriter, r *http.Request) {
	h.GetID(w, r, h.Client.FindMatch)
}

func (h *handler) PutMatch(w http.ResponseWriter, r *http.Request) {
	h.Put(w, r, h.Client.InsertMatch)
}

func (h *handler) UpdateMatch(w http.ResponseWriter, r *http.Request) {
	h.Update(w, r, h.Client.UpdateMatch)
}

func (h *handler) DeleteMatch(w http.ResponseWriter, r *http.Request) {
	h.Delete(w, r, h.Client.DeleteMatch)
}
