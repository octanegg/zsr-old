package helper

import (
	"github.com/octanegg/zsr/octane"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func UpdateGame(client octane.Client, id *primitive.ObjectID) error {
	g, err := client.Games().FindOne(bson.M{"_id": id})
	if err != nil {
		return err
	}
	game := g.(octane.Game)

	var statlines []interface{}
	var records []interface{}
	blue, orange := GameToStatlines(&game)
	for _, statline := range append(blue, orange...) {
		statlines = append(statlines, statline)
	}
	for _, record := range StatlinesToRecords(append(blue, orange...)) {
		records = append(records, record)
	}

	if _, err := client.Statlines().Delete(bson.M{"game._id": game.ID}); err != nil {
		return err
	}

	if _, err := client.Statlines().Insert(statlines); err != nil {
		return err
	}


	if _, err := client.Records().Delete(bson.M{"game._id": game.ID}); err != nil {
		return err
	}

	if _, err := client.Records().Insert(records); err != nil {
		return err
	}


	return nil
}
