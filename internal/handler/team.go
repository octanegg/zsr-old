package handler

import (
	"net/http"
	"net/url"
)

func (h *handler) GetTeams(w http.ResponseWriter, r *http.Request) {
	h.Get(w, r, h.contextFindTeams(r.URL.Query()))
}

func (h *handler) GetTeam(w http.ResponseWriter, r *http.Request) {
	h.GetID(w, r, h.Client.FindTeams)
}

func (h *handler) PutTeam(w http.ResponseWriter, r *http.Request) {
	h.Put(w, r, h.Client.InsertTeamWithReader)
}

func (h *handler) UpdateTeam(w http.ResponseWriter, r *http.Request) {
	h.Update(w, r, h.Client.UpdateTeamWithReader)
}

func (h *handler) DeleteTeam(w http.ResponseWriter, r *http.Request) {
	h.Delete(w, r, h.Client.DeleteTeam)
}

func (h *handler) contextFindTeams(v url.Values) *FindContext {
	return &FindContext{
		Do:         h.Client.FindTeams,
		Filter:     getBasicFilters(v),
		Pagination: getPagination(v),
		Sort:       getSort(v),
	}
}
