package helper

import (
	"errors"
	"fmt"
	"os"

	"github.com/octanegg/zsr/ballchasing"
	"github.com/octanegg/zsr/internal/config"
	"github.com/octanegg/zsr/octane"
	"github.com/octanegg/zsr/octane/pipelines"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func UseBallchasing(client octane.Client, game *octane.Game) (*octane.Game, error) {
	b := ballchasing.New(os.Getenv(config.EnvBallchasing))
	replay, err := b.GetReplay(game.BallchasingID)
	if err != nil {
		return nil, err
	} else if game.FlipBallchasing {
		replay.Blue, replay.Orange = replay.Orange, replay.Blue
	}

	game.BallchasingID = replay.ID
	game.Duration = 300 + replay.OvertimeSeconds
	game.Date = &replay.Date
	game.Map = &octane.Map{
		ID:   replay.MapCode,
		Name: replay.MapName,
	}

	// if len(replay.Blue.Players) > game.Match.Event.Mode {
	// 	sort.SliceStable(replay.Blue.Players, func(i, j int) bool {
	// 		return replay.Blue.Players[i].Stats.Core.Score > replay.Blue.Players[j].Stats.Core.Score
	// 	})
	// 	if len(replay.Orange.Players) >= game.Match.Event.Mode {
	// 		replay.Blue.Players = replay.Blue.Players[:game.Match.Event.Mode]
	// 	} else {
	// 		replay.Blue.Players = replay.Blue.Players[:game.Match.Event.Mode-len(replay.Orange.Players)]
	// 	}
	// }

	// if len(replay.Orange.Players) > game.Match.Event.Mode {
	// 	sort.SliceStable(replay.Orange.Players, func(i, j int) bool {
	// 		return replay.Orange.Players[i].Stats.Core.Score > replay.Orange.Players[j].Stats.Core.Score
	// 	})
	// 	if len(replay.Blue.Players) >= game.Match.Event.Mode {
	// 		replay.Orange.Players = replay.Orange.Players[:game.Match.Event.Mode]
	// 	} else {
	// 		replay.Orange.Players = replay.Orange.Players[:game.Match.Event.Mode-len(replay.Blue.Players)]
	// 	}
	// }

	bluePlayers, err := BallchasingToPlayerInfos(client, replay.Blue.Players)
	if err != nil {
		return nil, err
	}

	if len(game.Blue.Players) == len(bluePlayers) {
		for i, player := range game.Blue.Players {
			bluePlayers[i].Player = player.Player
		}
	}

	game.Blue.Players = bluePlayers

	orangePlayers, err := BallchasingToPlayerInfos(client, replay.Orange.Players)
	if err != nil {
		return nil, err
	}

	if len(game.Orange.Players) == len(orangePlayers) {
		for i, player := range game.Orange.Players {
			orangePlayers[i].Player = player.Player
		}
	}

	game.Orange.Players = orangePlayers

	game.Blue.Team.Stats = PlayerStatsToTeamStats(game.Blue.Players)
	game.Orange.Team.Stats = PlayerStatsToTeamStats(game.Orange.Players)
	game.Blue.Winner = game.Blue.Team.Stats.Core.Goals > game.Orange.Team.Stats.Core.Goals
	game.Orange.Winner = game.Orange.Team.Stats.Core.Goals > game.Blue.Team.Stats.Core.Goals

	if replay.Blue.Stats.Ball != nil && replay.Orange.Stats.Ball != nil {
		game.Blue.Team.Stats.Ball = &octane.TeamBall{
			PossessionTime: replay.Blue.Stats.Ball.PossessionTime,
			TimeInSide:     replay.Blue.Stats.Ball.TimeInSide,
		}

		game.Orange.Team.Stats.Ball = &octane.TeamBall{
			PossessionTime: replay.Orange.Stats.Ball.PossessionTime,
			TimeInSide:     replay.Orange.Stats.Ball.TimeInSide,
		}
	}
	return game, nil
}

func BallchasingToPlayerInfos(client octane.Client, players []ballchasing.Player) ([]*octane.PlayerInfo, error) {
	var res []*octane.PlayerInfo

	var teamGoals float64
	for _, b := range players {
		player, err := BallchasingToPlayer(client, &b)
		if err != nil {
			return nil, err
		}

		playerInfo := &octane.PlayerInfo{
			Player: &octane.Player{
				ID:      player.ID,
				Slug:    player.Slug,
				Tag:     player.Tag,
				Country: player.Country,
			},
			Stats: BallchasingToPlayerStats(&b.Stats),
			Advanced: &octane.AdvancedStats{
				MVP: b.Mvp,
			},
		}

		if playerInfo.Stats.Core.Shots > 0 {
			playerInfo.Stats.Core.ShootingPercentage = float64(playerInfo.Stats.Core.Goals) / float64(playerInfo.Stats.Core.Shots) * 100
		}

		teamGoals += playerInfo.Stats.Core.Goals
		res = append(res, playerInfo)
	}

	pipeline := pipelines.Averages()
	data, err := client.Statlines().Pipeline(pipeline.Pipeline, pipeline.Decode)
	if err != nil {
		return nil, err
	}

	if len(data) == 0 {
		return nil, errors.New("no averages found")
	}

	for _, player := range res {
		if teamGoals > 0 {
			player.Advanced.GoalParticipation = float64(player.Stats.Core.Goals+player.Stats.Core.Assists) / teamGoals * 100
		}

		player.Advanced.Rating = Rating(data[0], player)
	}
	return res, nil
}

func BallchasingToPlayer(client octane.Client, b *ballchasing.Player) (*octane.Player, error) {
	p, _ := client.Players().FindOne(bson.M{
		"accounts": bson.M{
			"$elemMatch": bson.M{
				"platform": b.ID.Platform,
				"id":       b.ID.ID,
			},
		},
	})

	if p != nil {
		player := p.(octane.Player)
		return &octane.Player{
			ID:      player.ID,
			Slug:    player.Slug,
			Tag:     player.Tag,
			Country: player.Country,
		}, nil
	}

	p, _ = client.Players().FindOne(bson.M{
		"tag": bson.M{
			"$regex": primitive.Regex{
				Pattern: fmt.Sprintf("^%s$", b.Name),
				Options: "i",
			},
		},
	})

	if p != nil {
		player := p.(octane.Player)
		if _, err := client.Players().UpdateOne(bson.M{"_id": player.ID}, bson.M{
			"$addToSet": bson.M{
				"accounts": bson.M{
					"id":       b.ID.ID,
					"platform": b.ID.Platform,
				},
			},
		}); err != nil {
			return nil, err
		}

		return &octane.Player{
			ID:      player.ID,
			Slug:    player.Slug,
			Tag:     player.Tag,
			Country: player.Country,
		}, nil
	}

	id := primitive.NewObjectID()
	newPlayer := &octane.Player{
		ID:  &id,
		Tag: b.Name,
		Accounts: []*octane.Account{
			{
				ID:       b.ID.ID,
				Platform: b.ID.Platform,
			},
		},
	}
	newPlayer.Slug = PlayerSlug(newPlayer)

	if _, err := client.Players().InsertOne(newPlayer); err != nil {
		return nil, err
	}

	return &octane.Player{
		ID:   newPlayer.ID,
		Slug: PlayerSlug(newPlayer),
		Tag:  newPlayer.Tag,
	}, nil
}
