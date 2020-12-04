package stats

import (
	"sort"

	"github.com/octanegg/zsr/octane"
	"go.mongodb.org/mongo-driver/bson"
)

// PlayerAggregate .
type PlayerAggregate struct {
	Player        *octane.Player           `json:"player" bson:"player,omitempty"`
	Team          *octane.Team             `json:"team,omitempty" bson:"team,omitempty"`
	Event         *octane.Event            `json:"event,omitempty" bson:"event,omitempty"`
	Games         int                      `json:"games" bson:"games"`
	Wins          int                      `json:"wins" bson:"wins"`
	WinPercentage float64                  `json:"win_percentage" bson:"win_percentage"`
	Totals        *PlayerAggregateTotals   `json:"totals" bson:"totals"`
	Averages      *PlayerAggregateAverages `json:"averages" bson:"averages"`
}

// PlayerAggregateTotals .
type PlayerAggregateTotals struct {
	Score   int     `json:"score" bson:"score"`
	Goals   int     `json:"goals" bson:"goals"`
	Assists int     `json:"assists" bson:"assists"`
	Saves   int     `json:"saves" bson:"saves"`
	Shots   int     `json:"shots" bson:"shots"`
	Rating  float64 `json:"rating" bson:"rating"`
}

// PlayerAggregateAverages .
type PlayerAggregateAverages struct {
	Score   float64 `json:"score" bson:"score"`
	Goals   float64 `json:"goals" bson:"goals"`
	Assists float64 `json:"assists" bson:"assists"`
	Saves   float64 `json:"saves" bson:"saves"`
	Shots   float64 `json:"shots" bson:"shots"`
	Rating  float64 `json:"rating" bson:"rating"`
}

func (s *stats) GetPlayerAggregate(groupBy string, filter bson.M) ([]*PlayerAggregate, error) {
	data, err := s.Statlines.Find(filter, nil, nil)
	if err != nil {
		return nil, err
	}

	index := make(map[string]int)
	res := []*PlayerAggregate{}
	for _, d := range data {
		s := d.(octane.Statline)
		g := getGroupByField(groupBy, &s)

		i, ok := index[g]
		if !ok {
			i = len(res)
			index[g] = i
			res = append(res, &PlayerAggregate{
				Player:   s.Player,
				Team:     s.Team,
				Event:    s.Game.Match.Event,
				Totals:   &PlayerAggregateTotals{},
				Averages: &PlayerAggregateAverages{},
			})
		}

		res[i].Games++
		if s.Winner {
			res[i].Wins++
		}

		res[i].Totals.Score += s.Stats.Core.Score
		res[i].Totals.Goals += s.Stats.Core.Goals
		res[i].Totals.Assists += s.Stats.Core.Assists
		res[i].Totals.Saves += s.Stats.Core.Saves
		res[i].Totals.Shots += s.Stats.Core.Shots
		res[i].Totals.Rating += s.Stats.Core.Rating

		res[i].WinPercentage = float64(res[i].Wins) / float64(res[i].Games)
		res[i].Averages.Score = float64(res[i].Totals.Score) / float64(res[i].Games)
		res[i].Averages.Goals = float64(res[i].Totals.Goals) / float64(res[i].Games)
		res[i].Averages.Assists = float64(res[i].Totals.Assists) / float64(res[i].Games)
		res[i].Averages.Saves = float64(res[i].Totals.Saves) / float64(res[i].Games)
		res[i].Averages.Shots = float64(res[i].Totals.Shots) / float64(res[i].Games)
		res[i].Averages.Rating = float64(res[i].Totals.Rating) / float64(res[i].Games)
	}

	sort.Slice(res, func(i, j int) bool {
		return res[i].Games > res[j].Games
	})

	return res, nil
}

func getGroupByField(field string, s *octane.Statline) string {
	switch field {
	case "player":
		return s.Player.ID.Hex()
	case "team":
		return s.Team.ID.Hex()
	case "event":
		return s.Game.Match.Event.ID.Hex()
	default:
		return ""
	}
}
