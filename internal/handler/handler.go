package handler

import (
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/octanegg/core/octane"
)

const (
	contentType     = "Content-Type"
	applicationJSON = "application/json"
	errContentType  = "Content-Type header is not application/json"
)

// Error .
type Error struct {
	Timestamp time.Time `json:"timestamp"`
	Error     string    `json:"error"`
}

type handler struct {
	Client octane.Client
}

// Handler .
type Handler interface {
	Health(http.ResponseWriter, *http.Request)

	GetEvent(http.ResponseWriter, *http.Request)
	GetEvents(http.ResponseWriter, *http.Request)
	GetMatch(http.ResponseWriter, *http.Request)
	GetMatches(http.ResponseWriter, *http.Request)
	GetGame(http.ResponseWriter, *http.Request)
	GetGames(http.ResponseWriter, *http.Request)
	GetPlayer(http.ResponseWriter, *http.Request)
	GetPlayers(http.ResponseWriter, *http.Request)
	GetTeam(http.ResponseWriter, *http.Request)
	GetTeams(http.ResponseWriter, *http.Request)

	PutEvent(http.ResponseWriter, *http.Request)
	PutMatch(http.ResponseWriter, *http.Request)
	PutGame(http.ResponseWriter, *http.Request)
	PutPlayer(http.ResponseWriter, *http.Request)
	PutTeam(http.ResponseWriter, *http.Request)

	UpdateEvent(http.ResponseWriter, *http.Request)
	UpdateMatch(http.ResponseWriter, *http.Request)
	UpdateGame(http.ResponseWriter, *http.Request)
	UpdatePlayer(http.ResponseWriter, *http.Request)
	UpdateTeam(http.ResponseWriter, *http.Request)

	DeleteEvent(http.ResponseWriter, *http.Request)
	DeleteMatch(http.ResponseWriter, *http.Request)
	DeleteGame(http.ResponseWriter, *http.Request)
	DeletePlayer(http.ResponseWriter, *http.Request)
	DeleteTeam(http.ResponseWriter, *http.Request)
}

// NewHandler .
func NewHandler(client octane.Client) Handler {
	return &handler{
		Client: client,
	}
}

func getPagination(v url.Values) *octane.Pagination {
	page, _ := strconv.ParseInt(v.Get("page"), 10, 64)
	perPage, _ := strconv.ParseInt(v.Get("per_page"), 10, 64)
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
	switch v.Get("order") {
	case "asc":
		order = 1
	case "desc":
		order = -1
	default:
		return nil
	}

	return &octane.Sort{
		Field: v.Get("sort"),
		Order: order,
	}
}
