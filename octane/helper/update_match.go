package helper

import (
	"github.com/octanegg/zsr/octane"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func UpdateMatch(client octane.Client, old, new *primitive.ObjectID) error {
	o, err := client.Matches().FindOne(bson.M{"_id": old})
	if err != nil {
		return err
	}
	oldMatch := o.(octane.Match)

	newMatch := o.(octane.Match)
	if old.Hex() != new.Hex() {
		n, err := client.Matches().FindOne(bson.M{"_id": new})
		if err != nil {
			return err
		}
		newMatch = n.(octane.Match)
	}

	var blueWinner, orangeWinner bool
	if newMatch.Blue != nil && newMatch.Orange != nil {
		blueWinner = newMatch.Blue.Score > newMatch.Orange.Score
		orangeWinner = newMatch.Orange.Score > newMatch.Blue.Score
	}

	if _, err := client.Games().Update(
		bson.M{
			"match._id": oldMatch.ID,
		},
		bson.M{
			"$set": bson.M{
				"match._id":           newMatch.ID,
				"match.slug":          newMatch.Slug,
				"match.format":        newMatch.Format,
				"blue.match_winner":   blueWinner,
				"orange.match_winner": orangeWinner,
			},
		},
	); err != nil {
		return err
	}

	if _, err := client.Statlines().Update(
		bson.M{
			"game.match._id": oldMatch.ID,
		},
		bson.M{
			"$set": bson.M{
				"game.match._id":    newMatch.ID,
				"game.match.slug":   newMatch.Slug,
				"game.match.format": newMatch.Format,
			},
		},
	); err != nil {
		return err
	}

	return nil
}

func UpdateMatchAggregate(client octane.Client, id *primitive.ObjectID) error {
	m, err := client.Matches().FindOne(bson.M{"_id": id})
	if err != nil {
		return err
	}
	match := m.(octane.Match)

	if match.Blue == nil || match.Orange == nil || match.Blue.Team == nil || match.Orange.Team == nil {
		return nil
	}

	res, err := client.Games().Find(bson.M{"match._id": id}, bson.M{"number": 1}, nil)
	if err != nil {
		return err
	}

	match.Blue.Score = 0
	match.Orange.Score = 0

	var games []*octane.Game
	for _, g := range res {
		game := g.(octane.Game)
		games = append(games, &game)
		if game.Blue.Team.Stats.Core.Goals > game.Orange.Team.Stats.Core.Goals {
			match.Blue.Score++
		} else {
			match.Orange.Score++
		}
	}

	var blueStatlines, orangeStatlines []*octane.Statline
	for _, game := range games {
		blue, orange := GameToStatlines(game)
		blueStatlines = append(blueStatlines, blue...)
		orangeStatlines = append(orangeStatlines, orange...)
	}

	match.ID = nil
	match.Blue.Players = StatlinesToAggregatePlayerStats(blueStatlines)
	match.Orange.Players = StatlinesToAggregatePlayerStats(orangeStatlines)
	match.Blue.Team.Stats = PlayerStatsToTeamStats(match.Blue.Players)
	match.Orange.Team.Stats = PlayerStatsToTeamStats(match.Orange.Players)
	match.ReverseSweepAttempt, match.ReverseSweep = ReverseSweep(games)
	match.Games = GamesToGameOverviews(games)
	match.Blue.Winner = match.Blue.Score > match.Orange.Score
	match.Orange.Winner = match.Orange.Score > match.Blue.Score

	if _, err := client.Matches().UpdateOne(bson.M{"_id": id}, bson.M{"$set": match}); err != nil {
		return err
	}

	return nil
}
