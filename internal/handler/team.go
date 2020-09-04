package handler

import (
	"net/http"
)

func (h *handler) GetTeams(w http.ResponseWriter, r *http.Request) {
	h.Get(w, r, h.Client.FindTeams)
}

func (h *handler) GetTeam(w http.ResponseWriter, r *http.Request) {
	h.GetID(w, r, h.Client.FindTeam)
}

func (h *handler) PutTeam(w http.ResponseWriter, r *http.Request) {
	h.Put(w, r, h.Client.InsertTeam)
}

func (h *handler) UpdateTeam(w http.ResponseWriter, r *http.Request) {
	h.Update(w, r, h.Client.UpdateTeam)
}

func (h *handler) DeleteTeam(w http.ResponseWriter, r *http.Request) {
	h.Delete(w, r, h.Client.DeleteTeam)
}
