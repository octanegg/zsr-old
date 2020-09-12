package handler

import (
	"net/http"
	"net/url"
)

func (h *handler) GetTeams(w http.ResponseWriter, r *http.Request) {
	h.Get(w, r, h.contextFindTeams(r.URL.Query()))
}

func (h *handler) GetTeam(w http.ResponseWriter, r *http.Request) {
	h.GetID(w, r, h.Octane.FindTeams)
}

func (h *handler) PutTeam(w http.ResponseWriter, r *http.Request) {
	h.Put(w, r, h.Octane.InsertTeamWithReader)
}

func (h *handler) UpdateTeam(w http.ResponseWriter, r *http.Request) {
	h.Update(w, r, h.Octane.UpdateTeamWithReader)
}

func (h *handler) DeleteTeam(w http.ResponseWriter, r *http.Request) {
	h.Delete(w, r, h.Octane.DeleteTeam)
}

func (h *handler) contextFindTeams(v url.Values) *FindContext {
	return &FindContext{
		Do:         h.Octane.FindTeams,
		Filter:     getBasicFilters(v),
		Pagination: getPagination(v),
		Sort:       getSort(v),
	}
}
