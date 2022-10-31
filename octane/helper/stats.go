package helper

import (
	"fmt"
	"math"

	"github.com/octanegg/zsr/ballchasing"
	"github.com/octanegg/zsr/octane"
	"github.com/octanegg/zsr/octane/pipelines"
)

func BallchasingToPlayerStats(b *ballchasing.PlayerStats) *octane.PlayerStats {
	return &octane.PlayerStats{
		Core: &octane.PlayerCore{
			Score:   b.Core.Score,
			Goals:   b.Core.Goals,
			Assists: b.Core.Assists,
			Saves:   b.Core.Saves,
			Shots:   b.Core.Shots,
		},
		Boost: &octane.PlayerBoost{
			Bpm:                       b.Boost.Bpm,
			Bcpm:                      b.Boost.Bcpm,
			AvgAmount:                 b.Boost.AvgAmount,
			AmountCollected:           b.Boost.AmountCollected,
			AmountStolen:              b.Boost.AmountStolen,
			AmountCollectedBig:        b.Boost.AmountCollectedBig,
			AmountStolenBig:           b.Boost.AmountStolenBig,
			AmountCollectedSmall:      b.Boost.AmountCollectedSmall,
			AmountStolenSmall:         b.Boost.AmountStolenSmall,
			CountCollectedBig:         b.Boost.CountCollectedBig,
			CountStolenBig:            b.Boost.CountStolenBig,
			CountCollectedSmall:       b.Boost.CountCollectedSmall,
			CountStolenSmall:          b.Boost.CountStolenSmall,
			AmountOverfill:            b.Boost.AmountOverfill,
			AmountOverfillStolen:      b.Boost.AmountOverfillStolen,
			AmountUsedWhileSupersonic: b.Boost.AmountUsedWhileSupersonic,
			TimeZeroBoost:             b.Boost.TimeZeroBoost,
			PercentZeroBoost:          b.Boost.PercentZeroBoost,
			TimeFullBoost:             b.Boost.TimeFullBoost,
			PercentFullBoost:          b.Boost.PercentFullBoost,
			TimeBoost025:              b.Boost.TimeBoost025,
			PercentBoost025:           b.Boost.PercentBoost025,
			TimeBoost2550:             b.Boost.TimeBoost2550,
			PercentBoost2550:          b.Boost.PercentBoost2550,
			TimeBoost5075:             b.Boost.TimeBoost5075,
			PercentBoost5075:          b.Boost.PercentBoost5075,
			TimeBoost75100:            b.Boost.TimeBoost75100,
			PercentBoost75100:         b.Boost.PercentBoost75100,
		},
		Movement: &octane.PlayerMovement{
			AvgSpeed:               b.Movement.AvgSpeed,
			TotalDistance:          b.Movement.TotalDistance,
			TimeSupersonicSpeed:    b.Movement.TimeSupersonicSpeed,
			TimeBoostSpeed:         b.Movement.TimeBoostSpeed,
			TimeSlowSpeed:          b.Movement.TimeSlowSpeed,
			TimeGround:             b.Movement.TimeGround,
			TimeLowAir:             b.Movement.TimeLowAir,
			TimeHighAir:            b.Movement.TimeHighAir,
			TimePowerslide:         b.Movement.TimePowerslide,
			CountPowerslide:        b.Movement.CountPowerslide,
			AvgPowerslideDuration:  b.Movement.AvgPowerslideDuration,
			AvgSpeedPercentage:     b.Movement.AvgSpeedPercentage,
			PercentSlowSpeed:       b.Movement.PercentSlowSpeed,
			PercentBoostSpeed:      b.Movement.PercentBoostSpeed,
			PercentSupersonicSpeed: b.Movement.PercentSupersonicSpeed,
			PercentGround:          b.Movement.PercentGround,
			PercentLowAir:          b.Movement.PercentLowAir,
			PercentHighAir:         b.Movement.PercentHighAir,
		},
		Positioning: &octane.PlayerPositioning{
			AvgDistanceToBall:             b.Positioning.AvgDistanceToBall,
			AvgDistanceToBallPossession:   b.Positioning.AvgDistanceToBallPossession,
			AvgDistanceToBallNoPossession: b.Positioning.AvgDistanceToBallNoPossession,
			AvgDistanceToMates:            b.Positioning.AvgDistanceToMates,
			TimeDefensiveThird:            b.Positioning.TimeDefensiveThird,
			TimeNeutralThird:              b.Positioning.TimeNeutralThird,
			TimeOffensiveThird:            b.Positioning.TimeOffensiveThird,
			TimeDefensiveHalf:             b.Positioning.TimeDefensiveHalf,
			TimeOffensiveHalf:             b.Positioning.TimeOffensiveHalf,
			TimeBehindBall:                b.Positioning.TimeBehindBall,
			TimeInfrontBall:               b.Positioning.TimeInfrontBall,
			TimeMostBack:                  b.Positioning.TimeMostBack,
			TimeMostForward:               b.Positioning.TimeMostForward,
			GoalsAgainstWhileLastDefender: b.Positioning.GoalsAgainstWhileLastDefender,
			TimeClosestToBall:             b.Positioning.TimeClosestToBall,
			TimeFarthestFromBall:          b.Positioning.TimeFarthestFromBall,
			PercentDefensiveThird:         b.Positioning.PercentDefensiveThird,
			PercentOffensiveThird:         b.Positioning.PercentOffensiveThird,
			PercentNeutralThird:           b.Positioning.PercentNeutralThird,
			PercentDefensiveHalf:          b.Positioning.PercentDefensiveHalf,
			PercentOffensiveHalf:          b.Positioning.PercentOffensiveHalf,
			PercentBehindBall:             b.Positioning.PercentBehindBall,
			PercentInfrontBall:            b.Positioning.PercentInfrontBall,
			PercentMostBack:               b.Positioning.PercentMostBack,
			PercentMostForward:            b.Positioning.PercentMostForward,
			PercentClosestToBall:          b.Positioning.PercentClosestToBall,
			PercentFarthestFromBall:       b.Positioning.PercentFarthestFromBall,
		},
		Demolitions: &octane.PlayerDemolitions{
			Inflicted: b.Demolitions.Inflicted,
			Taken:     b.Demolitions.Taken,
		},
	}
}

