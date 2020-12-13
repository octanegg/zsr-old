package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/gorilla/mux"
	"github.com/octanegg/zsr/octane/collection"
	"github.com/octanegg/zsr/octane/filter"
	"github.com/octanegg/zsr/octane/pipelines"
	"go.mongodb.org/mongo-driver/bson"
)

func (h *handler) GetGameRecords(w http.ResponseWriter, r *http.Request) {
	ctx := h.getRecordsContext(r)
	if ctx == nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Error{time.Now(), "invalid category, type, or stat"})
		return

	}

	var (
		data []interface{}
		err  error
	)

	if ctx.Pipeline != nil {
		data, err = ctx.Collection.Pipeline(ctx.Pipeline.Pipeline, ctx.Pipeline.Decode)
	} else {
		data, err = ctx.Collection.Find(
			ctx.Filter,
			ctx.Sort,
			&collection.Pagination{Page: 1, PerPage: 25},
		)
	}

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Error{time.Now(), err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}

type recordsContext struct {
	Collection collection.Collection
	Filter     bson.M
	Sort       bson.M
	Pipeline   *pipelines.Pipeline
}

func (h *handler) getRecordsContext(r *http.Request) *recordsContext {
	var (
		vars     = mux.Vars(r)
		category = vars["category"]
		typ      = vars["type"]
		stat     = vars["stat"]
	)

	switch category {
	case "games":
		switch typ {
		case "players":
			return &recordsContext{
				Collection: h.Octane.Statlines(),
				Filter:     statlinesFilter(r.URL.Query()),
				Sort:       bson.M{fmt.Sprintf("stats.core.%s", stat): -1},
			}
		case "teams":
			return &recordsContext{
				Collection: h.Octane.Statlines(),
				Pipeline: pipelines.GameTeamRecords(
					statlinesFilter(r.URL.Query()),
					stat,
				),
			}
		case "totals":
			return &recordsContext{
				Collection: h.Octane.Games(),
				Pipeline: pipelines.GameTotalRecords(
					gamesFilter(r.URL.Query()),
					stat,
				),
			}
		case "differentials":
			return &recordsContext{
				Collection: h.Octane.Games(),
				Pipeline: pipelines.GameDifferentialRecords(
					gamesFilter(r.URL.Query()),
					stat,
				),
			}
		case "overtimes":
			return &recordsContext{
				Collection: h.Octane.Games(),
				Filter:     gamesFilter(r.URL.Query()),
				Sort:       bson.M{"duration": -1},
			}
		default:
			return nil
		}
	case "matches":
		switch typ {
		case "players":
			return nil
		case "teams":
			return nil
		case "totals":
			return nil
		case "differentials":
			return nil
		default:
			return nil
		}
	default:
		return nil
	}
}

func statlinesFilter(v url.Values) bson.M {
	return filter.New(
		filter.Strings("game.match.event.tier", v["tier"]),
		filter.Strings("game.match.event.region", v["region"]),
		filter.Ints("game.match.event.mode", v["mode"]),
		filter.ObjectIDs("player._id", v["player"]),
		filter.ObjectIDs("team._id", v["team"]),
		filter.ObjectIDs("opponent._id", v["opponent"]),
		filter.Dates("game.date", v.Get("before"), v.Get("after")),
		filter.Bool("winner", v.Get("winner")),
	)
}

func teamlinesFilter(v url.Values) bson.M {
	return filter.New(
		filter.Strings("game.match.event.tier", v["tier"]),
		filter.Strings("game.match.event.region", v["region"]),
		filter.Ints("game.match.event.mode", v["mode"]),
		filter.ObjectIDs("player._id", v["player"]),
		filter.ObjectIDs("team._id", v["team"]),
		filter.ObjectIDs("opponent._id", v["opponent"]),
		filter.Dates("game.date", v.Get("before"), v.Get("after")),
		filter.Bool("winner", v.Get("winner")),
	)
}
