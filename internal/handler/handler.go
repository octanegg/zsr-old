package handler

import (
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/octanegg/zsr/internal/config"
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
	GetPlayerTeamStats(http.ResponseWriter, *http.Request)
	GetPlayerOpponentStats(http.ResponseWriter, *http.Request)
	GetPlayerEventStats(http.ResponseWriter, *http.Request)

	GetTeamStats(http.ResponseWriter, *http.Request)
	GetTeamOpponentStats(http.ResponseWriter, *http.Request)
	GetTeamEventStats(http.ResponseWriter, *http.Request)

	GetPlayerMetrics(http.ResponseWriter, *http.Request)
	GetTeamMetrics(http.ResponseWriter, *http.Request)

	GetActiveTeams(http.ResponseWriter, *http.Request)
	GetEventParticipants(http.ResponseWriter, *http.Request)

	CreateEvent(http.ResponseWriter, *http.Request)
	CreatePlayer(http.ResponseWriter, *http.Request)
	CreateTeam(http.ResponseWriter, *http.Request)
	CreateMatch(http.ResponseWriter, *http.Request)
	CreateGame(http.ResponseWriter, *http.Request)

	UpdateEvent(http.ResponseWriter, *http.Request)
	UpdatePlayer(http.ResponseWriter, *http.Request)
	UpdateTeam(http.ResponseWriter, *http.Request)
	UpdateMatch(http.ResponseWriter, *http.Request)
	UpdateGame(http.ResponseWriter, *http.Request)

	DeleteEvent(http.ResponseWriter, *http.Request)
	DeleteMatch(http.ResponseWriter, *http.Request)
	DeleteGame(http.ResponseWriter, *http.Request)

	UpdateMatches(http.ResponseWriter, *http.Request)
	MergePlayers(http.ResponseWriter, *http.Request)

	GetMatchGame(http.ResponseWriter, *http.Request)
	GetMatchGames(http.ResponseWriter, *http.Request)

	GetEventMatches(http.ResponseWriter, *http.Request)

	Search(http.ResponseWriter, *http.Request)
	SearchAdvanced(http.ResponseWriter, *http.Request)
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

	if os.Getenv(config.EnvIsInternal) == "true" && (p == 0 || pp == 0) {
		return nil
	}

	if p == 0 {
		p = 1
	}

	if pp == 0 {
		pp = 50
	}

	return &collection.Pagination{
		Page:    p,
		PerPage: pp,
	}
}
