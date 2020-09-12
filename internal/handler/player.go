package handler

import (
	"net/http"
	"net/url"
)

func (h *handler) GetPlayers(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	h.Get(w, r, h.contextFindPlayers(r.URL.Query()))
}

func (h *handler) GetPlayer(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	h.GetID(w, r, h.Octane.FindPlayers)
}

func (h *handler) PutPlayer(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	h.Put(w, r, h.Octane.InsertPlayerWithReader)
}

func (h *handler) UpdatePlayer(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	h.Update(w, r, h.Octane.UpdatePlayerWithReader)
}

func (h *handler) DeletePlayer(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	h.Delete(w, r, h.Octane.DeletePlayer)
}

func (h *handler) contextFindPlayers(v url.Values) *FindContext {
	return &FindContext{
		Do:         h.Octane.FindPlayers,
		Filter:     getBasicFilters(v),
		Pagination: getPagination(v),
		Sort:       getSort(v),
	}
}
