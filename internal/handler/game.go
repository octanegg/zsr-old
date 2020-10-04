package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/octanegg/zsr/internal/config"
	"go.mongodb.org/mongo-driver/bson"
)

func (h *handler) GetGames(w http.ResponseWriter, r *http.Request) {
	v := r.URL.Query()
	a := bson.A{getBasicFilters(v)}
	if playersFilter := getPTFiltersWithElemMatch(v); playersFilter != nil {
		a = append(a, playersFilter)
	}

	data, err := h.Octane.FindGames(bson.M{config.KeyAnd: a}, getPagination(v), getSort(v))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Error{time.Now(), err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}
