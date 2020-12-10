package stats

import (
	"github.com/octanegg/zsr/octane/collection"
	"go.mongodb.org/mongo-driver/bson"
)

const maxRecords = 25

func (s *stats) GetGameRecords(filter, sort bson.M) ([]interface{}, error) {
	data, err := s.Statlines.Find(filter, sort, &collection.Pagination{Page: 1, PerPage: 50})
	if err != nil {
		return nil, err
	}

	if len(data) > maxRecords {
		data = data[:maxRecords]
	}

	return data, nil
}
