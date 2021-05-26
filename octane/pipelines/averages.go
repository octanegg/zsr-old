package pipelines

import (
	"github.com/octanegg/zsr/octane/filter"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func AverageScore(goals, assists, saves, shots int) *Pipeline {
	pipeline := New(
		Match(bson.M{
			"player.stats.core.goals":   goals,
			"player.stats.core.assists": assists,
			"player.stats.core.saves":   saves,
			"player.stats.core.shots":   shots,
		}),
		Group(bson.M{
			"_id": 0,
			"score": bson.M{
				"$avg": "$player.stats.core.score",
			},
		}),
	)

	return &Pipeline{
		Pipeline: pipeline,
		Decode: func(cursor *mongo.Cursor) (interface{}, error) {
			var score struct {
				Score float64 `json:"score" bson:"score"`
			}
			if err := cursor.Decode(&score); err != nil {
				return nil, err
			}
			return score, nil
		},
	}
}

func Averages() *Pipeline {
	pipeline := New(
		Match(filter.New(
			filter.AfterDate("game.date", "2019-01-01"),
			filter.Strings("game.match.event.tier", []string{"A", "S"}),
			filter.Ints("game.match.event.mode", []string{"3"}),
			filter.Bool("game.match.stage.qualifier", "false"),
		)),
		Group(bson.M{
			"_id": 0,
			"games": bson.M{
				"$sum": 1,
			},
			"score": bson.M{
				"$sum": "$player.stats.core.score",
			},
			"goals": bson.M{
				"$sum": "$player.stats.core.goals",
			},
			"assists": bson.M{
				"$sum": "$player.stats.core.assists",
			},
			"saves": bson.M{
				"$sum": "$player.stats.core.saves",
			},
			"shots": bson.M{
				"$sum": "$player.stats.core.shots",
			},
			"teamGoals": bson.M{
				"$sum": "$team.stats.core.goals",
			},
		}),
		Project(bson.M{
			"score": bson.M{
				"$divide": bson.A{"$score", "$games"},
			},
			"goals": bson.M{
				"$divide": bson.A{"$goals", "$games"},
			},
			"assists": bson.M{
				"$divide": bson.A{"$assists", "$games"},
			},
			"saves": bson.M{
				"$divide": bson.A{"$saves", "$games"},
			},
			"shots": bson.M{
				"$divide": bson.A{"$shots", "$games"},
			},
			"shootingPercentage": bson.M{
				"$divide": bson.A{"$goals", "$shots"},
			},
			"goalParticipation": bson.M{
				"$divide": bson.A{bson.M{"$add": bson.A{"$goals", "$assists"}}, "$teamGoals"},
			},
		}),
	)

	return &Pipeline{
		Pipeline: pipeline,
		Decode: func(cursor *mongo.Cursor) (interface{}, error) {
			var score struct {
				Score              float64 `json:"score" bson:"score"`
				Goals              float64 `json:"goals" bson:"goals"`
				Assists            float64 `json:"assists" bson:"assists"`
				Saves              float64 `json:"saves" bson:"saves"`
				Shots              float64 `json:"shots" bson:"shots"`
				ShootingPercentage float64 `json:"shootingPercentage" bson:"shootingPercentage"`
				GoalParticipation  float64 `json:"goalParticipation" bson:"goalParticipation"`
			}
			if err := cursor.Decode(&score); err != nil {
				return nil, err
			}
			return score, nil
		},
	}
}