func PlayerStatsToTeamStats(players []*octane.PlayerInfo) *octane.TeamStats {
	team := &octane.TeamStats{
		Core: &octane.TeamCore{},
	}

	for _, player := range players {
		team.Core.Score += player.Stats.Core.Score
		team.Core.Goals += player.Stats.Core.Goals
		team.Core.Assists += player.Stats.Core.Assists
		team.Core.Saves += player.Stats.Core.Saves
		team.Core.Shots += player.Stats.Core.Shots

		if player.Stats.Boost != nil {
			if team.Boost == nil {
				team.Boost = &octane.TeamBoost{}
			}
			team.Boost.Bpm += player.Stats.Boost.Bpm
			team.Boost.Bcpm += player.Stats.Boost.Bcpm
			team.Boost.AvgAmount += player.Stats.Boost.AvgAmount
			team.Boost.AmountCollected += player.Stats.Boost.AmountCollected
			team.Boost.AmountStolen += player.Stats.Boost.AmountStolen
			team.Boost.AmountCollectedBig += player.Stats.Boost.AmountCollectedBig
			team.Boost.AmountStolenBig += player.Stats.Boost.AmountStolenBig
			team.Boost.AmountCollectedSmall += player.Stats.Boost.AmountCollectedSmall
			team.Boost.AmountStolenSmall += player.Stats.Boost.AmountStolenSmall
			team.Boost.CountCollectedBig += player.Stats.Boost.CountCollectedBig
			team.Boost.CountStolenBig += player.Stats.Boost.CountStolenBig
			team.Boost.CountCollectedSmall += player.Stats.Boost.CountCollectedSmall
			team.Boost.CountStolenSmall += player.Stats.Boost.CountStolenSmall
			team.Boost.AmountOverfill += player.Stats.Boost.AmountOverfill
			team.Boost.AmountOverfillStolen += player.Stats.Boost.AmountOverfillStolen
			team.Boost.AmountUsedWhileSupersonic += player.Stats.Boost.AmountUsedWhileSupersonic
			team.Boost.TimeZeroBoost += player.Stats.Boost.TimeZeroBoost
			team.Boost.TimeFullBoost += player.Stats.Boost.TimeFullBoost
			team.Boost.TimeBoost025 += player.Stats.Boost.TimeBoost025
			team.Boost.TimeBoost2550 += player.Stats.Boost.TimeBoost2550
			team.Boost.TimeBoost5075 += player.Stats.Boost.TimeBoost5075
			team.Boost.TimeBoost75100 += player.Stats.Boost.TimeBoost75100
		}

		if player.Stats.Movement != nil {
			if team.Movement == nil {
				team.Movement = &octane.TeamMovement{}
			}
			team.Movement.TotalDistance += player.Stats.Movement.TotalDistance
			team.Movement.TimeSupersonicSpeed += player.Stats.Movement.TimeSupersonicSpeed
			team.Movement.TimeBoostSpeed += player.Stats.Movement.TimeBoostSpeed
			team.Movement.TimeSlowSpeed += player.Stats.Movement.TimeSlowSpeed
			team.Movement.TimeGround += player.Stats.Movement.TimeGround
			team.Movement.TimeLowAir += player.Stats.Movement.TimeLowAir
			team.Movement.TimeHighAir += player.Stats.Movement.TimeHighAir
			team.Movement.TimePowerslide += player.Stats.Movement.TimePowerslide
			team.Movement.CountPowerslide += player.Stats.Movement.CountPowerslide
		}

		if player.Stats.Positioning != nil {
			if team.Positioning == nil {
				team.Positioning = &octane.TeamPositioning{}
			}
			team.Positioning.TimeDefensiveThird += player.Stats.Positioning.TimeDefensiveThird
			team.Positioning.TimeNeutralThird += player.Stats.Positioning.TimeNeutralThird
			team.Positioning.TimeOffensiveThird += player.Stats.Positioning.TimeOffensiveThird
			team.Positioning.TimeDefensiveHalf += player.Stats.Positioning.TimeDefensiveHalf
			team.Positioning.TimeOffensiveHalf += player.Stats.Positioning.TimeOffensiveHalf
			team.Positioning.TimeBehindBall += player.Stats.Positioning.TimeBehindBall
			team.Positioning.TimeInfrontBall += player.Stats.Positioning.TimeInfrontBall
		}

		if player.Stats.Demolitions != nil {
			if team.Demolitions == nil {
				team.Demolitions = &octane.TeamDemolitions{}
			}
			team.Demolitions.Inflicted += player.Stats.Demolitions.Inflicted
			team.Demolitions.Taken += player.Stats.Demolitions.Taken
		}
	}

	if team.Core.Shots > 0 {
		team.Core.ShootingPercentage = float64(team.Core.Goals) / float64(team.Core.Shots) * 100
	}

	return team
}

