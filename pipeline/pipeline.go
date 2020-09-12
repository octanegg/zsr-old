package pipeline

import "go.mongodb.org/mongo-driver/bson"

var (
	lookupTeams = []bson.M{
		{
			"$lookup": bson.M{
				"from":         "teams",
				"localField":   "blue.team",
				"foreignField": "_id",
				"as":           "blue.team",
			},
		},
		{
			"$unwind": "$blue.team",
		},
		{
			"$lookup": bson.M{
				"from":         "teams",
				"localField":   "orange.team",
				"foreignField": "_id",
				"as":           "orange.team",
			},
		},
		{
			"$unwind": "$orange.team",
		},
	}
)

// MatchesWithTeamLookup .
func MatchesWithTeamLookup(filter bson.M) []bson.M {
	return append([]bson.M{{"$match": filter}}, lookupTeams...)
}
