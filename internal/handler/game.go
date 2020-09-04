package handler

import (
	"net/http"
	"net/url"

	"go.mongodb.org/mongo-driver/bson"
)

func (h *handler) GetGames(w http.ResponseWriter, r *http.Request) {
	h.Get(w, r, h.contextFindGames(r.URL.Query()))
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

func (h *handler) contextFindGames(v url.Values) *FindContext {
	a := bson.A{getBasicFilters(v)}
	if playersFilter := getPTFiltersWithElemMatch(v); playersFilter != nil {
		a = append(a, playersFilter)
	}

	return &FindContext{
		Do:         h.Client.FindGames,
		Filter:     bson.M{"$and": a},
		Pagination: getPagination(v),
		Sort:       getSort(v),
	}
}
