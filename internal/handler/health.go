package handler

import (
	"net/http"
)

func (h *handler) Health(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
