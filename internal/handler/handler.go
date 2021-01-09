package handler

import (
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/octanegg/zsr/octane"
	"github.com/octanegg/zsr/octane/collection"
	"go.mongodb.org/mongo-driver/bson"
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

	GetPlayerRecords(http.ResponseWriter, *http.Request)
	GetTeamRecords(http.ResponseWriter, *http.Request)
	GetGameRecords(http.ResponseWriter, *http.Request)
	GetSeriesRecords(http.ResponseWriter, *http.Request)

	GetPlayerStats(http.ResponseWriter, *http.Request)
	GetTeamStats(http.ResponseWriter, *http.Request)

	GetPlayerTeams(http.ResponseWriter, *http.Request)
}

// New .
func New(o octane.Client) Handler {
	return &handler{o}
}

func sort(v url.Values) bson.M {
	vals := v["sort"]
	sort := bson.M{}
	for _, val := range vals {
		fields := strings.Split(val, ":")

		var order int
		switch strings.ToLower(fields[1]) {
		case "asc":
			order = 1
		case "desc":
			order = -1
		default:
			continue
		}

		sort[strings.ToLower(fields[0])] = order
	}

	return sort
}

func pagination(v url.Values) *collection.Pagination {
	page, perPage := v.Get("page"), v.Get("perPage")
	p, _ := strconv.ParseInt(page, 10, 64)
	pp, _ := strconv.ParseInt(perPage, 10, 64)
	if p == 0 || pp == 0 {
		return nil
	}

	return &collection.Pagination{
		Page:    p,
		PerPage: pp,
	}
}
