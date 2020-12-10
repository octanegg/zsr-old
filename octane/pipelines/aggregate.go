package pipelines

import "go.mongodb.org/mongo-driver/bson"

var (
	groupByPlayer = bson.M{
		"$group": bson.M{
			"_id": "$player._id",
			"player": bson.M{
				"$first": "$player",
			},
			"team": bson.M{
				"$first": "$team",
			},
			"event": bson.M{
				"$first": "$game.match.event",
			},
			"games": bson.M{
				"$sum": 1,
			},
			"wins": bson.M{
				"$sum": bson.M{
					"$cond": bson.A{
						bson.M{
							"$eq": bson.A{
								"$winner", true,
							},
						}, 1, 0,
					},
				},
			},
			"score_total": bson.M{
				"$sum": "$stats.core.score",
			},
			"goals_total": bson.M{
				"$sum": "$stats.core.goals",
			},
			"assists_total": bson.M{
				"$sum": "$stats.core.assists",
			},
			"saves_total": bson.M{
				"$sum": "$stats.core.saves",
			},
			"shots_total": bson.M{
				"$sum": "$stats.core.shots",
			},
			"score_avg": bson.M{
				"$avg": "$stats.core.score",
			},
			"goals_avg": bson.M{
				"$avg": "$stats.core.goals",
			},
			"assists_avg": bson.M{
				"$avg": "$stats.core.assists",
			},
			"saves_avg": bson.M{
				"$avg": "$stats.core.saves",
			},
			"shots_avg": bson.M{
				"$avg": "$stats.core.shots",
			},
			"rating_avg": bson.M{
				"$avg": "$stats.core.rating",
			},
		},
	}

	projectPlayerAggregate = bson.M{
		"$project": bson.M{
			"_id":    "$_id",
			"player": "$player",
			"team":   "$team",
			"event":  "$event",
			"games":  "$games",
			"wins":   "$wins",
			"win_percentage": bson.M{
				"$divide": bson.A{
					"$wins", "$games",
				},
			},
			"totals": bson.M{
				"score":   "$score_total",
				"goals":   "$goals_total",
				"assists": "$assists_total",
				"saves":   "$saves_total",
				"shots":   "$shots_total",
			},
			"averages": bson.M{
				"score":   "$score_avg",
				"goals":   "$goals_avg",
				"assists": "$assists_avg",
				"saves":   "$saves_avg",
				"shots":   "$shots_avg",
				"rating":  "$rating_avg",
			},
		},
	}
)

// PlayerAggregate .
func PlayerAggregate(filter bson.M, having bson.M) []bson.M {
	pipeline := []bson.M{{"$match": filter}, groupByPlayer}

	if having != nil {
		pipeline = append(pipeline, bson.M{"$match": having})
	}

	return append(pipeline, projectPlayerAggregate)
}
