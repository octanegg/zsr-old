package handler

import (
	"encoding/json"
	"net/http"
	"time"
)

func (h *handler) Health(w http.ResponseWriter, r *http.Request) {
	if err := h.Octane.Ping(); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Error{time.Now(), err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
}
