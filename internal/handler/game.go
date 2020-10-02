package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/octanegg/core/internal/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (h *handler) GetGames(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get(config.HeaderContentType) != config.HeaderApplicationJSON {
		w.WriteHeader(http.StatusUnsupportedMediaType)
		json.NewEncoder(w).Encode(Error{time.Now(), config.ErrInvalidContentType})
		return
	}
	defer r.Body.Close()

	var filter bson.M
	json.NewDecoder(r.Body).Decode(&filter)
	for _, field := range config.ObjectIDFields {
		if v, ok := filter[field]; ok {
			filter[field], _ = primitive.ObjectIDFromHex(v.(string))
		}
	}

	data, err := h.Octane.FindGames(filter, getPagination(r.URL.Query()), getSort(r.URL.Query()))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Error{time.Now(), err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}
