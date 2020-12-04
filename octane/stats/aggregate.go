package stats

import (
	"sort"

	"github.com/octanegg/zsr/octane"
	"go.mongodb.org/mongo-driver/bson"
)

// PlayerAggregate .
type PlayerAggregate struct {
	Player  *octane.Player `json:"player" bson:"player,omitempty"`
	Team    *octane.Team   `json:"team,omitempty" bson:"team,omitempty"`
	Event   *octane.Event  `json:"event,omitempty" bson:"event,omitempty"`
	Games   int            `json:"games" bson:"games"`
	Wins    int            `json:"wins" bson:"wins"`
	Score   int            `json:"score" bson:"score"`
	Goals   int            `json:"goals" bson:"goals"`
	Assists int            `json:"assists" bson:"assists"`
	Saves   int            `json:"saves" bson:"saves"`
	Shots   int            `json:"shots" bson:"shots"`
	Rating  float64        `json:"rating" bson:"rating"`
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
				Player: s.Player,
				Team:   s.Team,
				Event:  s.Game.Match.Event,
			})
		}

		res[i].Games++
		res[i].Score += s.Stats.Core.Score
		res[i].Goals += s.Stats.Core.Goals
		res[i].Assists += s.Stats.Core.Assists
		res[i].Shots += s.Stats.Core.Shots
		res[i].Saves += s.Stats.Core.Saves
		res[i].Rating += s.Stats.Core.Rating

		if s.Winner {
			res[i].Wins++
		}
	}

	sort.Slice(res, func(i,j int) bool {
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
