package handler

import (
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/octanegg/core/internal/config"
	"github.com/octanegg/core/octane"
)

// Error .
type Error struct {
	Timestamp time.Time `json:"timestamp"`
	Error     string    `json:"error"`
}

type handler struct {
	Octane octane.Client
}

// Handler .
type Handler interface {
	Health(http.ResponseWriter, *http.Request)

	GetEvents(http.ResponseWriter, *http.Request)
	GetMatches(http.ResponseWriter, *http.Request)
	GetGames(http.ResponseWriter, *http.Request)
	GetPlayers(http.ResponseWriter, *http.Request)
	GetTeams(http.ResponseWriter, *http.Request)
}

// New .
func New(o octane.Client) Handler {
	return &handler{o}
}

func getPagination(v url.Values) *octane.Pagination {
	page, _ := strconv.ParseInt(v.Get(config.ParamPage), 10, 64)
	perPage, _ := strconv.ParseInt(v.Get(config.ParamPage), 10, 64)
	if page == 0 || perPage == 0 {
		return nil
	}

	return &octane.Pagination{
		Page:    page,
		PerPage: perPage,
	}
}

func getSort(v url.Values) *octane.Sort {
	var order int
	switch v.Get(config.ParamOrder) {
	case config.ParamAscending:
		order = 1
	case config.ParamDescending:
		order = -1
	default:
		return nil
	}

	return &octane.Sort{
		Field: v.Get(config.ParamSort),
		Order: order,
	}
}