package helper

import (
	"github.com/octanegg/zsr/octane"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func UpdateEvent(client octane.Client, old, new *primitive.ObjectID) error {
	o, err := client.Events().FindOne(bson.M{"_id": old})
	if err != nil {
		return err
	}
	oldEvent := o.(octane.Event)

	n, err := client.Events().FindOne(bson.M{"_id": new})
	if err != nil {
		return err
	}
	newEvent := n.(octane.Event)

	if _, err := client.Matches().Update(
		bson.M{
			"event._id": oldEvent.ID,
		},
		bson.M{
			"$set": bson.M{
				"event": bson.M{
					"_id":    newEvent.ID,
					"name":   newEvent.Name,
					"mode":   newEvent.Mode,
					"region": newEvent.Region,
					"tier":   newEvent.Tier,
					"image":  newEvent.Image,
					"groups": newEvent.Groups,
				},
			},
		},
	); err != nil {
		return err
	}

	if _, err := client.Games().Update(
		bson.M{
			"match.event._id": oldEvent.ID,
		},
		bson.M{
			"$set": bson.M{
				"match.event": bson.M{
					"_id":    newEvent.ID,
					"name":   newEvent.Name,
					"mode":   newEvent.Mode,
					"region": newEvent.Region,
					"tier":   newEvent.Tier,
					"image":  newEvent.Image,
					"groups": newEvent.Groups,
				},
			},
		},
	); err != nil {
		return err
	}

	if _, err := client.Statlines().Update(
		bson.M{
			"game.match.event._id": oldEvent.ID,
		},
		bson.M{
			"$set": bson.M{
				"game.match.event": bson.M{
					"_id":    newEvent.ID,
					"name":   newEvent.Name,
					"mode":   newEvent.Mode,
					"region": newEvent.Region,
					"tier":   newEvent.Tier,
					"image":  newEvent.Image,
					"groups": newEvent.Groups,
				},
			},
		},
	); err != nil {
		return err
	}

	return nil
}
