package handler

import (
	"net/http"
	"net/url"
)

func (h *handler) GetEvents(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	h.Get(w, r, h.contextFindEvents(r.URL.Query()))
}

func (h *handler) GetEvent(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	h.GetID(w, r, h.Octane.FindEvents)
}

func (h *handler) PutEvent(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	h.Put(w, r, h.Octane.InsertEventWithReader)
}

func (h *handler) UpdateEvent(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	h.Update(w, r, h.Octane.UpdateEventWithReader)
}

func (h *handler) DeleteEvent(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	h.Delete(w, r, h.Octane.DeleteEvent)
}

func (h *handler) contextFindEvents(v url.Values) *FindContext {
	return &FindContext{
		Do:         h.Octane.FindEvents,
		Filter:     getBasicFilters(v),
		Pagination: getPagination(v),
		Sort:       getSort(v),
	}
}
