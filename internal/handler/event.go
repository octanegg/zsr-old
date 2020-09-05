package handler

import (
	"net/http"
	"net/url"
)

func (h *handler) GetEvents(w http.ResponseWriter, r *http.Request) {
	h.Get(w, r, h.contextFindEvents(r.URL.Query()))
}

func (h *handler) GetEvent(w http.ResponseWriter, r *http.Request) {
	h.GetID(w, r, h.Client.FindEvent)
}

func (h *handler) PutEvent(w http.ResponseWriter, r *http.Request) {
	h.Put(w, r, h.Client.InsertEventWithReader)
}

func (h *handler) UpdateEvent(w http.ResponseWriter, r *http.Request) {
	h.Update(w, r, h.Client.UpdateEventWithReader)
}

func (h *handler) DeleteEvent(w http.ResponseWriter, r *http.Request) {
	h.Delete(w, r, h.Client.DeleteEvent)
}

func (h *handler) contextFindEvents(v url.Values) *FindContext {
	return &FindContext{
		Do:         h.Client.FindEvents,
		Filter:     getBasicFilters(v),
		Pagination: getPagination(v),
		Sort:       getSort(v),
	}
}
