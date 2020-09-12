package admin

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/octanegg/core/internal/config"
	"github.com/octanegg/core/internal/deprecated"
)

func (h *handler) UpdateMatch(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get(config.HeaderContentType) != config.HeaderApplicationJSON {
		w.WriteHeader(http.StatusUnsupportedMediaType)
		json.NewEncoder(w).Encode(Error{time.Now(), config.ErrInvalidContentType})
		return
	}

	var ctx deprecated.UpdateMatchContext
	if err := json.NewDecoder(r.Body).Decode(&ctx); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Error{time.Now(), err.Error()})
		return
	}

	if err := h.Deprecated.UpdateMatch(&ctx); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Error{time.Now(), err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
}
