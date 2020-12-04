package stats

import (
	"go.mongodb.org/mongo-driver/bson"
)

const maxRecords = 25

func (s *stats) GetGameRecords(filter, sort bson.M) ([]interface{}, error) {
	data, err := s.Statlines.Find(filter, sort, nil)
	if err != nil {
		return nil, err
	}

	if len(data) > maxRecords {
		data = data[:maxRecords]
	}

	return data, nil
}
