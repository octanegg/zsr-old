package handler

import (
	"net/url"
	"strconv"

	"github.com/octanegg/core/internal/config"
	"github.com/octanegg/core/octane"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

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
			(*filter)[config.KeyID] = bson.M{"$in": ids}
		} else {
			(*filter)[key] = bson.M{"$in": ids}
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
