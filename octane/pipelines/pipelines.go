package pipelines

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// Pipeline .
type Pipeline struct {
	Pipeline []bson.M
	Decode   func(cursor *mongo.Cursor) (interface{}, error)
}

// New .
func New(stages ...bson.M) []bson.M {
	var pipeline []bson.M
	for _, stage := range stages {
		if stage != nil {
			pipeline = append(pipeline, stage)
		}
	}
	return pipeline
}

// Match .
func Match(m bson.M) bson.M {
	if m == nil {
		return nil
	}

	return bson.M{
		"$match": m,
	}
}

// Group .
func Group(m bson.M) bson.M {
	if m == nil {
		return nil
	}

	return bson.M{
		"$group": m,
	}
}

// Project .
func Project(m bson.M) bson.M {
	if m == nil {
		return nil
	}

	return bson.M{
		"$project": m,
	}
}

// Sort .
func Sort(key string, descending bool) bson.M {
	var order int
	switch descending {
	case true:
		order = -1
	case false:
		order = 1
	}
	return bson.M{
		"$sort": bson.M{key: order},
	}
}

// Limit .
func Limit(n int) bson.M {
	return bson.M{
		"$limit": n,
	}
}