func PlayerStatsToPlayer(playerStats []*octane.PlayerInfo) []*octane.Player {
	var players []*octane.Player
	for _, stats := range playerStats {
		players = append(players, stats.Player)
	}
	return players
}

func GameToStatlines(game *octane.Game) ([]*octane.Statline, []*octane.Statline) {
	var blue, orange []*octane.Statline

	for _, p := range game.Blue.Players {
		blue = append(blue, &octane.Statline{
			Game: &octane.Game{
				ID:            game.ID,
				Number:        game.Number,
				Match:         game.Match,
				Date:          game.Date,
				Map:           game.Map,
				Duration:      game.Duration,
				Overtime:      game.Overtime,
				BallchasingID: game.BallchasingID,
			},
			Team: &octane.StatlineSide{
				Score:       game.Blue.Team.Stats.Core.Goals,
				Winner:      game.Blue.Winner,
				MatchWinner: game.Blue.MatchWinner,
				Team:        game.Blue.Team.Team,
				Stats:       game.Blue.Team.Stats,
				Players:     PlayerStatsToPlayer(game.Blue.Players),
			},
			Opponent: &octane.StatlineSide{
				Score:       game.Orange.Team.Stats.Core.Goals,
				Winner:      game.Orange.Winner,
				MatchWinner: game.Orange.MatchWinner,
				Team:        game.Orange.Team.Team,
				Stats:       game.Orange.Team.Stats,
				Players:     PlayerStatsToPlayer(game.Orange.Players),
			},
			Player: p,
		})
	}

	for _, p := range game.Orange.Players {
		orange = append(orange, &octane.Statline{
			Game: &octane.Game{
				ID:            game.ID,
				Number:        game.Number,
				Match:         game.Match,
				Date:          game.Date,
				Map:           game.Map,
				Duration:      game.Duration,
				Overtime:      game.Overtime,
				BallchasingID: game.BallchasingID,
			},
			Team: &octane.StatlineSide{
				Score:       game.Orange.Team.Stats.Core.Goals,
				Winner:      game.Orange.Winner,
				MatchWinner: game.Orange.MatchWinner,
				Team:        game.Orange.Team.Team,
				Stats:       game.Orange.Team.Stats,
				Players:     PlayerStatsToPlayer(game.Orange.Players),
			},
			Opponent: &octane.StatlineSide{
				Score:       game.Blue.Team.Stats.Core.Goals,
				Winner:      game.Blue.Winner,
				MatchWinner: game.Blue.MatchWinner,
				Team:        game.Blue.Team.Team,
				Stats:       game.Blue.Team.Stats,
				Players:     PlayerStatsToPlayer(game.Blue.Players),
			},
			Player: p,
		})
	}

	return blue, orange
}

