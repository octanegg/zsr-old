package handler

import (
	"net/http"
)

func (h *handler) GetEvents(w http.ResponseWriter, r *http.Request) {
	h.Get(w, r, h.Client.FindEvents)
}

func (h *handler) GetEvent(w http.ResponseWriter, r *http.Request) {
	h.GetID(w, r, h.Client.FindEvent)
}

func (h *handler) PutEvent(w http.ResponseWriter, r *http.Request) {
	h.Put(w, r, h.Client.InsertEvent)
}

func (h *handler) UpdateEvent(w http.ResponseWriter, r *http.Request) {
	h.Update(w, r, h.Client.UpdateEvent)
}

func (h *handler) DeleteEvent(w http.ResponseWriter, r *http.Request) {
	h.Delete(w, r, h.Client.DeleteEvent)
}
