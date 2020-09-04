package handler

import (
	"net/http"
	"net/url"

	"go.mongodb.org/mongo-driver/bson"
)

func (h *handler) GetMatches(w http.ResponseWriter, r *http.Request) {
	h.Get(w, r, h.contextFindMatches(r.URL.Query()))
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

func (h *handler) contextFindMatches(v url.Values) *FindContext {
	a := bson.A{getBasicFilters(v)}
	if playersFilter := getPTFilters(v); playersFilter != nil {
		a = append(a, playersFilter)
	}

	return &FindContext{
		Do:         h.Client.FindMatches,
		Filter:     bson.M{"$and": a},
		Pagination: getPagination(v),
		Sort:       getSort(v),
	}
}
