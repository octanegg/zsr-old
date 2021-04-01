package stats

import (
	"sort"

	"github.com/octanegg/zsr/octane"
)

func GetEventParticipants(matches []interface{}, stages []int) []*octane.Participant {
	participantsMap := make(map[string]*octane.Participant)
	teams := []string{}
	for _, m := range matches {
		match := m.(octane.Match)
		if match.Blue != nil {
			id := match.Blue.Team.Team.ID.Hex()
			if (len(stages) == 0 || containsInt(stages, match.Stage.ID)) && !containsString(teams, id) {
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
			if (len(stages) == 0 || containsInt(stages, match.Stage.ID)) && !containsString(teams, id) {
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

func containsInt(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func containsString(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}