package handler

import (
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/octanegg/zsr/internal/config"
	"github.com/octanegg/zsr/octane"
	"go.mongodb.org/mongo-driver/bson"
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

func getBasicFilters(v url.Values) bson.M {
	filter := bson.M{}
	put(&filter, v, config.ParamTier)
	put(&filter, v, config.ParamRegion)
	put(&filter, v, config.ParamCountry)
	put(&filter, v, config.ParamTag)
	put(&filter, v, config.ParamActiveTeam)
	put(&filter, v, config.ParamName)
	put(&filter, v, config.ParamAccountID)
	put(&filter, v, config.ParamAccountPlatform)
	putInt(&filter, v, config.ParamMode)
	putInt(&filter, v, config.ParamStage)
	putInt(&filter, v, config.ParamSubstage)
	putID(&filter, v, config.ParamEvent)
	putID(&filter, v, config.ParamMatch)
	putID(&filter, v, config.ParamID)
	return filter
}

func getPTFilters(v url.Values) bson.M {
	blue, orange := bson.M{}, bson.M{}
	if players := stringsToObjectIDs(v[config.ParamPlayer]); len(players) > 0 {
		blue["blue.players"] = bson.M{"$all": players}
		orange["orange.players"] = bson.M{"$all": players}
		if opponents := stringsToObjectIDs(v[config.ParamOpponent]); len(opponents) > 0 {
			blue["orange.players"] = bson.M{"$all": opponents}
			orange["blue.players"] = bson.M{"$all": opponents}
		}
	}

	if teams := stringsToObjectIDs(v[config.ParamTeam]); len(teams) > 0 {
		blue["blue.team"] = bson.M{"$in": teams}
		orange["orange.team"] = bson.M{"$in": teams}
	}

	if vs := stringsToObjectIDs(v[config.ParamVs]); len(vs) > 0 {
		blue["orange.team"] = bson.M{"$in": vs}
		orange["blue.team"] = bson.M{"$in": vs}
	}

	return bson.M{config.KeyOr: bson.A{blue, orange}}
}

func getPTFiltersWithElemMatch(v url.Values) bson.M {
	blue, orange := bson.M{}, bson.M{}
	if players := stringsToObjectIDs(v[config.ParamPlayer]); len(players) > 0 {
		blue["blue.players"] = bson.M{
			"$all": bson.A{
				bson.M{"$elemMatch": bson.M{"player": bson.M{"$in": players}}},
			},
		}
		orange["orange.players"] = bson.M{
			"$all": bson.A{
				bson.M{"$elemMatch": bson.M{"player": bson.M{"$in": players}}},
			},
		}

		if opponents := stringsToObjectIDs(v[config.ParamOpponent]); len(opponents) > 0 {
			blue["orange.players"] = bson.M{
				"$all": bson.A{
					bson.M{"$elemMatch": bson.M{"player": bson.M{"$in": opponents}}},
				},
			}
			orange["blue.players"] = bson.M{
				"$all": bson.A{
					bson.M{"$elemMatch": bson.M{"player": bson.M{"$in": opponents}}},
				},
			}
		}
	}

	if teams := stringsToObjectIDs(v[config.ParamTeam]); len(teams) > 0 {
		blue["blue.team"] = bson.M{"$in": teams}
		orange["orange.team"] = bson.M{"$in": teams}
	}

	if vs := stringsToObjectIDs(v[config.ParamVs]); len(vs) > 0 {
		blue["orange.team"] = bson.M{"$in": vs}
		orange["blue.team"] = bson.M{"$in": vs}
	}

	return bson.M{config.KeyOr: bson.A{blue, orange}}

}

func put(filter *bson.M, v url.Values, key string) {
	if vals, ok := v[key]; ok {
		(*filter)[key] = bson.M{"$in": vals}
	}
}

func putInt(filter *bson.M, v url.Values, key string) {
	if vals, ok := v[key]; ok {
		var a []int
		for _, val := range vals {
			if i, err := strconv.Atoi(val); err == nil {
				a = append(a, i)
			}
		}
		(*filter)[key] = bson.M{"$in": a}
	}
}

func putID(filter *bson.M, v url.Values, key string) {
	if vals, ok := v[key]; ok {
		ids := stringsToObjectIDs(vals)
		if key == config.ParamID {
			(*filter)["_id"] = bson.M{"$in": ids}
		} else {
			(*filter)[key] = bson.M{"$in": ids}
		}
	}
}

func putAfterDate(filter *bson.M, v url.Values, key string) {
	if val := v.Get(key); val != "" {
		if t, err := time.Parse("2006-01-02T03:04:05Z", val); err == nil {
			(*filter)[config.ParamStartDate] = bson.M{"$gte": t}
		}
	}
}

func stringsToObjectIDs(vals []string) []primitive.ObjectID {
	var a []primitive.ObjectID
	for _, val := range vals {
		if i, err := primitive.ObjectIDFromHex(val); err == nil {
			a = append(a, i)
		}
	}
	return a
}
