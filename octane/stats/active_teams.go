package stats

import (
	"sort"

	"github.com/octanegg/zsr/octane"
	"github.com/octanegg/zsr/util"
)

func ActiveTeams(players []interface{}, regions []string) []*octane.Participant {
	participantsMap := make(map[string]*octane.Participant)
	for _, p := range players {
		player := p.(octane.Player)
		id := player.Team.ID.Hex()
		if _, ok := participantsMap[id]; !ok {
			participantsMap[id] = &octane.Participant{
				Team:    player.Team,
				Players: []*octane.Player{},
			}
		}
		participantsMap[id].Players = append(participantsMap[id].Players, &player)
	}

	participants := []*octane.Participant{}
	for _, participant := range participantsMap {
		if len(regions) == 0 || util.ContainsString(regions, participant.Team.Region) {
			sort.Slice(participant.Players, func(i, j int) bool {
				a := participant.Players[i]
				b := participant.Players[j]

				if a.Coach && b.Coach {
					return a.Tag < b.Tag
				}

				if a.Coach || b.Coach {
					return !a.Coach
				}

				if a.Substitute && b.Substitute {
					return a.Tag < b.Tag
				}

				if a.Substitute || b.Substitute {
					return !a.Substitute
				}

				return a.Tag < b.Tag
			})
			participants = append(participants, participant)
		}
	}

	sort.Slice(participants, func(i, j int) bool {
		return participants[i].Team.Name < participants[j].Team.Name
	})

	return participants
}
