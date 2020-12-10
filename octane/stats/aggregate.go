package stats

import (
	"github.com/octanegg/zsr/octane/pipelines"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (s *stats) GetPlayerAggregate(filter, having bson.M) ([]interface{}, error) {
	data, err := s.Statlines.Pipeline(pipelines.PlayerAggregate(filter, having), toPlayer)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func toPlayer(cursor *mongo.Cursor) (interface{}, error) {
	var player pipelines.Player
	if err := cursor.Decode(&player); err != nil {
		return nil, err
	}
	return player, nil
}