func StatlinesToAggregatePlayerStats(statlines []*octane.Statline) []*octane.PlayerInfo {
	var teamGoals int
	playerMap := map[string]*octane.PlayerInfo{}
	for _, statline := range statlines {
		player := statline.Player.Player.ID.Hex()

		if _, ok := playerMap[player]; !ok {
			playerMap[player] = &octane.PlayerInfo{
				Player: statline.Player.Player,
				Stats: &octane.PlayerStats{
					Core: &octane.PlayerCore{},
				},
				Advanced: &octane.AdvancedStats{},
			}
		}

		playerMap[player].Stats.Core.Score += statline.Player.Stats.Core.Score
		playerMap[player].Stats.Core.Goals += statline.Player.Stats.Core.Goals
		playerMap[player].Stats.Core.Assists += statline.Player.Stats.Core.Assists
		playerMap[player].Stats.Core.Saves += statline.Player.Stats.Core.Saves
		playerMap[player].Stats.Core.Shots += statline.Player.Stats.Core.Shots

		if statline.Player.Stats.Boost != nil {
			if playerMap[player].Stats.Boost == nil {
				playerMap[player].Stats.Boost = &octane.PlayerBoost{}
			}
			playerMap[player].Stats.Boost.Bpm += statline.Player.Stats.Boost.Bpm
			playerMap[player].Stats.Boost.Bcpm += statline.Player.Stats.Boost.Bcpm
			playerMap[player].Stats.Boost.AvgAmount += statline.Player.Stats.Boost.AvgAmount
			playerMap[player].Stats.Boost.AmountCollected += statline.Player.Stats.Boost.AmountCollected
			playerMap[player].Stats.Boost.AmountStolen += statline.Player.Stats.Boost.AmountStolen
			playerMap[player].Stats.Boost.AmountCollectedBig += statline.Player.Stats.Boost.AmountCollectedBig
			playerMap[player].Stats.Boost.AmountStolenBig += statline.Player.Stats.Boost.AmountStolenBig
			playerMap[player].Stats.Boost.AmountCollectedSmall += statline.Player.Stats.Boost.AmountCollectedSmall
			playerMap[player].Stats.Boost.AmountStolenSmall += statline.Player.Stats.Boost.AmountStolenSmall
			playerMap[player].Stats.Boost.CountCollectedBig += statline.Player.Stats.Boost.CountCollectedBig
			playerMap[player].Stats.Boost.CountStolenBig += statline.Player.Stats.Boost.CountStolenBig
			playerMap[player].Stats.Boost.CountCollectedSmall += statline.Player.Stats.Boost.CountCollectedSmall
			playerMap[player].Stats.Boost.CountStolenSmall += statline.Player.Stats.Boost.CountStolenSmall
			playerMap[player].Stats.Boost.AmountOverfill += statline.Player.Stats.Boost.AmountOverfill
			playerMap[player].Stats.Boost.AmountOverfillStolen += statline.Player.Stats.Boost.AmountOverfillStolen
			playerMap[player].Stats.Boost.AmountUsedWhileSupersonic += statline.Player.Stats.Boost.AmountUsedWhileSupersonic
			playerMap[player].Stats.Boost.TimeZeroBoost += statline.Player.Stats.Boost.TimeZeroBoost
			playerMap[player].Stats.Boost.PercentZeroBoost += statline.Player.Stats.Boost.PercentZeroBoost
			playerMap[player].Stats.Boost.TimeFullBoost += statline.Player.Stats.Boost.TimeFullBoost
			playerMap[player].Stats.Boost.PercentFullBoost += statline.Player.Stats.Boost.PercentFullBoost
			playerMap[player].Stats.Boost.TimeBoost025 += statline.Player.Stats.Boost.TimeBoost025
			playerMap[player].Stats.Boost.TimeBoost2550 += statline.Player.Stats.Boost.TimeBoost2550
			playerMap[player].Stats.Boost.TimeBoost5075 += statline.Player.Stats.Boost.TimeBoost5075
			playerMap[player].Stats.Boost.TimeBoost75100 += statline.Player.Stats.Boost.TimeBoost75100
			playerMap[player].Stats.Boost.PercentBoost025 += statline.Player.Stats.Boost.PercentBoost025
			playerMap[player].Stats.Boost.PercentBoost2550 += statline.Player.Stats.Boost.PercentBoost2550
			playerMap[player].Stats.Boost.PercentBoost5075 += statline.Player.Stats.Boost.PercentBoost5075
			playerMap[player].Stats.Boost.PercentBoost75100 += statline.Player.Stats.Boost.PercentBoost75100
		}

		if statline.Player.Stats.Movement != nil {
			if playerMap[player].Stats.Movement == nil {
				playerMap[player].Stats.Movement = &octane.PlayerMovement{}
			}
			playerMap[player].Stats.Movement.AvgSpeed += statline.Player.Stats.Movement.AvgSpeed
			playerMap[player].Stats.Movement.TotalDistance += statline.Player.Stats.Movement.TotalDistance
			playerMap[player].Stats.Movement.TimeSupersonicSpeed += statline.Player.Stats.Movement.TimeSupersonicSpeed
			playerMap[player].Stats.Movement.TimeBoostSpeed += statline.Player.Stats.Movement.TimeBoostSpeed
			playerMap[player].Stats.Movement.TimeSlowSpeed += statline.Player.Stats.Movement.TimeSlowSpeed
			playerMap[player].Stats.Movement.TimeGround += statline.Player.Stats.Movement.TimeGround
			playerMap[player].Stats.Movement.TimeLowAir += statline.Player.Stats.Movement.TimeLowAir
			playerMap[player].Stats.Movement.TimeHighAir += statline.Player.Stats.Movement.TimeHighAir
			playerMap[player].Stats.Movement.TimePowerslide += statline.Player.Stats.Movement.TimePowerslide
			playerMap[player].Stats.Movement.CountPowerslide += statline.Player.Stats.Movement.CountPowerslide
			playerMap[player].Stats.Movement.AvgPowerslideDuration += statline.Player.Stats.Movement.AvgPowerslideDuration
			playerMap[player].Stats.Movement.AvgSpeedPercentage += statline.Player.Stats.Movement.AvgSpeedPercentage
			playerMap[player].Stats.Movement.PercentSlowSpeed += statline.Player.Stats.Movement.PercentSlowSpeed
			playerMap[player].Stats.Movement.PercentBoostSpeed += statline.Player.Stats.Movement.PercentBoostSpeed
			playerMap[player].Stats.Movement.PercentSupersonicSpeed += statline.Player.Stats.Movement.PercentSupersonicSpeed
			playerMap[player].Stats.Movement.PercentGround += statline.Player.Stats.Movement.PercentGround
			playerMap[player].Stats.Movement.PercentLowAir += statline.Player.Stats.Movement.PercentLowAir
			playerMap[player].Stats.Movement.PercentHighAir += statline.Player.Stats.Movement.PercentHighAir
		}

		if statline.Player.Stats.Positioning != nil {
			if playerMap[player].Stats.Positioning == nil {
				playerMap[player].Stats.Positioning = &octane.PlayerPositioning{}
			}
			playerMap[player].Stats.Positioning.AvgDistanceToBall += statline.Player.Stats.Positioning.AvgDistanceToBall
			playerMap[player].Stats.Positioning.AvgDistanceToBallPossession += statline.Player.Stats.Positioning.AvgDistanceToBallPossession
			playerMap[player].Stats.Positioning.AvgDistanceToBallNoPossession += statline.Player.Stats.Positioning.AvgDistanceToBallNoPossession
			playerMap[player].Stats.Positioning.AvgDistanceToMates += statline.Player.Stats.Positioning.AvgDistanceToMates
			playerMap[player].Stats.Positioning.TimeDefensiveThird += statline.Player.Stats.Positioning.TimeDefensiveThird
			playerMap[player].Stats.Positioning.TimeNeutralThird += statline.Player.Stats.Positioning.TimeNeutralThird
			playerMap[player].Stats.Positioning.TimeOffensiveThird += statline.Player.Stats.Positioning.TimeOffensiveThird
			playerMap[player].Stats.Positioning.TimeDefensiveHalf += statline.Player.Stats.Positioning.TimeDefensiveHalf
			playerMap[player].Stats.Positioning.TimeOffensiveHalf += statline.Player.Stats.Positioning.TimeOffensiveHalf
			playerMap[player].Stats.Positioning.TimeBehindBall += statline.Player.Stats.Positioning.TimeBehindBall
			playerMap[player].Stats.Positioning.TimeInfrontBall += statline.Player.Stats.Positioning.TimeInfrontBall
			playerMap[player].Stats.Positioning.TimeMostBack += statline.Player.Stats.Positioning.TimeMostBack
			playerMap[player].Stats.Positioning.TimeMostForward += statline.Player.Stats.Positioning.TimeMostForward
			playerMap[player].Stats.Positioning.GoalsAgainstWhileLastDefender += statline.Player.Stats.Positioning.GoalsAgainstWhileLastDefender
			playerMap[player].Stats.Positioning.TimeClosestToBall += statline.Player.Stats.Positioning.TimeClosestToBall
			playerMap[player].Stats.Positioning.TimeFarthestFromBall += statline.Player.Stats.Positioning.TimeFarthestFromBall
			playerMap[player].Stats.Positioning.PercentDefensiveThird += statline.Player.Stats.Positioning.PercentDefensiveThird
			playerMap[player].Stats.Positioning.PercentOffensiveThird += statline.Player.Stats.Positioning.PercentOffensiveThird
			playerMap[player].Stats.Positioning.PercentNeutralThird += statline.Player.Stats.Positioning.PercentNeutralThird
			playerMap[player].Stats.Positioning.PercentDefensiveHalf += statline.Player.Stats.Positioning.PercentDefensiveHalf
			playerMap[player].Stats.Positioning.PercentOffensiveHalf += statline.Player.Stats.Positioning.PercentOffensiveHalf
			playerMap[player].Stats.Positioning.PercentBehindBall += statline.Player.Stats.Positioning.PercentBehindBall
			playerMap[player].Stats.Positioning.PercentInfrontBall += statline.Player.Stats.Positioning.PercentInfrontBall
			playerMap[player].Stats.Positioning.PercentMostBack += statline.Player.Stats.Positioning.PercentMostBack
			playerMap[player].Stats.Positioning.PercentMostForward += statline.Player.Stats.Positioning.PercentMostForward
			playerMap[player].Stats.Positioning.PercentClosestToBall += statline.Player.Stats.Positioning.PercentClosestToBall
			playerMap[player].Stats.Positioning.PercentFarthestFromBall += statline.Player.Stats.Positioning.PercentFarthestFromBall
		}

		if statline.Player.Stats.Demolitions != nil {
			if playerMap[player].Stats.Demolitions == nil {
				playerMap[player].Stats.Demolitions = &octane.PlayerDemolitions{}
			}
			playerMap[player].Stats.Demolitions.Inflicted += statline.Player.Stats.Demolitions.Inflicted
			playerMap[player].Stats.Demolitions.Taken += statline.Player.Stats.Demolitions.Taken
		}

		playerMap[player].Advanced.Rating += statline.Player.Advanced.Rating

		teamGoals += int(statline.Player.Stats.Core.Goals)
	}

	games := float64(len(statlines)) / float64(len(playerMap))

	var players []*octane.PlayerInfo
	for _, player := range playerMap {
		if player.Stats.Boost != nil {
			player.Stats.Boost.PercentZeroBoost /= games
			player.Stats.Boost.PercentFullBoost /= games
			player.Stats.Boost.PercentBoost025 /= games
			player.Stats.Boost.PercentBoost2550 /= games
			player.Stats.Boost.PercentBoost5075 /= games
			player.Stats.Boost.PercentBoost75100 /= games
		}

		if player.Stats.Movement != nil {
			player.Stats.Movement.AvgSpeedPercentage /= games
			player.Stats.Movement.PercentSlowSpeed /= games
			player.Stats.Movement.PercentBoostSpeed /= games
			player.Stats.Movement.PercentSupersonicSpeed /= games
			player.Stats.Movement.PercentGround /= games
			player.Stats.Movement.PercentLowAir /= games
			player.Stats.Movement.PercentHighAir /= games
		}

		if player.Stats.Positioning != nil {
			player.Stats.Positioning.PercentDefensiveThird /= games
			player.Stats.Positioning.PercentOffensiveThird /= games
			player.Stats.Positioning.PercentNeutralThird /= games
			player.Stats.Positioning.PercentDefensiveHalf /= games
			player.Stats.Positioning.PercentOffensiveHalf /= games
			player.Stats.Positioning.PercentBehindBall /= games
			player.Stats.Positioning.PercentInfrontBall /= games
			player.Stats.Positioning.PercentMostBack /= games
			player.Stats.Positioning.PercentMostForward /= games
			player.Stats.Positioning.PercentClosestToBall /= games
			player.Stats.Positioning.PercentFarthestFromBall /= games
		}

		if player.Stats.Core.Shots > 0 {
			player.Stats.Core.ShootingPercentage = player.Stats.Core.Goals / player.Stats.Core.Shots * 100
		}

		if teamGoals > 0 {
			player.Advanced.GoalParticipation = float64(player.Stats.Core.Goals+player.Stats.Core.Assists) / float64(teamGoals) * 100
		}

		player.Advanced.Rating /= games
		players = append(players, player)
	}

	return players
}

