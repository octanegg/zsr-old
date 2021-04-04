package stats

import (
	"sort"

	"github.com/octanegg/zsr/octane"
	"github.com/octanegg/zsr/util"
)

func EventParticipants(matches []interface{}, stages []int) []*octane.Participant {
	participantsMap := make(map[string]*octane.Participant)
	teams := []string{}
	for _, m := range matches {
		match := m.(octane.Match)
		if match.Blue != nil {
			id := match.Blue.Team.Team.ID.Hex()
			if (len(stages) == 0 || util.ContainsInt(stages, match.Stage.ID)) && !util.ContainsString(teams, id) {
				teams = append(teams, id)
			}

			if _, ok := participantsMap[id]; !ok {
				participantsMap[id] = &octane.Participant{
					Team:    match.Blue.Team.Team,
					Players: []*octane.Player{},
				}
			}
			participantsMap[id].Players = append(participantsMap[id].Players, getPlayers(participantsMap[id].Players, match.Blue.Players)...)
		}
		if match.Orange != nil {
			id := match.Orange.Team.Team.ID.Hex()
			if (len(stages) == 0 || util.ContainsInt(stages, match.Stage.ID)) && !util.ContainsString(teams, id) {
				teams = append(teams, id)
			}

			if _, ok := participantsMap[id]; !ok {
				participantsMap[id] = &octane.Participant{
					Team:    match.Orange.Team.Team,
					Players: []*octane.Player{},
				}
			}
			participantsMap[id].Players = append(participantsMap[id].Players, getPlayers(participantsMap[id].Players, match.Orange.Players)...)
		}
	}

	participants := []*octane.Participant{}
	for _, team := range teams {
		participant := participantsMap[team]
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

		if len(participant.Players) > 0 {
			participants = append(participants, participant)
		}
	}

	sort.Slice(participants, func(i, j int) bool {
		return participants[i].Team.Name < participants[j].Team.Name
	})

	return participants
}

func getPlayers(exists []*octane.Player, toAdd []*octane.PlayerStats) []*octane.Player {
	m := make(map[string]bool)
	for _, p := range exists {
		m[p.ID.Hex()] = true
	}

	var players []*octane.Player
	for _, p := range toAdd {
		if _, ok := m[p.Player.ID.Hex()]; !ok {
			players = append(players, p.Player)
		}
	}
	return players
}
