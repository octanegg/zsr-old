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

	n, err := client.Matches().FindOne(bson.M{"_id": new})
	if err != nil {
		return err
	}
	newMatch := n.(octane.Match)

	if _, err := client.Games().Update(
		bson.M{
			"match._id": oldMatch.ID,
		},
		bson.M{
			"$set": bson.M{
				"match._id":    newMatch.ID,
				"match.format": newMatch.Format,
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
				"game.match.format": newMatch.Format,
			},
		},
	); err != nil {
		return err
	}

	return nil
}