func ReverseSweep(games []*octane.Game) (bool, bool) {
	if games == nil || len(games) == 0 || games[0].Match.Format == nil {
		return false, false
	}

	format := games[0].Match.Format.Length

	if len(games) != format {
		return false, false
	}

	for i := 1; i < format/2; i++ {
		if games[i].Blue.Winner != games[i-1].Blue.Winner {
			return false, false
		}
	}

	for i := format/2 + 1; i < format-1; i++ {
		if games[i].Blue.Winner != games[i-1].Blue.Winner {
			return false, false
		}
	}

	return true, games[0].Blue.Winner != games[format-1].Blue.Winner
}

func GamesToGameOverviews(games []*octane.Game) []*octane.GameOverview {
	var gameOverviews []*octane.GameOverview
	for _, game := range games {
		for game.Number > len(gameOverviews)+1 {
			gameOverviews = append(gameOverviews, &octane.GameOverview{})
		}
		gameOverviews = append(gameOverviews, &octane.GameOverview{
			ID:            game.ID,
			Blue:          game.Blue.Team.Stats.Core.Goals,
			Orange:        game.Orange.Team.Stats.Core.Goals,
			Duration:      game.Duration,
			Overtime:      game.Overtime,
			BallchasingID: game.BallchasingID,
		})
	}
	return gameOverviews
}

