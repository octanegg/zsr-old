package helper

import (
	"github.com/octanegg/zsr/octane"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func UpdateTeam(client octane.Client, old, new *primitive.ObjectID) error {
	o, err := client.Teams().FindOne(bson.M{"_id": old})
	if err != nil {
		return err
	}
	oldTeam := o.(octane.Team)

	newTeam := o.(octane.Team)
	if old.Hex() != new.Hex() {
		n, err := client.Teams().FindOne(bson.M{"_id": new})
		if err != nil {
			return err
		}
		newTeam = n.(octane.Team)
	}

	if _, err := client.Players().Update(
		bson.M{"team._id": oldTeam.ID},
		bson.M{
			"$set": bson.M{
				"team": bson.M{
					"_id":    newTeam.ID,
					"slug":   newTeam.Slug,
					"name":   newTeam.Name,
					"region": newTeam.Region,
					"image":  newTeam.Image,
				},
			},
		}); err != nil {
		return err
	}

	if _, err := client.Statlines().Update(
		bson.M{"team.team._id": oldTeam.ID},
		bson.M{
			"$set": bson.M{
				"team.team": bson.M{
					"_id":   newTeam.ID,
					"slug":  newTeam.Slug,
					"name":  newTeam.Name,
					"image": newTeam.Image,
				},
			},
		}); err != nil {
		return err
	}

	if _, err := client.Games().Update(
		bson.M{"blue.team.team._id": oldTeam.ID},
		bson.M{
			"$set": bson.M{
				"blue.team.team": bson.M{
					"_id":   newTeam.ID,
					"slug":  newTeam.Slug,
					"name":  newTeam.Name,
					"image": newTeam.Image,
				},
			},
		}); err != nil {
		return err
	}

	if _, err := client.Games().Update(
		bson.M{"orange.team.team._id": oldTeam.ID},
		bson.M{
			"$set": bson.M{
				"orange.team.team": bson.M{
					"_id":   newTeam.ID,
					"slug":  newTeam.Slug,
					"name":  newTeam.Name,
					"image": newTeam.Image,
				},
			},
		}); err != nil {
		return err
	}

	if _, err := client.Matches().Update(
		bson.M{"blue.team.team._id": oldTeam.ID},
		bson.M{
			"$set": bson.M{
				"blue.team.team": bson.M{
					"_id":   newTeam.ID,
					"slug":  newTeam.Slug,
					"name":  newTeam.Name,
					"image": newTeam.Image,
				},
			},
		}); err != nil {
		return err
	}

	if _, err := client.Matches().Update(
		bson.M{"orange.team.team._id": oldTeam.ID},
		bson.M{
			"$set": bson.M{
				"orange.team.team": bson.M{
					"_id":   newTeam.ID,
					"slug":  newTeam.Slug,
					"name":  newTeam.Name,
					"image": newTeam.Image,
				},
			},
		}); err != nil {
		return err
	}

	return nil
}
