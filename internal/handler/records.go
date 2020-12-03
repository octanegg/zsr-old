package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"github.com/octanegg/zsr/octane/records"
)

func (h *handler) GetRecords(w http.ResponseWriter, r *http.Request) {
	stat := mux.Vars(r)["stat"]
	if !records.IsValidStat(strings.ToLower(stat)) {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Error{time.Now(), fmt.Sprintf("valid stats: %s", records.ValidStats())})
		return
	}

	data, err := h.Octane.Records().Get(&records.Context{
		Category: "game",
		Type:     "player",
		Stat:     stat,
		Query:    r.URL.Query(),
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Error{time.Now(), err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}
