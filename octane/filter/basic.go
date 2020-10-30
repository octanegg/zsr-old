package filter

import (
	"net/url"

	"go.mongodb.org/mongo-driver/bson"
)

// Events .
func Events(v url.Values) bson.M {
	return New(
		Strings("name", v["name"]),
		Strings("tier", v["tier"]),
		Strings("region", v["region"]),
		Strings("mode", v["mode"]),
		BeforeDate("start_date", v.Get("before")),
		AfterDate("start_date", v.Get("after")),
	)
}

// Matches .
func Matches(v url.Values) bson.M {
	return New(
		ObjectIDs("event._id", v["event"]),
		Strings("event.tier", v["tier"]),
		Strings("event.region", v["region"]),
		Ints("event.mode", v["mode"]),
		Ints("stage._id", v["stage"]),
		Ints("substage", v["substage"]),
		BeforeDate("date", v.Get("before")),
		AfterDate("date", v.Get("after")),
	)
}

// Games .
func Games(v url.Values) bson.M {
	return New(
		ObjectIDs("match.event._id", v["event"]),
		Strings("match.event.tier", v["tier"]),
		Strings("match.event.region", v["region"]),
		Ints("match.event.mode", v["mode"]),
		Ints("match.stage._id", v["stage"]),
		Ints("match.substage", v["substage"]),
		ObjectIDs("match._id", v["event"]),
		BeforeDate("date", v.Get("before")),
		AfterDate("date", v.Get("after")),
	)
}

// Players .
func Players(v url.Values) bson.M {
	return New(
		Strings("country", v["country"]),
		Strings("tag", v["tag"]),
		ObjectIDs("team", v["team"]),
	)
}

// Teams .
func Teams(v url.Values) bson.M {
	return New(Strings("name", v["name"]))
}
