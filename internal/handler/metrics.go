package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/octanegg/zsr/octane/pipelines"
)

func (h *handler) GetPlayerMetrics(w http.ResponseWriter, r *http.Request) {
	v := r.URL.Query()

	pipeline := pipelines.PlayerMetrics(statlinesFilter(v), v["stat"])
	data, err := h.Octane.Statlines().Pipeline(pipeline.Pipeline, pipeline.Decode)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Error{time.Now(), err.Error()})
		return
	}

	resp := struct {
		Metrics []interface{} `json:"metrics"`
	}{data}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

func (h *handler) GetTeamMetrics(w http.ResponseWriter, r *http.Request) {
	v := r.URL.Query()

	pipeline := pipelines.TeamMetrics(statlinesFilter(v), v["stat"])
	data, err := h.Octane.Statlines().Pipeline(pipeline.Pipeline, pipeline.Decode)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Error{time.Now(), err.Error()})
		return
	}

	resp := struct {
		Metrics []interface{} `json:"metrics"`
	}{data}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}
