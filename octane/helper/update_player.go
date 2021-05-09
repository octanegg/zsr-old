package helper

import (
	"github.com/octanegg/zsr/octane"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func UpdatePlayer(client octane.Client, old, new *primitive.ObjectID) error {
	o, err := client.Players().FindOne(bson.M{"_id": old})
	if err != nil {
		return err
	}
	oldPlayer := o.(octane.Player)

	newPlayer := o.(octane.Player)
	if old.Hex() != new.Hex() {
		n, err := client.Players().FindOne(bson.M{"_id": new})
		if err != nil {
			return err
		}
		newPlayer = n.(octane.Player)
	}

	if _, err := client.Statlines().Update(
		bson.M{"player.player._id": oldPlayer.ID},
		bson.M{
			"$set": bson.M{
				"player.player": bson.M{
					"_id":     newPlayer.ID,
					"slug":    newPlayer.Slug,
					"tag":     newPlayer.Tag,
					"country": newPlayer.Country,
				},
			},
		}); err != nil {
		return err
	}

	if _, err := client.Games().Update(
		bson.M{"blue.players.player._id": oldPlayer.ID},
		bson.M{
			"$set": bson.M{
				"blue.players.$.player": bson.M{
					"_id":     newPlayer.ID,
					"slug":    newPlayer.Slug,
					"tag":     newPlayer.Tag,
					"country": newPlayer.Country,
				},
			},
		}); err != nil {
		return err
	}

	if _, err := client.Games().Update(
		bson.M{"orange.players.player._id": oldPlayer.ID},
		bson.M{
			"$set": bson.M{
				"orange.players.$.player": bson.M{
					"_id":     newPlayer.ID,
					"slug":    newPlayer.Slug,
					"tag":     newPlayer.Tag,
					"country": newPlayer.Country,
				},
			},
		}); err != nil {
		return err
	}

	if _, err := client.Matches().Update(
		bson.M{"blue.players.player._id": oldPlayer.ID},
		bson.M{
			"$set": bson.M{
				"blue.players.$.player": bson.M{
					"_id":     newPlayer.ID,
					"slug":    newPlayer.Slug,
					"tag":     newPlayer.Tag,
					"country": newPlayer.Country,
				},
			},
		}); err != nil {
		return err
	}

	if _, err := client.Matches().Update(
		bson.M{"orange.players.player._id": oldPlayer.ID},
		bson.M{
			"$set": bson.M{
				"orange.players.$.player": bson.M{
					"_id":     newPlayer.ID,
					"slug":    newPlayer.Slug,
					"tag":     newPlayer.Tag,
					"country": newPlayer.Country,
				},
			},
		}); err != nil {
		return err
	}

	return nil
}
