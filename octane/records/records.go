package records

import (
	"fmt"
	"net/url"

	"github.com/octanegg/zsr/octane/collection"
	"github.com/octanegg/zsr/octane/filter"
	"go.mongodb.org/mongo-driver/bson"
)

const maxRecords = 25
const game = "game"

// Records .
type Records interface {
	Get(*Context) ([]interface{}, error)
}

type records struct {
	Stats collection.Collection
}

// Context .
type Context struct {
	Category string
	Type     string
	Stat     string
	Query    url.Values
}

// New .
func New(c collection.Collection) Records {
	return &records{c}
}

// Get .
func (r *records) Get(ctx *Context) ([]interface{}, error) {
	data, err := r.Stats.Find(getFilter(ctx), getSort(ctx), nil)
	if err != nil {
		return nil, err
	}

	if len(data) > maxRecords {
		data = data[:maxRecords]
	}

	return data, nil
}

func getSort(ctx *Context) bson.M {
	if ctx.Category == game {
		if ctx.Type == "player" {
			return bson.M{fmt.Sprintf("stats.core.%s", ctx.Stat): -1}
		}
	}
	return nil
}

func getFilter(ctx *Context) bson.M {
	if ctx.Category == game {
		return filter.New(
			filter.Strings("game.match.event.tier", ctx.Query["tier"]),
			filter.Strings("game.match.event.region", ctx.Query["region"]),
			filter.Ints("game.match.event.mode", ctx.Query["mode"]),
			filter.ObjectIDs("player._id", ctx.Query["player"]),
			filter.ObjectIDs("team._id", ctx.Query["team"]),
			filter.ObjectIDs("opponent._id", ctx.Query["opponent"]),
			filter.Dates("game.date", ctx.Query.Get("before"), ctx.Query.Get("after")),
			filter.Bool("winner", ctx.Query.Get("winner")),
		)
	}

	return nil
}
