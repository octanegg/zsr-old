package handler

import (
	"net/http"
	"net/url"
	"strconv"

	"github.com/octanegg/core/internal/config"
	"go.mongodb.org/mongo-driver/bson"
)

func (h *handler) GetMatches(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	ctx := h.contextFindMatches(r.URL.Query())
	if b, err := strconv.ParseBool(r.URL.Query().Get(config.ParamLookupTeams)); err == nil && b {
		ctx.Do = h.Octane.FindMatchesWithTeamLookup
	}
	h.Get(w, r, ctx)
}

func (h *handler) GetMatch(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	h.GetID(w, r, h.Octane.FindMatches)
}

func (h *handler) PutMatch(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	h.Put(w, r, h.Octane.InsertMatchWithReader)
}

func (h *handler) UpdateMatch(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	h.Update(w, r, h.Octane.UpdateMatchWithReader)
}

func (h *handler) DeleteMatch(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	h.Delete(w, r, h.Octane.DeleteMatch)
}

func (h *handler) contextFindMatches(v url.Values) *FindContext {
	a := bson.A{getBasicFilters(v)}
	if playersFilter := getPTFilters(v); playersFilter != nil {
		a = append(a, playersFilter)
	}

	return &FindContext{
		Do:         h.Octane.FindMatches,
		Filter:     bson.M{config.KeyAnd: a},
		Pagination: getPagination(v),
		Sort:       getSort(v),
	}
}
