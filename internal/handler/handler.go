package handler

import (
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/octanegg/zsr/octane"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

	GetEvent(http.ResponseWriter, *http.Request)
	GetMatch(http.ResponseWriter, *http.Request)
	GetGame(http.ResponseWriter, *http.Request)
	GetPlayer(http.ResponseWriter, *http.Request)
	GetTeam(http.ResponseWriter, *http.Request)
}

// New .
func New(o octane.Client) Handler {
	return &handler{o}
}

func getPagination(v url.Values) *octane.Pagination {
	page, _ := strconv.ParseInt(v.Get("page"), 10, 64)
	perPage, _ := strconv.ParseInt(v.Get("page"), 10, 64)
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

func toObjectIDs(vals []string) []primitive.ObjectID {
	var a []primitive.ObjectID
	for _, val := range vals {
		if i, err := primitive.ObjectIDFromHex(val); err == nil {
			a = append(a, i)
		}
	}
	return a
}

func toInts(vals []string) []int {
	var a []int
	for _, val := range vals {
		if i, err := strconv.Atoi(val); err == nil {
			a = append(a, i)
		}
	}
	return a
}