func AverageScore(client octane.Client, core *octane.PlayerCore) (float64, error) {
	pipeline := pipelines.AverageScore(int(core.Goals), int(core.Assists), int(core.Saves), int(core.Shots))

	data, err := client.Statlines().Pipeline(pipeline.Pipeline, pipeline.Decode)
	if err != nil {
		return 0, err
	}

	if len(data) == 0 {
		return 0, nil
	}

	s := data[0].(struct {
		Score float64 `json:"score" bson:"score"`
	})

	return math.Floor(s.Score), nil
}

func Rating(stats *octane.PlayerInfo) float64 {
	rating := float64(0)
	rating += float64(stats.Stats.Core.Score) / 369.8394212252121
	rating += float64(stats.Stats.Core.Goals) / 0.6616665799198459
	rating += float64(stats.Stats.Core.Assists) / 0.5248321449018893
	rating += float64(stats.Stats.Core.Saves) / 1.5932962056940614
	rating += float64(stats.Stats.Core.Shots) / 2.7166189559152656
	rating += float64(stats.Stats.Core.ShootingPercentage) / 24.35625277807581
	rating += float64(stats.Advanced.GoalParticipation) / 59.77324334387406

	return rating / 7
}

func StatlinesToRecords(statlines []*octane.Statline) []*octane.Record {
	var records []*octane.Record
	playerRecords := map[string]map[string]map[string]float64{}
	teamRecords := map[string]map[string]map[string]float64{}

	for _, statline := range statlines {
		stats := statline.Player.Stats

		if playerRecords[statline.Player.Player.ID.Hex()] == nil {
			playerRecords[statline.Player.Player.ID.Hex()] = map[string]map[string]float64{}
		}

		if playerRecords[statline.Player.Player.ID.Hex()][statline.Game.Match.ID.Hex()] == nil {
			playerRecords[statline.Player.Player.ID.Hex()][statline.Game.Match.ID.Hex()] = map[string]float64{}
		}

		if teamRecords[statline.Team.Team.ID.Hex()] == nil {
			teamRecords[statline.Team.Team.ID.Hex()] = map[string]map[string]float64{}
		}

		if teamRecords[statline.Team.Team.ID.Hex()][statline.Game.Match.ID.Hex()] == nil {
			teamRecords[statline.Team.Team.ID.Hex()][statline.Game.Match.ID.Hex()] = map[string]float64{}
		}

		if teamRecords[statline.Team.Team.ID.Hex()][statline.Game.Match.ID.Hex()+fmt.Sprint(statline.Game.Number)] == nil {
			teamRecords[statline.Team.Team.ID.Hex()][statline.Game.Match.ID.Hex()+fmt.Sprint(statline.Game.Number)] = map[string]float64{}
		}

		playerRecords[statline.Player.Player.ID.Hex()][statline.Game.Match.ID.Hex()]["games"] += 1
		playerRecords[statline.Player.Player.ID.Hex()][statline.Game.Match.ID.Hex()]["duration"] += float64(statline.Game.Duration)

		createRecord := func(stat string, value float64) *octane.Record {
			teamRecords[statline.Team.Team.ID.Hex()][statline.Game.Match.ID.Hex()][stat] += value
			teamRecords[statline.Team.Team.ID.Hex()][statline.Game.Match.ID.Hex()+fmt.Sprint(statline.Game.Number)][stat] += value
			playerRecords[statline.Player.Player.ID.Hex()][statline.Game.Match.ID.Hex()][stat] += value
			return &octane.Record{
				Game: &octane.Game{
					ID:     statline.Game.ID,
					Number: statline.Game.Number,
					Match: &octane.Match{
						ID:                  statline.Game.Match.ID,
						Slug:                statline.Game.Match.Slug,
						Event:               statline.Game.Match.Event,
						Stage:               statline.Game.Match.Stage,
						Substage:            statline.Game.Match.Substage,
						Date:                statline.Game.Match.Date,
						Format:              statline.Game.Match.Format,
						Blue:                statline.Game.Match.Blue,
						Orange:              statline.Game.Match.Orange,
						Number:              statline.Game.Match.Number,
						ReverseSweep:        statline.Game.Match.ReverseSweep,
						ReverseSweepAttempt: statline.Game.Match.ReverseSweepAttempt,
					},
					Map:      statline.Game.Map,
					Duration: statline.Game.Duration,
					Overtime: statline.Game.Overtime,
					Date:     statline.Game.Date,
				},
				Team: &octane.RecordSide{
					Score:       statline.Team.Score,
					Team:        statline.Team.Team,
					Winner:      statline.Team.Winner,
					MatchWinner: statline.Team.MatchWinner,
					Players:     statline.Team.Players,
				},
				Opponent: &octane.RecordSide{
					Score:       statline.Opponent.Score,
					Team:        statline.Opponent.Team,
					Winner:      statline.Opponent.Winner,
					MatchWinner: statline.Opponent.MatchWinner,
					Players:     statline.Opponent.Players,
				},
				Player:      statline.Player.Player,
				Stat:        stat,
				PlayerValue: value,
			}
		}

		records = append(records, createRecord("score", stats.Core.Score))
		records = append(records, createRecord("goals", stats.Core.Goals))
		records = append(records, createRecord("assists", stats.Core.Assists))
		records = append(records, createRecord("saves", stats.Core.Saves))
		records = append(records, createRecord("shots", stats.Core.Shots))
		records = append(records, createRecord("shooting_percentage", stats.Core.ShootingPercentage))
		records = append(records, createRecord("goal_participation", statline.Player.Advanced.GoalParticipation))
		records = append(records, createRecord("rating", statline.Player.Advanced.Rating))

		if stats.Boost != nil {
			records = append(records, createRecord("bpm", stats.Boost.Bpm))
			records = append(records, createRecord("bcpm", stats.Boost.Bcpm))
			records = append(records, createRecord("avg_amount", stats.Boost.AvgAmount))
			records = append(records, createRecord("amount_collected", stats.Boost.AmountCollected))
			records = append(records, createRecord("amount_stolen", stats.Boost.AmountStolen))
			records = append(records, createRecord("amount_collected_big", stats.Boost.AmountCollectedBig))
			records = append(records, createRecord("amount_stolen_big", stats.Boost.AmountStolenBig))
			records = append(records, createRecord("amount_collected_small", stats.Boost.AmountCollectedSmall))
			records = append(records, createRecord("amount_stolen_small", stats.Boost.AmountStolenSmall))
			records = append(records, createRecord("count_collected_big", stats.Boost.CountCollectedBig))
			records = append(records, createRecord("count_stolen_big", stats.Boost.CountStolenBig))
			records = append(records, createRecord("count_collected_small", stats.Boost.CountCollectedSmall))
			records = append(records, createRecord("count_stolen_small", stats.Boost.CountStolenSmall))
			records = append(records, createRecord("amount_overfill", stats.Boost.AmountOverfill))
			records = append(records, createRecord("amount_overfill_stolen", stats.Boost.AmountOverfillStolen))
			records = append(records, createRecord("amount_used_while_supersonic", stats.Boost.AmountUsedWhileSupersonic))
			records = append(records, createRecord("time_zero_boost", stats.Boost.TimeZeroBoost))
			records = append(records, createRecord("percent_zero_boost", stats.Boost.PercentZeroBoost))
			records = append(records, createRecord("time_full_boost", stats.Boost.TimeFullBoost))
			records = append(records, createRecord("percent_full_boost", stats.Boost.PercentFullBoost))
			records = append(records, createRecord("time_boost_0_25", stats.Boost.TimeBoost025))
			records = append(records, createRecord("time_boost_25_50", stats.Boost.TimeBoost2550))
			records = append(records, createRecord("time_boost_50_75", stats.Boost.TimeBoost5075))
			records = append(records, createRecord("time_boost_75_100", stats.Boost.TimeBoost75100))
			records = append(records, createRecord("percent_boost_0_25", stats.Boost.PercentBoost025))
			records = append(records, createRecord("percent_boost_25_50", stats.Boost.PercentBoost2550))
			records = append(records, createRecord("percent_boost_50_75", stats.Boost.PercentBoost5075))
			records = append(records, createRecord("percent_boost_75_100", stats.Boost.PercentBoost75100))
		}

		if stats.Movement != nil {
			records = append(records, createRecord("avg_speed", stats.Movement.AvgSpeed))
			records = append(records, createRecord("total_distance", stats.Movement.TotalDistance))
			records = append(records, createRecord("time_supersonic_speed", stats.Movement.TimeSupersonicSpeed))
			records = append(records, createRecord("time_boost_speed", stats.Movement.TimeBoostSpeed))
			records = append(records, createRecord("time_slow_speed", stats.Movement.TimeSlowSpeed))
			records = append(records, createRecord("time_ground", stats.Movement.TimeGround))
			records = append(records, createRecord("time_low_air", stats.Movement.TimeLowAir))
			records = append(records, createRecord("time_high_air", stats.Movement.TimeHighAir))
			records = append(records, createRecord("time_powerslide", stats.Movement.TimePowerslide))
			records = append(records, createRecord("count_powerslide", stats.Movement.CountPowerslide))
			records = append(records, createRecord("avg_powerslide_duration", stats.Movement.AvgPowerslideDuration))
			records = append(records, createRecord("avg_speed_percentage", stats.Movement.AvgSpeedPercentage))
			records = append(records, createRecord("percent_slow_speed", stats.Movement.PercentSlowSpeed))
			records = append(records, createRecord("percent_boost_speed", stats.Movement.PercentBoostSpeed))
			records = append(records, createRecord("percent_supersonic_speed", stats.Movement.PercentSupersonicSpeed))
			records = append(records, createRecord("percent_ground", stats.Movement.PercentGround))
			records = append(records, createRecord("percent_low_air", stats.Movement.PercentLowAir))
			records = append(records, createRecord("percent_high_air", stats.Movement.PercentHighAir))
		}

		if stats.Positioning != nil {
			records = append(records, createRecord("avg_distance_to_ball", stats.Positioning.AvgDistanceToBall))
			records = append(records, createRecord("avg_distance_to_ball_possession", stats.Positioning.AvgDistanceToBallPossession))
			records = append(records, createRecord("avg_distance_to_ball_no_possession", stats.Positioning.AvgDistanceToBallNoPossession))
			records = append(records, createRecord("time_defensive_third", stats.Positioning.TimeDefensiveThird))
			records = append(records, createRecord("time_neutral_third", stats.Positioning.TimeNeutralThird))
			records = append(records, createRecord("time_offensive_third", stats.Positioning.TimeOffensiveThird))
			records = append(records, createRecord("time_defensive_half", stats.Positioning.TimeDefensiveHalf))
			records = append(records, createRecord("time_offensive_half", stats.Positioning.TimeOffensiveHalf))
			records = append(records, createRecord("time_behind_ball", stats.Positioning.TimeBehindBall))
			records = append(records, createRecord("time_infront_ball", stats.Positioning.TimeInfrontBall))
			records = append(records, createRecord("time_most_back", stats.Positioning.TimeMostBack))
			records = append(records, createRecord("time_most_forward", stats.Positioning.TimeMostForward))
			records = append(records, createRecord("goals_against_while_last_defender", stats.Positioning.GoalsAgainstWhileLastDefender))
			records = append(records, createRecord("time_closest_to_ball", stats.Positioning.TimeClosestToBall))
			records = append(records, createRecord("time_farthest_from_ball", stats.Positioning.TimeFarthestFromBall))
			records = append(records, createRecord("percent_defensive_third", stats.Positioning.PercentDefensiveThird))
			records = append(records, createRecord("percent_offensive_third", stats.Positioning.PercentOffensiveThird))
			records = append(records, createRecord("percent_neutral_third", stats.Positioning.PercentNeutralThird))
			records = append(records, createRecord("percent_defensive_half", stats.Positioning.PercentDefensiveHalf))
			records = append(records, createRecord("percent_offensive_half", stats.Positioning.PercentOffensiveHalf))
			records = append(records, createRecord("percent_behind_ball", stats.Positioning.PercentBehindBall))
			records = append(records, createRecord("percent_infront_ball", stats.Positioning.PercentInfrontBall))
			records = append(records, createRecord("percent_most_back", stats.Positioning.PercentMostBack))
			records = append(records, createRecord("percent_most_forward", stats.Positioning.PercentMostForward))
			records = append(records, createRecord("percent_closest_to_ball", stats.Positioning.PercentClosestToBall))
			records = append(records, createRecord("percent_farthest_from_ball", stats.Positioning.PercentFarthestFromBall))
		}

		if stats.Demolitions != nil {
			records = append(records, createRecord("inflicted", stats.Demolitions.Inflicted))
			records = append(records, createRecord("taken", stats.Demolitions.Taken))
		}
	}

	for _, record := range records {
		record.Duration = playerRecords[record.Player.ID.Hex()][record.Game.Match.ID.Hex()]["duration"]
		record.PlayerMatchValue = playerRecords[record.Player.ID.Hex()][record.Game.Match.ID.Hex()][record.Stat]
		record.PlayerMatchAverage = record.PlayerMatchValue / playerRecords[record.Player.ID.Hex()][record.Game.Match.ID.Hex()]["games"]
		record.TeamValue = teamRecords[record.Team.Team.ID.Hex()][record.Game.Match.ID.Hex()+fmt.Sprint(record.Game.Number)][record.Stat]
		record.TeamMatchValue = teamRecords[record.Team.Team.ID.Hex()][record.Game.Match.ID.Hex()][record.Stat]
		record.TeamMatchAverage = record.TeamMatchValue / playerRecords[record.Player.ID.Hex()][record.Game.Match.ID.Hex()]["games"]
		record.GameValue = teamRecords[record.Team.Team.ID.Hex()][record.Game.Match.ID.Hex()+fmt.Sprint(record.Game.Number)][record.Stat] + teamRecords[record.Opponent.Team.ID.Hex()][record.Game.Match.ID.Hex()+fmt.Sprint(record.Game.Number)][record.Stat]
		record.GameDifferential = teamRecords[record.Team.Team.ID.Hex()][record.Game.Match.ID.Hex()+fmt.Sprint(record.Game.Number)][record.Stat] - teamRecords[record.Opponent.Team.ID.Hex()][record.Game.Match.ID.Hex()+fmt.Sprint(record.Game.Number)][record.Stat]
		record.GameAverage = record.GameValue / playerRecords[record.Player.ID.Hex()][record.Game.Match.ID.Hex()]["games"]
		record.MatchValue = teamRecords[record.Team.Team.ID.Hex()][record.Game.Match.ID.Hex()][record.Stat] + teamRecords[record.Opponent.Team.ID.Hex()][record.Game.Match.ID.Hex()+fmt.Sprint(record.Game.Number)][record.Stat]
		record.MatchDifferential = teamRecords[record.Team.Team.ID.Hex()][record.Game.Match.ID.Hex()][record.Stat] - teamRecords[record.Opponent.Team.ID.Hex()][record.Game.Match.ID.Hex()+fmt.Sprint(record.Game.Number)][record.Stat]
	}

	return records
}
